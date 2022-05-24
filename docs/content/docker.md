# Docker

**docker仓库** https://hub.docker.com/

**参考书籍** Docker实践(第2版)

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

### 解决apt-get/pip下载很慢的问题

更换国内源，如果基于ubuntu的镜像

```dockerfile
# syntax=docker/dockerfile:1
FROM ubuntu:20.04

RUN sed -i s@/archive.ubuntu.com/@/mirrors.aliyun.com/@g /etc/apt/sources.list && \
    apt-get clean && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
    curl \
    vim
```

如果是python或者go

```dockerfile
# syntax=docker/dockerfile:1
FROM python:3.10.2-bullseye

# 安装方便调试的工具
RUN sed -i 's/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    sed -i 's/security.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    sed -i 's/security-cdn.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    apt-get clean && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
    curl \
    vim

# pip更换国内源
RUN pip3 config set global.index-url https://mirrors.aliyun.com/pypi/simple \
    && pip3 config set install.trusted-host mirrors.aliyun.com
```

### 在容器运行图形化软件

在主机安装X11软件, Centos 7/8

```text
yum -y install xorg-x11-xauth
```

> Docker实践(第2版) 技巧29 在Docker里运行GUI

分两种场景

**在本地Linux启动docker的GUI**

以 [core](https://coreemu.github.io/core/) 为例，启动容器的命令

```text
sudo docker run -itd --name core -e DISPLAY -v /tmp/.X11-unix:/tmp/.X11-unix:rw --privileged core
```

这里从主机挂载了X11图形化软件，GUI需要的环境变量: DISPLAY

:::tip
DISPLAY是显示器名称
:::

**本地远程服务器使用docker exec的GUI(不推荐)**

如果是Win系统，先下载Xming, 然后参考博客 https://blog.csdn.net/ywxuan/article/details/118462658

```text
docker run -itd --name core --net=host -e DISPLAY -v $HOME/.Xauthority:/root/.Xauthority -v /tmp/.X11-unix:/tmp/.X11-unix:rw -v /opt/core:/opt/core --privileged core
```

关键参数: --net=host -e DISPLAY -v $HOME/.Xauthority:/root/.Xauthority -v /tmp/.X11-unix:/tmp/.X11-unix:rw

注意, 每次XShell重新连接，都要重启容器，然后进入容器修改DISPLAY的值跟宿主机的一致，才能打开图形化界面软件。因为这个过程挺繁琐的，所以不推荐这种方法。

**本地ssh直接容器，然后打开GUI(推荐)**

1. 运行容器
```text
docker run --net=host --name gui image_name
```

注意，关键参数只有--net=host, 不要挂载主机的 /.Xauthority 和 /tmp/.X11-unix ,可能会产生冲突

2. 容器启动ssh服务允许远程连接, 参考 [本地远程连接服务器的容器](#本地远程连接服务器的容器)
3. ssh连接容器, 假设端口号为222

当然，还有一种情况是服务器没有开放222端口，这时可以先加上-X参数ssh连接服务器, 然后服务器上再用-X参数ssh连接容器, 这样也能达到一样的效果。

**Unix系统**
假设本机是Ubuntu系统，并且装好了图形化软件
```text
ssh -X root@host -p 222
apt-get install xarclock
xarclock 
```

**Win系统**

用powershell连接到容器, 打开powershell设置DISPLAY
```text
PS C:\Users\39713> setx DISPLAY "localhost:0.0"
```
重启powershell让DISPLAY生效, 输出DISPLAY
```text
PS C:\Users\39713> $env:DISPLAY
localhost:0.0
```
有值代表设置成功

远程连接容器并打开GUI软件
```text
PS C:\Users\39713> ssh -Y root@host -p 222
```

```text
root@iZwz9ivwxya29whzkf6aejZ:~# apt-get install xarclock
root@iZwz9ivwxya29whzkf6aejZ:~# xarclock 
```

注意，-X是关键参数, 允许X11转发

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



