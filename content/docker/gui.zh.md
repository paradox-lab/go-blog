---
title: "在容器运行图形化软件"
date: 2022-06-15T16:58:49+08:00
draft: false
---

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

这里从主机挂载了X11图形化软件，GUI需要的环境变量: DISPLAY 。 DISPLAY是显示器名称

**本地远程服务器使用docker exec的GUI(不推荐)**

如果是Win系统，先下载Xming, 然后参考博客 https://blog.csdn.net/ywxuan/article/details/118462658

```text
docker run -itd --name core --net=host -e DISPLAY -v $HOME/.Xauthority:/root/.Xauthority -v /tmp/.X11-unix:/tmp/.X11-unix:rw -v /opt/core:/opt/core --privileged core
```

关键参数: ```--net=host -e DISPLAY -v $HOME/.Xauthority:/root/.Xauthority -v /tmp/.X11-unix:/tmp/.X11-unix:rw```

注意, 每次XShell重新连接，都要重启容器，然后进入容器修改DISPLAY的值跟宿主机的一致，才能打开图形化界面软件。因为这个过程挺繁琐的，所以不推荐这种方法。

**本地ssh直接容器，然后打开GUI(推荐)**

1. 运行容器
```text
docker run --net=host --name gui image_name
```

对于Linux系统, 关键参数只有 ```--net=host``` (Win系统不支持 ```--net=host```), 不要挂载主机的 /.Xauthority 和 /tmp/.X11-unix ,可能会产生冲突

2. 容器启动ssh服务允许远程连接, 参考 [本地远程连接服务器的容器](/go-blog/docker/#本地远程连接服务器的容器)
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