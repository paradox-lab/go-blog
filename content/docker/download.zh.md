---
title: "Download"
date: 2022-06-30T16:20:35+08:00
draft: true
---

# 解决下载慢的问题

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

## alpine(用apk add)

```dockerfile
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
```

或者

```dockerfile
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
```

## node.js

```dockerfile
RUN npm config set registry https://registry.npmmirror.com
```
