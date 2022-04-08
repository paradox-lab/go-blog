# Docker

**docker仓库** https://hub.docker.com/

**理论支撑** Docker实践(第2版)

![Docker实践](/images/Golang/Docker实践.png)

## 命令行使用

### commit

将当前的容器创建为一个镜像

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

.. code-block:: python
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

**本地远程服务器启动docker的GUI**

先下载Xming, 然后参考博客 https://blog.csdn.net/ywxuan/article/details/118462658

```text
docker run -itd --name core --net=host -e DISPLAY -v $HOME/.Xauthority:/root/.Xauthority -v /opt/core:/opt/core --privileged core
```

关键参数: --net=host -e DISPLAY -v $HOME/.Xauthority:/root/.Xauthority

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

* docker安装ssh服务
* 修改ssh配置文件，允许远程连接
* 本地生成公钥，将文本复制到容器中
* 将容器提交到新的镜像，重新运行容器，暴露22端口
* 阿里云开放暴露的端口
* 本地测试连接: ssh root@{host} -p 221

### 查看启动参数

```text
pip install runlike
runlike 容器名称
```

## 源码

docker: https://github.com/moby/moby



