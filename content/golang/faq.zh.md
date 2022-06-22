---
title: "Faq"
date: 2022-06-17T16:27:39+08:00
draft: false
---

# 常见问题

## 处理中文乱码

参见: https://studygolang.com/articles/8897

导入github.com/axgle/mahonia库

```go
enc := mahonia.NewDecoder("gbk")
goStr:= enc.ConvertString(robotgo.GetTitle())
fmt.Println(goStr)
```

注意: robot.GetTitle()返回C.GoString的值，遇到中文会显示乱码

## 如何永久阻塞程序

```go
sig := make(chan os.Signal, 2)
signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
<-sig
```

## http请求参数空格符转义

使用url.QueryEscape