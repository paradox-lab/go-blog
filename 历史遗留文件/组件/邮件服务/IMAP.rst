========================
IMAP
========================

依赖的第三方库

* `go-imap`_ An IMAP4rev1 library written in Go.

.. _go-imap: https://github.com/emersion/go-imap

客户端登录
------------------------

.. code-block:: go

    // 连接邮件服务器
    c, err := client.DialTLS("smtp.exmail.qq.com:993", nil)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Connected")

    // Don't forget to logout
    defer c.Logout()

    // 使用账号密码登录
    if err := c.Login("username", "******"); err != nil {
        log.Fatal(err)
    }
    log.Println("Logged in")

    // List mailboxes
    // 列出所有的邮件箱
    mailboxes := make(chan *imap.MailboxInfo, 10)
    done := make(chan error, 1)
    go func () {
        done <- c.List("", "*", mailboxes)
    }()

    log.Println("Mailboxes:")
    for m := range mailboxes {
        // 打印邮件箱名称
        log.Println("* " + m.Name)
    }

    if err := <-done; err != nil {
        log.Fatal(err)
    }

    // Select INBOX
    // 选择收件箱
    mbox, err := c.Select("INBOX", false)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Flags for INBOX:", mbox.Flags)

检索邮件
----------------------

.. code-block:: go

    // Get the last 4 messages
    // 获取最新的4封邮件
    from := uint32(1)
    to := mbox.Messages
    if mbox.Messages > 3 {
        // We're using unsigned integers here, only subtract if the result is > 0
        from = mbox.Messages - 3
    }
    seqset := new(imap.SeqSet)
    seqset.AddRange(from, to)

获取邮件内容
---------------------------

.. code-block:: go

    messages := make(chan *imap.Message, 10)
    done = make(chan error, 1)
    go func() {
        // 抓取邮件消息体传入到messages信道
        done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
    }()

    log.Println("Last 4 messages:")
    for msg := range messages {
        log.Println("* " + msg.Envelope.Subject)
    }

    if err := <-done; err != nil {
        log.Fatal(err)
    }

    log.Println("Done!")

抓取整封邮件

.. code-block:: go

    package main

    import (
        "fmt"
        "gitee.com/luzhenxiong/aje/task/goldenbrige"
        "github.com/emersion/go-imap"
        "github.com/emersion/go-message/mail"
        "io"
        "io/ioutil"
        "log"
        "os"
    )

    func main() {
        // Connect to server
        // 连接邮件服务器
        c, err := goldenbrige.NewImapClient()
        if err != nil {
            log.Fatal(err)
        }
        // Don't forget to logout
        defer c.Logout()

        // Select INBOX
        // 选择收件箱
        mbox, err := c.Select("INBOX", false)
        if err != nil {
            log.Fatal(err)
        }

        // Get the last message
        // 获取最新的邮件
        if mbox.Messages == 0 {
            log.Fatal("No message in mailbox")
        }
        seqSet := new(imap.SeqSet)
        seqSet.AddNum(mbox.Messages)

        // Get the whole message body
        // 获取邮件主体
        var section imap.BodySectionName

        items := []imap.FetchItem{section.FetchItem()}

        messages := make(chan *imap.Message, 1)
        go func() {
            if err := c.UidFetch(seqSet, items, messages); err != nil {
                log.Fatal(err)
            }
        }()

        msg := <-messages
        if msg == nil {
            log.Fatal("Server didn't returned message")
        }

        r := msg.GetBody(&section)
        if r == nil {
            log.Fatal("Server didn't returned message body")
        }

        // Create a new mail reader
        // 创建邮件阅读器
        mr, err := mail.CreateReader(r)
        if err != nil {
            log.Fatal(err)
        }

        // Print some info about the message
        // 打印消息体的一些信息
        header := mr.Header
        if date, err := header.Date(); err == nil {
            log.Println("Date:", date)
        }
        if from, err := header.AddressList("From"); err == nil {
            log.Println("From:", from)
        }
        if to, err := header.AddressList("To"); err == nil {
            log.Println("To:", to)
        }
        if subject, err := header.Subject(); err == nil {
            log.Println("Subject:", subject)
        }

        // Process each message's part
        // 处理消息体的每个part
        for {
            p, err := mr.NextPart()
            if err == io.EOF {
                break
            } else if err != nil {
                log.Fatal(err)
            }

            switch h := p.Header.(type) {
            case *mail.InlineHeader:
                // This is the message's text (can be plain-text or HTML)
                // text或者html
                b, _ := ioutil.ReadAll(p.Body)
                log.Println("Got text: ", string(b))
            case *mail.AttachmentHeader:
                // This is an attachment
                // 附件
                filename, _ := h.Filename()
                log.Println("Got attachment: ", filename)
                b, _ := ioutil.ReadAll(p.Body)
                file, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
                defer file.Close()
                n, err := file.Write(b)
                if err != nil {
                    fmt.Println("写入文件异常", err.Error())
                } else {
                    fmt.Println("写入Ok：", n)
                }

            }
        }
    }

.. seealso::

    上述代码参考自 `wiki`_ 。注意，Python是Fetch RFC822内容，但这里并没有用到 RFC822。

.. _wiki: https://github.com/emersion/go-imap/wiki/Fetching-messages

----------------------------------------

常见问题
========================================

如何做qp解码?
----------------------------------------

在Python的做法::

    subject, decode = decode_header(envelope.subject.decode())[0]
    subject = subject.decode(decode) if decode else str(subject)

在Go，只需要在入口函数开头加入一行代码即可解决

.. code-block:: go

    // 【字符集】  处理us-ascii和utf-8以外的字符集(例如gbk,gb2313等)时,
    imap.CharsetReader = charset.Reader

参考自 https://github.com/emersion/go-imap/wiki/Charset-handling