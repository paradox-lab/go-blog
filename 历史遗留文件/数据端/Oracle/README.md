# Oracle
# 第一部分 基础
## 安装
### 11g(11.2.0.4.0)
#### 客户端安装和配置
> 安装精简版即可
D:\instantclient-basic-windows.x64-11.2.0.4.0\instantclient_11_2

_配置_
- 环境变量Path=D:\instantclient-basic-windows.x64-11.2.0.4.0\instantclient_11_2
- 环境变量NLS_LANG=SIMPLIFIED CHINESE_CHINA.ZHS16GBK
- 环境变量TNS_ADMIN=D:\instantclient-basic-windows.x64-11.2.0.4.0\instantclient_11_2\network\admin
- 在instantclient_11_2目录下新建network\admin目录
- 在admin目录新建tnsnames.ora文件，写入以下内容：
```text
myDBName =

(DESCRIPTION =

    (ADDRESS = (PROTOCOL = TCP)(HOST = 127.0.0.1)(PORT = 1521))

    (CONNECT_DATA =

      (SERVER = DEDICATED)

      (SERVICE_NAME = dbname)

    )

)
```

> myDBName是客户端连接时填写或选择的数据库名称
> HOST是远程数据库的IP地址，SERVICE_NAME是远程数据库的服务名

做完以上动作，Python即可用cx_Oracle连接
```python
import cx_Oracle
cx_Oracle.connect('username/password@myDBName')
# 或
cx_Oracle.connect('username','password',"127.0.0.1:1521/DINSIGHT")
```

## pl/sql
### 存储过程调试传入date参数
![date参数](./img/存储过程调试存入Date类型参数.PNG)

> 传入to_date('8-1月-21')提示:'to_date('8-1月-21')' is not a valid date and time

# 第二部分 常用技巧
## 查找包含关键字的存储过程
```sql
select * from all_source where  type='PROCEDURE' and instr(lower(text),'insert into dif_hfs_sales_transaction_old')>0;
```
> all_source是系统表

## 常用系统表
```text
dba_开头..... 
dba_users    数据库用户信息
   dba_segments 表段信息
   dba_extents   数据区信息
   dba_objects   数据库对象信息
   dba_tablespaces  数据库表空间信息
   dba_data_files   数据文件设置信息
   dba_temp_files  临时数据文件信息
   dba_rollback_segs  回滚段信息
   dba_ts_quotas  用户表空间配额信息
   dba_free_space数据库空闲空间信息
   dba_profiles  数据库用户资源限制信息
   dba_sys_privs  用户的系统权限信息
   dba_tab_privs用户具有的对象权限信息
   dba_col_privs用户具有的列对象权限信息
   dba_role_privs用户具有的角色信息
   dba_audit_trail审计跟踪记录信息
   dba_stmt_audit_opts审计设置信息
   dba_audit_object  对象审计结果信息
   dba_audit_session会话审计结果信息
   dba_indexes用户模式的索引信息
user_开头
   user_objects  用户对象信息
   user_source  数据库用户的所有资源对象信息
   user_segments  用户的表段信息
   user_tables   用户的表对象信息
   user_tab_columns  用户的表列信息
   user_constraints  用户的对象约束信息
   user_sys_privs  当前用户的系统权限信息
   user_tab_privs  当前用户的对象权限信息
   user_col_privs  当前用户的表列权限信息
   user_role_privs  当前用户的角色权限信息
   user_indexes  用户的索引信息
   user_ind_columns用户的索引对应的表列信息
   user_cons_columns  用户的约束对应的表列信息
   user_clusters  用户的所有簇信息
   user_clu_columns 用户的簇所包含的内容信息
   user_cluster_hash_expressions   散列簇的信息
v$开头
   v$database   数据库信息
   v$datafile  数据文件信息
   v$controlfile控制文件信息
   v$logfile  重做日志信息
   v$instance  数据库实例信息
   v$log  日志组信息
   v$loghist 日志历史信息
   v$sga  数据库SGA信息
   v$parameter 初始化参数信息
   v$process  数据库服务器进程信息
   v$bgprocess  数据库后台进程信息
   v$controlfile_record_section  控制文件记载的各部分信息
   v$thread  线程信息
   v$datafile_header  数据文件头所记载的信息
   v$archived_log归档日志信息
   v$archive_dest  归档日志的设置信息
   v$logmnr_contents  归档日志分析的DML DDL结果信息
   v$logmnr_dictionary  日志分析的字典文件信息
   v$logmnr_logs  日志分析的日志列表信息
   v$tablespace  表空间信息
   v$tempfile  临时文件信息
   v$filestat  数据文件的I/O统计信息
   v$undostat  Undo数据信息
   v$rollname  在线回滚段信息
   v$session  会话信息
   v$transaction 事务信息
   v$rollstat   回滚段统计信息
   v$pwfile_users  特权用户信息
   v$sqlarea    当前查询过的sql语句访问过的资源及相关的信息
   v$sql         与v$sqlarea基本相同的相关信息
   v$sysstat   数据库系统状态信息
all_开头
   all_users  数据库所有用户的信息
   all_objects  数据库所有的对象的信息
   all_def_audit_opts  所有默认的审计设置信息
   all_tables  所有的表对象信息
   all_indexes所有的数据库对象索引的信息
session_开头
   session_roles   会话的角色信息
   session_privs   会话的权限信息
index_开头
   index_stats  索引的设置和存储信息
伪表
   dual  系统伪列表信息
```