# 语法

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