    /*=========================================
    procedure name:rcs_send_email
    Author:lzx
    Create date:2019/10/7 17:46:23
    Modify record:
  -------------------------
  2019/10/13  增加带附件发送功能
  2019/11/19  存储过程sub_splite_str增加初始化my_address_list数组，避免产生邮件接收人重复增加的问题
  2019/12/07  邮件标题和MIME内容增加Base64的转码
  -------------------------
  
    作用：用oracle发送邮件
    主要功能：1、支持多收件人。
              2、支持中文
              3、支持抄送人
              4、支持大于32K的附件
              5、支持多行正文
              6、支持多附件
              7、支持HTML邮件
    
    ===========================================*/
  
  PROCEDURE rcs_send_email(p_recipient       VARCHAR2, -- 邮件接收人
                           p_CC              VARCHAR2, --抄送人
                           p_subject         VARCHAR2, -- 邮件标题
                           p_message         clob default null, -- 纯文本邮件正文
                           p_html            clob default null, --html邮件
                           p_attachment_path varchar2, --附件地址（绝对路径）,多附件用逗号或分号分隔
                           o_statusflag      out char, --发送结果 Y-成功,F-失败
                           o_errtext         out varchar2) IS
    -- 下面四个变量请根据实际邮件服务器进行赋值
    v_mailhost VARCHAR2(30) := 'smtp.exmail.qq.com'; -- SMTP服务器地址
    v_user     VARCHAR2(30) := 'AsiaRCS@cyberway.net.cn'; -- 登录SMTP服务器的用户名
    v_pass     VARCHAR2(50) := 'RCS1qaz*2wsx'; -- 登录SMTP服务器的密码
    v_sender   VARCHAR2(50) := 'AsiaRCS@cyberway.net.cn'; -- 发送者邮箱，一般与 v_user 一样
  
    ---------------- 配置完成 -----------------
  
    v_conn UTL_SMTP.connection; -- 到邮件服务器的连接
    v_msg  VARCHAR2(4000); -- 邮件内容
    type address_list is table of varchar2(100) index by binary_integer;
    my_address_list address_list;
    type acct_list is table of varchar2(100) index by binary_integer;
    my_acct_list acct_list;
    v_receivers  varchar2(32767);
    v_CC         varchar2(32767);
    l_boundary   varchar2(255) default 'a1b2c3d4e3f2g1'; --邮件段体内的每个子段以此串定界
    l_body_html  clob := empty_clob; --This LOB will be the email message
    l_offset     number;
    l_filepos    pls_integer := 1;
    l_fil        bfile;
    l_file_len   number;
    l_modulo     number;
    l_pieces     number;
    l_data       raw(32767);
    l_amt        number := 8580;
    l_buf        raw(32767);
    l_amount     number;
  
    --分割邮件地址或者附件地址
    procedure sub_splite_str(p_str varchar2, p_splite_flag int default 1) is
      l_addr varchar2(254) := '';
      l_len  int;
      l_str  varchar2(4000);
      j      int := 0; --表示邮件地址或者附件的个数
    begin
      my_address_list.delete; --初始化数组
      /*处理接收邮件地址列表，将;转换为,等*/
      l_str := trim(rtrim(replace(p_str, ';', ','), ','));
    
      l_len := length(l_str);
      for i in 1 .. l_len loop
        if substr(l_str, i, 1) <> ',' then
          l_addr := l_addr || substr(l_str, i, 1);
        else
          j := j + 1;
          if p_splite_flag = 1 then
            --表示处理邮件地址
            --前后需要加上'<>'，否则很多邮箱将不能发送邮件
            l_addr := '<' || l_addr || '>';
            --调用邮件发送过程
            my_address_list(j) := l_addr;
          elsif p_splite_flag = 2 then
            --表示处理附件名称
            my_acct_list(j) := l_addr;
          end if;
          l_addr := '';
        end if;
        if i = l_len then
          j := j + 1;
          if p_splite_flag = 1 then
            l_addr := '<' || l_addr || '>';
            my_address_list(j) := l_addr;
          elsif p_splite_flag = 2 then
            my_acct_list(j) := l_addr;
          end if;
        end if;
      end loop;
    end;
  
    --返回附件源文件所在目录或者名称
    function sub_get_file(p_file varchar2, p_get int) return varchar2 is
      --p_get=1 表示返回目录
      --p_get=2 表示返回文件名
      l_file varchar2(1000);
    begin
      if instr(p_file, '\') > 0 then
        if p_get = 1 then
          l_file := substr(p_file, 1, instr(p_file, '\', -1) - 1) || '\';
        elsif p_get = 2 then
          l_file := substr(p_file,
                           - (length(p_file) - instr(p_file, '\', -1)));
        end if;
      elsif instr(p_file, '/') > 0 then
        if p_get = 1 then
          l_file := substr(p_file, 1, instr(p_file, '/', -1) - 1);
        elsif p_get = 2 then
          l_file := substr(p_file,
                           - (length(p_file) - instr(p_file, '/', -1)));
        end if;
      end if;
      return l_file;
    end;
  
    procedure sub_attachment(filename in varchar2, dt_name in varchar2) is
      --dt_name目录名称
      l_filename varchar2(1000);
      -- l_amount   number;
    begin
      --得到附件文件名称
      l_filename := sub_get_file(filename, 2);
      v_msg      := v_msg || UTL_TCP.CRLF || '--' || l_boundary ||
                    UTL_TCP.CRLF;
    
      v_msg    := v_msg || 'Content-Type: application/octet-stream;name="' ||
                  l_filename || '"' || UTL_TCP.CRLF ||
                  'Content-Transfer-Encoding: base64' || UTL_TCP.CRLF ||
                  'Content-Disposition: attachment;filename="' ||
                  l_filename || '"' || UTL_TCP.CRLF || UTL_TCP.CRLF;
      l_offset := dbms_lob.getlength(l_body_html) + 1;
      dbms_lob.write(l_body_html, length(v_msg), l_offset, v_msg);
      --dbms_output.put_line(l_body_html);
    
      --把附件分成多份，这样可以发送超过32k的附件
      l_filepos  := 1; --重置offset，在发送多个附件时，必须重置
      l_fil      := bfilename(dt_name, l_filename);
      l_file_len := dbms_lob.getlength(l_fil);
      l_modulo   := mod(l_file_len, l_amt);
      l_pieces   := trunc(l_file_len / l_amt);
      if (l_modulo <> 0) then
        l_pieces := l_pieces + 1;
      end if;
      dbms_lob.fileopen(l_fil, dbms_lob.file_readonly);
      l_data   := null;
      l_amount := l_amt;
    
      for i in 1 .. l_pieces loop
        dbms_lob.read(l_fil, l_amount, l_filepos, l_buf);
      
        l_filepos  := i * l_amount + 1;
        l_file_len := l_file_len - l_amount;
      
        l_offset := dbms_lob.getlength(l_body_html) + 1;
        dbms_lob.write(l_body_html,
                       length(utl_raw.cast_to_varchar2(utl_encode.base64_encode(l_buf))),
                       l_offset,
                       utl_raw.cast_to_varchar2(utl_encode.base64_encode(l_buf)));
      
        if i = l_pieces then
          l_amount := l_file_len;
        end if;
      end loop;
    
      dbms_lob.fileclose(l_fil);
      /*exception
        when others then
          dbms_lob.fileclose(l_fil);
          sub_end_boundary(conn);
          raise;
      end; --结束处理二进制附件*/
    end; --结束过程attachment
  
  BEGIN
    --连接邮件服务器
    v_conn := UTL_SMTP.open_connection(host                          => v_mailhost,
                                       port                          => '25',
                                       secure_connection_before_smtp => FALSE);
    --向邮件服务器打招呼
    UTL_SMTP.ehlo(v_conn, v_mailhost);
  
    UTL_SMTP.command(v_conn, 'AUTH LOGIN'); -- smtp服务器登录校验
    UTL_SMTP.command(v_conn,
                     UTL_RAW.cast_to_varchar2(UTL_ENCODE.base64_encode(UTL_RAW.cast_to_raw(v_user))));
    UTL_SMTP.command(v_conn,
                     UTL_RAW.cast_to_varchar2(UTL_ENCODE.base64_encode(UTL_RAW.cast_to_raw(v_pass))));
    UTL_SMTP.mail(v_conn, '<' || v_sender || '>'); --设置发件人
  
    sub_splite_str(p_recipient); --处理邮件地址
    for k in 1 .. my_address_list.count loop
      v_receivers := v_receivers || my_address_list(k) || ';';
      utl_smtp.rcpt(v_conn, my_address_list(k)); --设置收件人
    end loop;
  
    -- 创建要发送的邮件内容 注意报头信息和邮件正文之间要空一行
  
    v_msg := v_msg || 'Date:' || TO_CHAR(SYSDATE, 'yyyy mm dd hh24:mi:ss') ||
             UTL_TCP.CRLF;
    v_msg := v_msg || 'From: ' || v_sender || '' || UTL_TCP.CRLF;
    v_msg := v_msg || 'To: ' || v_receivers || '' || UTL_TCP.CRLF;
  
    IF p_CC is not null then
      sub_splite_str(p_CC); --处理邮件地址
      for k in 1 .. my_address_list.count loop
        v_CC := v_CC || my_address_list(k) || ';';
        utl_smtp.rcpt(v_conn, my_address_list(k));
      end loop;
      v_msg := v_msg || 'Cc: ' || v_CC || UTL_TCP.CRLF; --抄送
    END IF;
  
    --v_msg := v_msg || 'Bcc: ' || p_recipient || UTL_TCP.CRLF;--密送
    v_msg := v_msg || 'Subject: ' || '=?GB2312?B?' ||
             replace(replace(utl_raw.cast_to_varchar2(utl_encode.base64_encode(utl_raw.cast_to_raw(p_subject))),
                             chr(10),
                             chr(13)),
                     chr(13),
                     '') || '?=' || UTL_TCP.CRLF; --邮件标题
    v_msg := v_msg || 'Mime-Version: 1.0' || UTL_TCP.CRLF;
    v_msg := v_msg || 'Content-Type: multipart/mixed;boundary=' || chr(34) ||
             l_boundary || chr(34) || UTL_TCP.CRLF; -- 这前面是报头信息
  
    --dbms_output.put_line(v_msg);
    -- Write the headers 写入报头内容
    dbms_lob.createtemporary(l_body_html, false, 10); --建立临时LOB
    dbms_lob.write(l_body_html, length(v_msg), 1, v_msg); --把邮件内容写入临时LOB    
  
    IF p_message IS NOT NULL THEN
      -- Write the text boundary,邮件分段(纯文本)
      l_offset := dbms_lob.getlength(l_body_html) + 1;
      v_msg    := UTL_TCP.CRLF || '--' || l_boundary || UTL_TCP.CRLF;
      v_msg    := v_msg || 'content-type: text/plain; charset="GB2312"' ||
                  UTL_TCP.CRLF;
      v_msg    := v_msg || 'Content-Transfer-Encoding: base64' ||
                  UTL_TCP.CRLF || UTL_TCP.CRLF;
      dbms_lob.write(l_body_html, length(v_msg), l_offset, v_msg);
    
      --Write the plain text portion of the email，写入纯文本邮件内容
      l_offset := dbms_lob.getlength(l_body_html) + 1;
      dbms_lob.write(l_body_html,
                     length(utl_raw.cast_to_varchar2(utl_encode.base64_encode(utl_raw.cast_to_raw(p_message)))),
                     l_offset,
                     utl_raw.cast_to_varchar2(utl_encode.base64_encode(utl_raw.cast_to_raw(p_message))));
    END IF;
  
    IF p_html IS NOT NULL THEN
      -- Write the HTML boundary,邮件分段(html)
      v_msg    := UTL_TCP.CRLF || UTL_TCP.CRLF || '--' || l_boundary ||
                  UTL_TCP.CRLF;
      v_msg    := v_msg || 'Content-type: text/html; charset="GB2312"' ||
                  UTL_TCP.CRLF;
      v_msg    := v_msg || 'Content-Transfer-Encoding: base64' ||
                  UTL_TCP.CRLF || UTL_TCP.CRLF;
      l_offset := dbms_lob.getlength(l_body_html) + 1;
      dbms_lob.write(l_body_html, length(v_msg), l_offset, v_msg);
    
      -- Write the HTML portion of the message,写入html邮件内容
      l_offset := dbms_lob.getlength(l_body_html) + 1;
      dbms_lob.write(l_body_html,
                     length(utl_raw.cast_to_varchar2(utl_encode.base64_encode(utl_raw.cast_to_raw(p_html)))),
                     l_offset,
                     utl_raw.cast_to_varchar2(utl_encode.base64_encode(utl_raw.cast_to_raw(p_html))));
    
    END IF;
    -- v_msg := v_msg || p_message || UTL_TCP.CRLF; -- 这个是邮件正文
    -- v_msg := v_msg || 'X-Mailer: Foxmail 7, 1, 3, 48[cn]' || UTL_TCP.CRLF; --发送客户端
  
    v_msg := null;
  
    --添加附件 
    IF p_attachment_path is not null then
      sub_splite_str(p_attachment_path, 2); --获取附件名称列表
    
      for k in 1 .. my_acct_list.count loop
        sub_attachment(filename => my_acct_list(k), dt_name => 'RCSDIR'); --把附件内容写入MIME
      end loop;
    END IF;
  
    -- Write the final html boundary
    v_msg    := UTL_TCP.CRLF || UTL_TCP.CRLF || '--' || l_boundary || '--';
    l_offset := dbms_lob.getlength(l_body_html) + 1;
    dbms_lob.write(l_body_html, length(v_msg), l_offset, v_msg);
  
    -------------------------------------------------------------------------------------------  
    UTL_SMTP.open_data(v_conn); --打开流
    --   UTL_SMTP.write_raw_data(v_conn, UTL_RAW.cast_to_raw(l_body_html)); --这样写标题和内容都能用中文
  
    l_offset := 1;
    l_amount := 1900;
  
    while l_offset < dbms_lob.getlength(l_body_html) loop
      utl_smtp.write_raw_data(v_conn,
                              utl_raw.cast_to_raw(dbms_lob.substr(l_body_html,
                                                                  l_amount,
                                                                  l_offset)));
      l_offset := l_offset + l_amount;
      l_amount := least(1900, dbms_lob.getlength(l_body_html) - l_amount);
    end loop;
  
    UTL_SMTP.close_data(v_conn); --关闭流
    UTL_SMTP.quit(v_conn); --关闭连接
    dbms_lob.freetemporary(l_body_html); --释放内存
    o_statusflag := 'Y';
  EXCEPTION
    WHEN OTHERS THEN
      o_statusflag := 'F';
      o_errtext    := DBMS_UTILITY.format_error_stack;
      DBMS_OUTPUT.put_line(DBMS_UTILITY.format_error_stack);
      DBMS_OUTPUT.put_line(DBMS_UTILITY.format_call_stack);
  END rcs_send_email;