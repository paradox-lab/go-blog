# Mysql
> MySQL是一个关系型数据库管理系统，由瑞典MySQL AB 公司开发，属于 Oracle 旗下产品。

# Usage
## 拼接语句
```text
contact
简介:将查询结果拼接成一个字符串，返回结果为连接参数产生的字符串。如有任何一个参数为NULL ，则返回值为 NULL。
样例: eg：select contact('11','22','33'); 返回结果：112233
```

## 时间函数
```text
MySQL中有5个函数需要计算当前时间的值：

NOW.返回时间，格式如：2012-09-23 06:48:28
CURDATE，返回时间的日期，格式如：2012-09-23
CURTIME，返回时间，格式如：06:48:28
UNIX_TIMESTAMP，返回时间整数戳，如：1348408108
SYSDATE，返回时间，格式和time()函数返回时间一样，但是有区别。
除了本身定义所返回的区别以外，另一个区别是：前四个函数都是返回基于语句的开始执行时间，而SYSDATE返回time的值。

我们一般在执行语句的时候，都是用NOW()，因为SYSDATE获取当时实时的时间，这有可能导致主库和从库是执行的返回值是不一样的，导致主从数据不一致。
```

# DDL
## 建表
_复制表结构_
```sql
create table example_bak like example;
```

# DML
```sql
-- 开启事务
START TRANSACTION;
```