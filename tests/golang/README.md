# Go

## 包管理

**wiki** [https://github.com/golang/go/wiki/Modules](https://github.com/golang/go/wiki/Modules)

工程目录结构参考

```text
go.mod
go.sum
repo\main.go
     myPackage\myCode.go
     util.go 'package repo
     util_test.go
```

::: tip
每个包最多只能存在一个main.go文件
:::

---------------------------------------------

### 初始化

```shell
go mod init github.com/my/repo
```

如果不放远程仓库，可以省略仓库地址

```shell
go mod init repo
```

自动增加模块依赖和sums

```shell
go mod tidy
```

### 引用自定义包

```go
import "repo/repo/myPackage"
```

### go.mod

go.mod文件的4个关键字

* module语句指定包的名字（路径）
* require语句指定的依赖项模块
* replace语句可以替换依赖项模块
* exclude语句可以忽略依赖项模块

引入自定义包(不在项目内)

项目A引入自定义包B

A项目的main.go

```go
import "B/core/management"
```

A项目的go.mod需要加入以下内容

```text
require B v0.0.0

replace B => E:\godjan\B
```

::: tip
replace 改用相对路径也可以
:::

### go命令行

```golang
go generate
```

使用go mod并安装go依赖包

## 单元测试

:标准库单元测试学习进度: archive/tar/tar_test.go

-----------------------------------------------------

单元测试函数的定义一

* 创建一个名称以 _test.go 结尾的文件
* 文件包含 `TestXxx` 函数

单元测试函数的定义二

* package名以 _test 为结尾的文件
* 这个文件的函数

例子: archive/tar/example_test.go

运行go test命令将运行编写好的单元测试函数

测试函数运行子测试
===============================

使用T.Run方法，参考例子: archive/tar/reader_test.go line:622

并行测试
===============================

测试函数第一行加上 `t.Parallel()`

使用案例: chromedp/eval_test.go的TestEvaluateNumber函数

---------------------------------------------

设计模式
==============================

数据驱动测试，可参考 archive/tar/reader_test.go