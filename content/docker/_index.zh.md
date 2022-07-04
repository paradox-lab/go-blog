---
title: "Docker"
date: 2022-05-24T17:42:09+08:00
draft: false
---

**docker仓库** https://hub.docker.com/

**参考书籍** Docker实践(第2版)

## 安装

https://docs.docker.com/desktop/windows/install/

https://docs.microsoft.com/zh-cn/windows/wsl/install

## 常见问题

### 容器互相访问

**参考博客** https://www.cnblogs.com/shenh/p/9714547.html

server容器需要调用grpc容器服务以执行用户上传的代码，访问地址是"<容器名称>:<服务端口号>"。

在django admin pro项目中, 容器名称已在 ``docker-compose.yml`` 定义好

``/server/helper/grpc/client.py``

```python
class BaeGrpcClient:
    def __init__(self, address: str = settings.GRPC_ADDRESS):
        self.address = address
        ...
```

``/server/django_admin_pro/settings/prod.py``

```python
# grpc容器访问地址
GRPC_ADDRESS = "grpc:50051"
```

### 对已启动的容器修改端口

https://www.cnblogs.com/fps2tao/p/10557257.html

### docker run的volume

不指定-v参数时，默认放到 ``/var/lib/docker/{dir}`` 路径下, 例如

```shell
docker run -v db:/opt/db im
```

等价于 ``-v /var/lib/docker/db:/opt/db``

注意, wsl是看不到 ``/var/lib/docker`` 目录的，建议使用绝对路径

### 本地远程连接服务器的容器

https://www.cnblogs.com/jesse131/p/13543308.html

步骤清单

1. docker安装ssh服务
2. 修改ssh配置文件，允许远程连接

因为这些配置默认是注释掉的，所以除了用vim打开修改，也可以用追加的方式添加配置

```text
echo "UseDNS no" >> /etc/ssh/sshd_config
echo "AddressFamily inet" >> /etc/ssh/sshd_config
echo "SyslogFacility AUTHPRIV" >> /etc/ssh/sshd_config 
echo "PermitRootLogin yes" >> /etc/ssh/sshd_config
echo "PasswordAuthentication yes" >> /etc/ssh/sshd_config
```

修改端口号(可选)
```text
echo "Port 222" >> /etc/ssh/sshd_config 
```

3. 重启ssh服务
```text
/etc/init.d/ssh restart
```

4. 修改root密码

```text
passwd
```

5. 将容器提交到新的镜像，重新运行容器，暴露22端口
6. 阿里云开放暴露的端口
7. 本地测试连接: ssh root@{host} -p 221

### 本地ssh连接容器

配置ssh, 过程参考 [本地远程连接服务器的容器](#本地远程连接服务器的容器)

假设暴露端口222, 本地测试连接: ssh root@localhost -p 222

### 查看启动参数

```text
pip install runlike
runlike 容器名称
```

### vim从外部粘贴文本

如果粘贴外部文本显示可视化模式时, 用冒号输入命令

```text
set mouse-=a
```

之后就可以正常使用粘贴了。

### vim无法粘贴

```text
:set mouse-=a
```

## docker版本的软件

### jenkins

https://www.jenkins.io/zh/doc/book/installing/#docker

```dockerfile
# Docker实践(第2版) 技巧66
FROM jenkins/jenkins:lts-jdk11

USER root

RUN rm -rf /etc/localtime
RUN ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

RUN sed -i 's/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    sed -i 's/security.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    sed -i 's/security-cdn.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    apt-get clean && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
    python3.9 \
    python3-pip

RUN pip3 config set global.index-url https://mirrors.aliyun.com/pypi/simple && \
    pip3 config set install.trusted-host mirrors.aliyun.com && \
    pip3 install docker-compose
```

docker run
```text
docker run --name jenkins_server -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock -v /var/jenkins_home:/var/jenkins_home -v /usr/bin/docker:/usr/bin/docker -d --privileged jenkins_server
```

### nginx

https://hub.docker.com/_/nginx

## 源码

docker: https://github.com/moby/moby



