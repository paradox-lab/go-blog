---
title: "Golang"
date: 2022-05-24T17:42:09+08:00
draft: false
---

# Go

**github地址** https://github.com/golang/go

**Go参考** https://golang.google.cn/ref/spec

**awesome-go** https://github.com/avelino/awesome-go

**awesome-go中文版** https://github.com/jobbole/awesome-go-cn

**在线跑代码** https://goplay.tools/

**官网** https://golang.google.cn/

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

### 测试函数运行子测试

使用T.Run方法，参考例子: archive/tar/reader_test.go line:622

### 并行测试

测试函数第一行加上 `t.Parallel()`

使用案例: chromedp/eval_test.go的TestEvaluateNumber函数

---------------------------------------------

### 设计模式

数据驱动测试，可参考 archive/tar/reader_test.go

## linux下载和安装
> 下载地址: https://golang.google.cn/doc/install

`export PATH=$PATH:/usr/local/go/bin`

```text
假设是WSL的子系统使用Go，下载linux版Go安装包，接着按照教程做即可。
```

**升级**
> 例如1.16.4升级1.16.5
```commandline
go get golang.org/dl/go1.16.5
go1.16.5 download
```

从16升级到17同理

> 执行download命令时如果提示go1.16.5文件不存在，先切换至gopath变量的bin目录下执行

## 修改国内源
windos
```text
然后我们需要进行更改Go的GOPROXY值
go env -w GOPROXY=https://goproxy.cn,direct

SETX来设置一下Windows的环境变量
SETX GOPROXY=https://goproxy.cn,direct

接下来就可以正常使用了
go get -v github.com/astaxie/beego
```

linux
```text
$ export GOPROXY=https://goproxy.cn

或者
$ echo "export GO111MODULE=on" >> ~/.profile
$ echo "export GOPROXY=https://goproxy.cn" >> ~/.profile
$ source ~/.profile
```

## 项目工程结构

sample
```text
README.md
LICENSE
go.mod
go.sum
main.go
util.go 'package main
util_test.go
sample/sample.go 'package sample
      /test_sample.go 
      /other/main.go 'package main 可直接go run或者go build(go build后的可执行文件需要放置在go.mod所在目录下执行)
                     
```
> 执行:sample/other/main.go
```commandline
go run sample/other/main.go
```

# 第一部分 语法基础
## 25个关键字
### chan(信道)
> 信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。
```golang
ch <- v    // 将 v 发送至信道 ch。
v := <-ch  // 从 ch 接收值并赋予 v。
```
> “箭头”就是数据流的方向。

**创建chan**
```golang
// 双项通道，支持读写
ch := make(chan int)
// 单项通道，只能读
var ch_read <-chan int
// 单项通道，只能写
var ch_write chan<- int
```
> 定义只读和只写的channel意义不大，一般用于在参数传递中

```text
默认情况下，发送和接收操作在另一端准备好之前都会阻塞。
这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。
```

**遍历chan**
> 原文链接：https://blog.csdn.net/zhaominpro/article/details/77584534

> 读channel时，第二个值返回为false时表示channel被关闭。

方法一:
```golang
for{
    if value,ok:=<-ch;ok{
        //do somthing
    }else{
        break;//表示channel已经被关闭，退出循环
    }
}
```

方法二:
```golang
//range
ch:=make(chan int ,3)
ch<-1
ch<-2
ch<-3

for value:=range ch{
    fmt.Print(value)
}

//输出：123
//然后会一直阻塞当前协程，如果在其他协程中调用了close(ch),那么就会跳出for range循环。这也就是for range的特别之处
```

### else
```golang
   if age >= 18 {
        fmt.Println("成年人")
    } else {
        fmt.Println("未成年")
    }
```

### goto
> 将控制转移到被标记的语句
```golang
func main() {
	var C, c int //声明变量
	C = 1        /*这里不写入for循环是因为for语句执行之初会将C的值变为1，当goto A时for
	语句会重新执行（不是重新一轮循环） */
LOOP:
	for c < 50 {
		C++ //C=1不能写入for这里就不能写入
		for c = 2; c < C; c++ {
			if C%c == 0 {
				goto LOOP //若发现因子不是素数
			}
		}
		fmt.Printf("%d \t", c)
	}
}
```

### if
```txt
语句      描述
if语句        if语句由一个布尔表达式后紧跟一个或多个语句组成
if...else语句 if语句后可以使用可选的else语句，else语句中的表达式再布尔表达式为false时执行
if嵌套语句      可以在if或else if语句中潜入一个或多个if或else if语句
```

**if语句**
```txt
if 布尔表达式{
    /* 在布尔表达式为true时执行 */
}
```
```golang
if num%2 == 0 {
    fmt.Println(num, "偶数")
} else {
    fmt.Println(num, "奇数")
}
```

**if...else语句**
```txt
if 布尔表达式{
    /* 在布尔表达式为true时执行 */
} else if {
    /* 在布尔表达式为true时执行 */
...
} else{
    /* 在布尔表达式为false时执行 */
}
```

```golang
    if score >= 90 {
        fmt.Println("优秀")
    } else if score >= 80 {
        fmt.Println("良好")
    } else if score >= 70 {
        fmt.Println("中等")
    } else if score > 60 {
        fmt.Println("及格")
    } else if score < 60 {
        fmt.Println("不及格")
    }
```

**特殊写法**
```golang
if num := 10; num%2 == 0 {
    fmt.Println(num, "偶数")
} else {
    fmt.Println(num, "奇数")
}
```
> 需要注意的是，num的定义在if里，那么只能够在该if...else语句块中使用，否则编译器会报错。

### struct
> 结构体是由一系列相同类型或不同类型的数据构成的数据集合。

> 结构体是值类型。因此，如果要修改值，需要赋值指针地址。

```text
type 类型名 struct{
        成员属性1 类型1
        成员属性2 类型2
        成员属性3,成员属性4 类型 3
        ...
}
```

在使用结构体的过程中注意以下3点。
- 类型名是标识结构体的名称，在同一个包内不能重复。
- 结构体的属性，也叫字段，必须唯一。
- 同类型的成员属性可以写在一行。

```golang
type Teacher struct {
	name string
	age  int8
	sex  byte
}
```

**实例化**
```golang
type Emp struct {
	name string
	age  int8
	sex  byte
}

emp1 := new(Emp)
////////////////////////////
(*emp1).name = "David"
(*emp1).age = 30
(*emp1).sex = 1
// 或
emp1.name = "David2"
emp1.age = 31
emp1.sex = 1
```

**匿名结构体**
> 匿名结构体就是没有名字的结构体。无需通过type关键字就可以直接使用。创建匿名结构体时，同时要创建对象。匿名结构体由结构体定义和键值
对初始化两部分组成。语法格式示例如下。
```text
变量名:=struct{
//定义成员属性
} {//初始化成员属性}
```

```golang
//匿名结构体
addr := struct {
    province, city string
}{"山西省", "西安市"}
```

**匿名字段**
> 匿名字段就是在结构体中的字段没有名字，只包含一个没有字段名的类型。这些字段被称为匿名字段。

> 如果字段没有名字，那么默认使用类型作为字段名，同一个类型只能有一个匿名字段。结构体嵌套中采用匿名结构体字段可以模拟继承关系。

```golang
type User struct {
	string
	byte
	int8
	float64
}

//实例化结构体
user := User{"Steven", 'm', 35, 177.5}
fmt.Println(user)
//如果想依次输出姓名、年龄、身高、性别
fmt.Printf("姓名：%s \n", user.string)
fmt.Printf("身高：%.2f \n", user.float64)
fmt.Printf("性别：%c \n", user.byte)
fmt.Printf("年龄：%d \n", user.int8)
```

**结构体嵌套**
> 将一个结构体作为另一个结构体的属性(字段)，这种结构就是结构体嵌套。

结构体嵌套可以模拟面向对象编程中的以下两种关系。
- 聚合关系：一个类作为另一个类的属性。
- 继承关系：一个类作为另一个类的字类。子类和父类的关系。

_模拟聚合关系_
```golang
type Address struct {
	province, city string
}

type Person struct {
	name    string
	age     int
	address *Address
}

//模拟结构体对象之间的聚合关系
p := Person{}
p.name = "Steven"
p.age = 35
//赋值方式1
addr := Address{}
addr.province = "北京市"
addr.city = "海淀区"
p.address = &addr
```

_模拟继承关系_
```golang
type Person struct {
	name string
	age  int
	sex  string
}

type Student struct {
	Person
	schoolName string
}

//1.实例化并初始化Person
p1 := Person{"Steven", 35, "男"}
fmt.Println(p1)
fmt.Println("-----------------")
//2.实例化并初始化Student
//写法1:
s1 := Student{p1, "北航软件学院"}
//写法2:
s2 := Student{Person{"Josh", 30, "男"}, "北外高翻学院"}
//写法3:
s3 := Student{Person: Person{
    name: "Penn",
    age:  19,
    sex:  "男",
},
    schoolName: "北大元培学院",
}
//写法4:
s4 := Student{}
s4.name = "Daniel"
s4.sex = "男"
s4.age = 12
s4.schoolName = "北京十一龙樾"
```

**方法**
和函数的区别
```text
1.含义不同
函数(function)是一段具有独立功能的代码，可以被反复多次调用，从而实现代码复用。而方法(method)是一个类的行为功能，只有该类的对象才能调用。

2.方法有接受者，而函数无接受者
Go语言的方法(method)是一种作用于特定类型变量的函数。这种特定类型变量叫作接受者(receiver)。接收者的概念类似于传统面向对象语言中的this或
self关键字。
Go语言的接受者强调了方法具有作用对象，而函数没有作用对象。一个方法就是一个包含了接受者的函数。
Go语言中，接受者可以是结构体，也可以是结构体类型外的其他任何类型。

3.函数不可以重名，而方法可以重名。
只要接受者不同，方法名就可以相同。
```

语法格式
```text
func (接受者变量 接受者类型) 方法名(参数列表) (返回值列表) {
//方法体
}
```
> 接受者在func关键字和方法名之间编写，接受者可以是struct类型或非struct类型，可以是指针类型或非指针类型。接受者中的变量
在命名时，官方建议使用接受者类型的第一个小写字母。

```golang
package main

import (
	"fmt"
)

type Employee struct {
	name, currency string
	salary         float64
}

func main() {
	emp1 := Employee{"Daniel", "$", 2000}
	emp1.printSalary()
	printSalary(emp1)
}

//printSalary()方法
func (e Employee) printSalary() {
	fmt.Printf("员工姓名：%s，薪资：%s%.2f \n", e.name, e.currency, e.salary)
}

//printSalary函数
func (e Employee) printSalary {
	fmt.Printf("员工姓名：%s，薪资：%s%.2f \n", e.name, e.currency, e.salary)
}
```

> 若方法的接受者不是指针，实际只是获取了一个拷贝，而不能真正改变接受者中原来的数据。

**方法继承**
> 如果匿名字段实现了一个方法，那么包含这个匿名字段的struct也能调用该匿名字段中的方法。
```golang
package main

import (
	"fmt"
)

type Human struct {
	name, phone string
	age         int
}

type Student struct {
	Human  //匿名字段
	school string
}

type Employee struct {
	Human   //匿名字段
	company string
}

func main() {
	s1 := Student{Human{"Daniel", "15012345***", 13}, "十一中学"}
	e1 := Employee{Human{"Steven", "17812345***", 35}, "1000phone"}
	s1.SayHi()
	e1.SayHi()
}
func (h *Human) SayHi() {
	fmt.Printf("大家好！我是%s,我%d岁，我的联系方式是：%s\n", h.name, h.age, h.phone)
}
```

**方法重写**
```golang
package main

import (
	"fmt"
)

type Human struct {
	name, phone string
	age         int
}

type Student struct {
	Human  //匿名字段
	school string
}

type Employee struct {
	Human   //匿名字段
	company string
}

func main() {
	s1 := Student{Human{"Daniel", "15012345***", 13}, "十一中学"}
	e1 := Employee{Human{"Steven", "17812345***", 35}, "1000phone"}

	s1.SayHi()
	e1.SayHi()
}
func (h *Human) SayHi() {
	fmt.Printf("大家好！我是%s,我%d岁，联系方式：%s\n", h.name, h.age, h.phone)
}

func (s *Student) SayHi() {
	fmt.Printf("大家好！我是%s,我%d岁，我在%s上学，联系方式：%s\n", s.name, s.age, s.school, s.phone)
}

func (n *Employee) SayHi() {
	fmt.Printf("大家好！我是%s,我%d岁，我在%s工作，联系方式：%s\n", n.name, n.age, n.company, n.phone)
}
```

### switch
```txt
switch语句    switch语句用于基于不同条件执行不同动作
select      select语句类似于switch语句，但是select会随机执行一个可运行的case。如果没有case可运行，
        它将阻塞，直到有case可运行。
```

Go语言中switch语句的语法如下所示：
```golang
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
```

> switch语句执行的过程自上而下，直到找到case匹配项，匹配项中无须使用break，因为Go语言
中的switch默认给每个case自带break。因此匹配成功后不会向下执行其他的case分支，而是跳出
整个switch。可以添加fallthrough，强制执行后面的case分支。fallthrough必须放在case分支的
最后一行。如果它出现在中间的某个地方，编译器就会报错。

> 变量var1可以是任何类型，而val1和val2则可以是同类型的任意值。类型不局限于常量或整数，但
必须是相同类型或最终结果为相同类型的表达式。

> case后的值不能重复，但可以同时测试多个符合条件的值，也就是说case后可以有多个值，这些值
之间使用逗号分隔，例如：case val1,val2,val3。

> switch后的表达式可以省略，默认是switch true

```golang
/* 定义局部变量 */
grade := ""
score := 78.5
switch { //switch后面省略不写，默认相当于：switch true
case score >= 90:
    grade = "A"
case score >= 80:
    grade = "B"
case score >= 70:
    grade = "C"
case score >= 60:
    grade = "D"
default:
    grade = "E"
}

switch grade {
case "A":
    fmt.Printf("优秀!\n")
case "B":
    fmt.Printf("良好\n")
case "C":
    fmt.Printf("中等\n")
case "D":
    fmt.Printf("及格\n")
default:
    fmt.Printf("差\n")
}
/////////////////////////////
/* 定义局部变量 */
year := 2008
month := 2
days := 0
switch month {
case 1, 3, 5, 7, 8, 10, 12:
    days = 31
case 4, 6, 9, 11:
    days = 30
case 2:
    if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
        days = 29
    } else {
        days = 28
    }
default:
    days = -1
}
fmt.Printf("%d 年 %d 月的天数为: %d\n", year, month, days)
```

**类型转换**
> switch语句还可以被用于type switch(类型转换)来判断某个interface变量中实际存储的变量类型。
下面演示type switch的语法。其语法结构如下所示：

```golang
switch x.(type){
    case type:
         statement(s);
    case type:
         statement(s);
    /* 你可以定义任意个数的case */
    default: /* 可选 */
        statement(s);
```

```golang
// 判断interface变量中存储的变量类型。
var x interface{}
switch i := x.(type) {
case nil:
    fmt.Printf(" x 的类型:%T", i)
case int:
    fmt.Printf("x是int型")
case float64:
    fmt.Printf("x是float64型")
case func(int):
    fmt.Printf("x是func(int)型")
case bool, string:
    fmt.Printf("x是bool或string型")
default:
    fmt.Printf("未知型")
}
```

### select
> http://tour.studygolang.com/concurrency/5

```text
select 语句使一个 Go 程可以等待多个通信操作。

select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。
```
### type
**类型别名**
```text
类型别名的语法格式如下。
type 类型别名 = 类型
```

```golang
type byte = uint8
type rune = int32
```

**类型定义**
```text
定义类型的语法格式如下。
type 新的类型 类型
```

```golang
type NewString string
```

> 该语句是将NewString定义为string类型。通过type关键字，NewString会形成一种新的类型。
NewString本身依然具备string的特性。

>出于对程序性能的考虑，建议如下。
- 尽可能地使用：=去初始化声明一个变量(在函数内部)。
- 尽可能地使用字符代替字符串

## 数据类型

### 复合数据类型

#### array
> 数组是相同类型的一组数据构成的长度固定的序列，其中数据类型包含了基本数据类型、符合数据类型和自定义数据类型。
数组中的每一项被称为数组的元素。数组名是数组的唯一标识符，数组的每一个元素都是没有名字的，只能通过索引下标
(位置)进行访问。因为数组的内存是一段连续的存储区域，所以数组的检索速度是非常快的，但是数组也有一定的缺陷，就是定义
后长度不能更改。

**声明**
```text
var 变量名 [数组长度] 数据类型
```

**多维数组**
> 由于数据的复杂程度不一样，数组可能有多个下表。一般将数组元素下标的个数称为维数，根据维数，可将数组分为一维数组、
二维数组、三维数组、四维数组等。二维及以上的数组可称为多维数组。
```text
var variable_name [SIZE1][SIZE2]...[SIZEn] variable_type
```

**如何判断元素是否在切片中?**

在Python，可以使用in和not in语法做判断，但Go不行，只能遍历检查。 查资料浏览网友的说法，如果需要判断元素是否存在，建议使用map。

有一个例子，检查某个值是否存在于环境环境(这个是[]string类型)，使用了os.Getenv(key)。 查看源码，了解到最终是调用了windows的dll程序来获取结果。

**如何判断两个切片是否一致?**

使用reflect.DeepEqual函数

```text
二维数组的定义方式如下。
var arrayName [x][y] variable_type
```

#### interface
> 接口是一组方法签名。接口指定了类型应该具有的方法，类型决定了如果实现这些方法。当某个类型行为接口中的所有方法提供了具体
的实现细节时，这个类型就被称为实现了该接口。接口定义了一组方法，如果某个对象实现了该接口的所有方法，则此对象就实现了该接口。

> Go语言的类型都是隐式实现接口的。任何定义了接口中所有的方法的类型都被称为隐式地实现了该接口。

定义接口的语法格式如下。
```text
type 接口名字 interface{
方法1([参数列表]) [返回值]
方法2([参数列表]) [返回值]
...
方法n([参数列表]) [返回值]
}

实现该接口方法的语法格式如下：
func (变量名 结构体类型) 方法1([参数列表]) [返回值]{
    //方法体
}
func (变量名 结构体类型) 方法2([参数列表]) [返回值]{
    //方法体
}
...
func (变量名 结构体类型) 方法n([参数列表]) [返回值]{
    //方法体
}
```

```golang
package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type AndroidPhone struct {
}
type IPhone struct {
}

func (a AndroidPhone) call() {
	fmt.Println("我是安卓手机，可以打电话了")
}

func (i IPhone) call() {
	fmt.Println("我是苹果手机，可以打电话了")
}

func main() {
	//定义接口类型的变量
	var phone Phone
	phone = new(AndroidPhone)
	fmt.Printf("%T,%v,%p \n", phone, phone, &phone)
	phone.call()
	phone = AndroidPhone{}
	fmt.Printf("%T,%v,%p \n", phone, phone, &phone)
	phone.call()
	phone = new(IPhone)
	fmt.Printf("%T,%v,%p \n", phone, phone, &phone)
	phone.call()
	phone = IPhone{}
	fmt.Printf("%T,%v,%p \n", phone, phone, &phone)
	phone.call()
}
```

**关于duck typing**
- duck typing是描述事物的外部行为而非内部结构。
- duck typing关注的不是对象的类型本身，而是它是如何使用的。

> 代码实现请参阅https://gitee.com/luzhenxiong/codes/hyds70ploebkifv9654z292

**空接口**
> 空接口中没有任何方法。任意类型都可以实现该接口。空接口这样定义：interface{}，也就是包含0个方法(method)的interface。
> 空接口可表示任意数据类型，类似于Java中的object。

空接口常用于以下情形。
- println的参数就是空接口。
- 定义个map：key是string，value是任意数据类型。
- 定义一个切片，其中存储任意类型的数据。
```golang
package main

import (
	"fmt"
)

type A interface {
}

type Cat struct {
	name string
	age  int
}

type Person struct {
	name string
	sex  string
}

func main() {
	var a1 A = Cat{"Mimi", 1}
	var a2 A = Person{"Steven", ""}
	var a3 A = "Learn golang with me!"
	var a4 A = 100
	var a5 A = 3.14
	fmt.Println("-----------------")
	//1.fmt.println参数就是空接口
	fmt.Println("println的参数就是空接口，可以是任何数据类型", 100, 3.14, Cat{"旺旺", 2})
	//2.定义map,value是任何数据类型
	map1 := make(map[string]interface{})
	map1["name"] = "Daniel"
	map1["age"] = 13
	map1["height"] = 1.71
	fmt.Println(map1)
	fmt.Println("-----------------")
	//3.定义一个切片，其中存储任意数据类型
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, a1, a2, a3, a4, a5)
	fmt.Println(slice1)
}

```

接口对象类型
```golang
/*
instance,ok:=接口对象.(实际类型)

如果该接口对象是对应的实际类型，那么instance就是转型之后的对象，ok的值为true，配合
if...else if...语句使用。

接口对象转型第二种方式示例如下。
接口对象.(type)
*/
//样例：接口对象转型。
package main

import (
	"fmt"
	"math"
)

//1.定义接口
type Shape interface {
	perimeter() float64
	area() float64
}

//2.矩形
type Rectangle struct {
	a, b float64
}

//3.三角形
type Triangle struct {
	a, b, c float64
}

//4.圆形
type Circle struct {
	radius float64
}

//定义实现接口的方法
func (r Rectangle) perimeter() float64 {
	return (r.a + r.b) * 2
}
func (r Rectangle) area() float64 {
	return r.a * r.b
}
func (t Triangle) perimeter() float64 {
	return t.a + t.b + t.c
}
func (t Triangle) area() float64 {
	//海伦公式
	p := t.perimeter() / 2 //半周长
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

//接口对象转型方式1
//instance,ok:=接口对象.(实际类型)
func getType(s Shape) {
	if instance, ok := s.(Rectangle); ok {
		fmt.Printf("矩形：长度%.2f,宽度%.2f,", instance.a, instance.b)
	} else if instance, ok := s.(Triangle); ok {
		fmt.Printf("三角形：三边分别：%.2f,%.2f,%.2f,", instance.a, instance.b, instance.c)
	} else if instance, ok := s.(Circle); ok {
		fmt.Printf("圆形：半径%.2f,", instance.radius)
	}
}
//接口对象.(type)，配合switch和case语句使用
func getType2(s Shape) {
	switch instance := s.(type) {
	case Rectangle:
		fmt.Printf("矩形：长度为%.2f，宽为%.2f,\t", instance.a, instance.b)
	case Triangle:
		fmt.Printf("三角形：三边分别为%.2f,%.2f,%.2f,\t", instance.a, instance.b, instance.c)
	case Circle:
		fmt.Printf("圆形：半径为%.2f,\t", instance.radius)
	}
}

func getResult(s Shape) {
	getType2(s)
	fmt.Printf("周长：%.2f，面积:%.2f \n", s.perimeter(), s.area())
}
func main() {
	var s Shape
	s = Rectangle{3, 4}
	getResult(s)
	showInfo(s)
	s = Triangle{3, 4, 5}
	getResult(s)
	showInfo(s)
	s = Circle{1}
	getResult(s)
	showInfo(s)
	x := Triangle{3, 4, 5}
	fmt.Println(x)
}

func (t Triangle) String() string { //实现了系统接口，最后的打印部分会改变
	return fmt.Sprintf("Triangle对象，属性分别为：%.2f,%.2f,%.2f", t.a, t.b, t.c)
}

func showInfo(s Shape) {
	fmt.Printf("%T,%v \n", s, s)
	fmt.Println("----------------")
}
```

### 类型转换
```text
1.基本语法
Go语言采用数据类型前置加括号的方式进行类型转换，格式如：T(表达式)。T表示要转换的类型；表达式包括变量、
数值、函数返回值等。

类型转换时，需要考虑两种类型之间的关系和范围，是否会发生数值截断。值得注意的是，布尔型无法与其他类型
进行转换。

使用示例如下：
var a int = 100
b:=float64(a)  //将int型转换成float64型
c:=string(a)    //将int型转换成string型

2.浮点型与整型之间转换
float和int的类型精度不同，使用时需要注意float转int时精度的损失。
具体的使用方法：
package main

import "fmt"

func main() {
	chinese := 90
	english := 80.9
	avg := (chinese + int(english)) / 2
	avg2 := (float64(chinese) + english) / 2
	fmt.Printf("%T,%d\n", avg, avg)
	fmt.Printf("%T, %f\n", avg2, avg2)
}

运行结果：
int,85
float64, 85.450000

3.整型转字符串类型
这种类型的转换，其实相当于byte或rune转string。int数值是ASCII码的编号或unicode字符集的编号，转成
string就是根据字符集，将对应编号的字符查出来。当该数值超出unicode编号范围，则转成的字符串显示为
乱码。例如，19968转string，就是'一'。
备注：
·ASCII字符集中数字的十进制范围是48~57;
·ASCII字符集中大写字母的十进制范围是65~90；
·ASCII字符集中小写字母的十进制范围是97~122；
·unicode字符集中汉字的范围是4e00~9fa5，十进制范围是19968~40869。

具体的使用方法：
package main

import "fmt"

func main() {
	a := 97
	x := 19968
	result := string(a)
	fmt.Println(result)
	result = string(x)
	fmt.Println(result)
}

运行结果：
a
一

注意：在Go语言中，不允许字符串转int。
```

## 运算符
### 算数运算符
```text
运算符	描述	
+	相加	
-	相减
*	相减
/	相除
%	求余
++	自增
--	自减
```
### 关系运算符
```text
运算符：
==
!=
>
<
>=
<=
```

### 逻辑运算符
```text
运算符	描述
&&	逻辑AND运算符。如果两边的操作数都是True，则条件为True，否则为False
||	逻辑OR运算符。如果两边的操作数有一个True，则条件True，否则为False
!	逻辑NOT运算符。如果条件为True，则逻辑NOT条件False，否则为True。
```

### 位运算符
```text

位运算符对整数在内存中的二进制位进行操作。

位运算符比一般的算术运算符速度要快，而且可以实现一些算术运算符不能实现的功能。如果要开发高效率程序，
位运算符是必不可少的。位运算符用来对二进制位进行操作，包括：按位与(&)、按位或(|)、按位异或(^)、
按位左移(<<)、按位右移(>>)。

位运算符
运算符	描述					实例
&	按位运算符"&"是双目运算符。			(A & B)结果为12，二进制位0000 1100
	其功能是参与运算的两数各对应的二进制位与
|	按位或运算符"|"是双目运算符。其功能是参与	(A|B)结果位61，二进制位0011 1101
	的两数各对应的二进制位或
^	按位异或运算符"^"是双目运算符。其功能是		(A^B)结果为49，二进制位0011 0001
	参与运算的两数各对应的二进制位巷异或，当两
	对应的二进制位异或时，结果位1
<<	左移运算符"<<"是双目运算符。左移n位就是乘	(A<<2)结果为240，二进制为1111 0000
	以2的n次方。其功能是把"<<"左边的运算数的各
	二进制位全部左移若干位，由"<<"右边的数指定
	移动的位数，高位舍弃，低位补0
>>>	右移运算符">>"是双目运算符。右移n位就是除以2	(A>>2)结果位15，二进制为0000 1111
	的n次方。其功能是把">>"左边的运算数的各二进
	制全部右移若干位，">>"右边的数指定移动的位数

1.按位与
按位与(&)：对两个数进行操作，然后返回一个新的数，这个数的每一位都需要两个输入数的同一位都为1时才为
1.简单说就是：同一位同时为1则为1。

2.按位
按位或(|):对两个数进行操作，然后返回一个新的数，这个数的每一位只要注意任意一个输入数的同一位位1则位1。
简单说就是，同一位其中一个位1。

3.按位异或
按位异或(^)：对两个数进行操作，然后返回一个新的数，这个数的每一位只要两个输入数的同一位不同则为1，如果相同
就为0。简单说就是，同一位不相同则位1。

4.左移运算符(<<)
按二进制形式把所有的数字向左移动对应的位数，高位移出(舍弃)，低位的空位补0.
（1）语法格式
需要移位的数字<<移位的次数

例如，3<<4，则是将数字3左移4位。

（2）计算过程
3<<4

(3)数学意义
在数字没有溢出的前提下，对于正数和负数，左移1位都相当于乘以2的1次方，左移n位就相当于乘以2的n次方。
```

### 赋值运算符
```text
运算符	描述
=	简单的赋值运算符，将一个表达式的值赋给一个左值
+=	相加后再赋值
-=	相减后再赋值
*=	相乘后再赋值
/=	相除后再赋值
%=	求余后再赋值
<<=	左移后再赋值
>>=	右移后赋值
&=	按位与后赋值
^=	按位异或后赋值
|=	按位或后赋值
```

### 指针运算符
```text
运算符	描述	
&	返回变量存储地址	
*	指针变量
```

### 运算符优先级
```text
优先级	运算符
7	^!
6	*/%<<>>&&^
5	+-|^
4	==!=<<=>=>
3	<-
2	&&
1	||
```

## 指针
> 指针是存储另一个变量的内存地址的变量。

> 例如：变量b的值为156，存储在内存地址0x1040a124。变量a持有b的地址，则a被认为指向b。

> 在Go语言中使用取地址符(&)来获取变量的地址，一个变量前使用&，会返回该变量的内存地址。

> Go语言指针的最大特点是：指针不能运算(不同于C语言)

### 声明指针
> 声明指针，*T是指针变量的类型，它指向T类型的值。
```text
var 指针变量名 *指针类型
*号用于指定变量是一个指针。
```

```golang
var ip *int	//指向整型的指针
var fp *float32 //指向浮点型的指针
```

指针使用流程如下：
- 定义指针变量
- 为指针变量赋值。
- 访问指针变量中指向地址的值。

> 获取指针指向的变量值：在指针类型的变量前加上*号，如*a。

```golang
//声明实际变量
var a int = 120
//声明指针变量
var ip *int
//给指针变量赋值，将变量a的地址赋给ip
ip = &a
//打印ip的类型和值
fmt.Printf("ip的类型是%T，值是%v \n", ip, ip)
//打印变量*ip的类型和值
fmt.Printf("*ip变量的类型是%T，值是%v \n", *ip, *ip)
```

```golang
var s1 = Student{"Steven", 35, true, 1}
var s2 = Student{"Sunny", 20, false, 0}
var a *Student = &s1 //将s1的内存地址赋值给Student指针变量a
var b *Student = &s2 //将s2的内存地址赋值给Student指针变量b
fmt.Println(s1.name, s1.age, s1.married, s1.sex)
fmt.Println(a.name, a.age, a.married, a.sex)
fmt.Println(s2.name, s2.age, s2.married, s2.sex)
fmt.Println(b.name, b.age, b.married, b.sex)
fmt.Println((*a).name, (*a).age, (*a).married, (*a).sex)
fmt.Println((*b).name, (*b).age, (*b).married, (*b).sex)
fmt.Println(&a.name, &a.age, &a.married, &a.sex)
fmt.Println(&b.name, &b.age, &b.married, &b.sex)
```

### 空指针
> 当一个指针被定义后没有分配到任何变量时，它的值为nil。nil指针也成为空指针。nil在概念上和其他语言的 null、None、NULL一样，都指代零值或空值。

假设指针变量命名为ptr。空指针判断如下。
```golang
if(ptr !=nil)  //ptr不是空指针
if(ptr==nil) //ptr是空指针
```

### 使用指针
**通过指针修改变量的值**
```golang
b := 3158
a := &b
*a++
fmt.Println("b的新值：", b)
```

### 指针的指针
> 如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。当定义一个指向指针的指针变量
时，第一个指针存放第第二个指针的地址，第二个指针存放变量的地址。
```text
指向指针的指针变量声明格式如下。
var ptr **int

以上指向指针的指针变量为整型。
访问指向指针的指针变量值需要使用两个*号。
```
```golang
var a int
var ptr *int
var pptr **int
a = 1234
/* 指针ptr地址 */
ptr = &a
fmt.Println("ptr", ptr)
/*指向指针ptr的地址*/
pptr = &ptr
fmt.Println("pptr", ptr)
/*获取pptr的值*/
fmt.Printf("变量a=%d\n", a)
fmt.Printf("指针变量 *ptr=%d\n", *ptr)
fmt.Printf("指向指针的指针变量**ptr=%d\n", **pptr)
```

# 第二部分 提升
## 编译
```commandline
go build hello.go
```
### Windows编译Linux可执行文件
```commandline
set GOARCH=amd64
set GOOS=linux
go build
```
> 完成后记得设置回来set GOOS=windows

_注意_
```text
部分情况下交叉编译可能失败，原因是某些代码引用的包只能在特定操作系统上使用
```
> 解决办法：尽量统一开发环境，提前发现问题

### 编译为DLL文件
> 注意函数的注释:导出的间距是特定的,在//和export文本之间不应该有空格:
> 好像不能加注释

```commandline
go build -buildmode=c-shared -o mylib.dll file.go
```
> 编译后得到一个file.dll和file.h两个文件

#### 返回字符串
```golang
//export ReturnString
func ReturnString(val string) *C.char {
	cname, err := net.LookupCNAME(val)
	if err != nil {
		C.CString("Could not find CNAME")
	}
	return C.CString(cname)
}
```
#### 返回数值
```golang
//export ReturnInt
func ReturnInt(val int) int {
	return val + 3
}
```

#### 传入数组
```golang
//export Example
func Example(cArray *C.int, cSize C.int, i C.int) C.int {
    gSlice := (*[1 << 30]C.int)(unsafe.Pointer(cArray))[:cSize:cSize]
    return gSlice[i]
}
```

# 最佳实践
## 日志
> 通过标准包log实现。

> 包提供了一个预定义的“标准”Logger，可以通过辅助函数Print[f|ln]、Fatal[f|ln]和Panic[f|ln]访问。
> 和fmt包不同的是，日志开头会带上时间信息。

> log没有提供日志级别，可以通过prefix属性区分不同等级日志，另外的方法是使用第三方日志框架。

设置日志格式:
```golang
log.SetFlags(log.Llongfile| log.LstdFlags)
```
```golang
const (
    // 字位共同控制输出日志信息的细节。不能控制输出的顺序和格式。
    // 在所有项目后会有一个冒号：2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒分辨率：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件无路径名+行号：d.go:23（会覆盖掉Llongfile）
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
)
```

### 写入到文件
```golang
logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
if err != nil {
	panic(err)
}
logger = log.New(logFile, "[MAIN]", log.LstdFlags|log.Lshortfile) // 将文件设置为loger作为输出
```

### Fatal()和Panic()的区别
> 原文链接https://www.jianshu.com/p/f85ecae6e7df
- Fatal()的动作
    - 打印输出内容
    - 退出应用程序
    - defer函数不会执行
- Panic()的动作
    - 函数立刻停止执行 (注意是函数本身，不是应用程序停止)
    - defer函数被执行
    - 返回给调用者(caller)
    - 调用者函数假装也收到了一个panic函数，从而
        - 立即停止执行当前函数
        - 它defer函数被执行
        - 返回给它的调用者(caller)
    - (递归重复上述步骤，直到最上层函数)
    - 应用程序停止。
    - panic的行为

## 单元测试
> 标准库testing 提供对 Go 包的自动化测试的支持。
### 基准测试
```text
基准测试，是一种测试代码性能的方法，比如你有多种不同的方案，都可以解决问题，那么到底是那种方案性能更好呢？这时候基准测试就派上用场了。

基准测试主要是通过测试CPU和内存的效率问题，来评估被测试代码的性能，进而找到更好的解决方案。比如链接池的数量不是越多越好，那么哪个值才是最优值呢，
这就需要配合基准测试不断调优了。
```

**不错的资料**: https://www.cnblogs.com/saryli/p/11406937.html

### gomock
> https://gitee.com/luzhenxiong/devdoc/blob/master/%E5%90%8E%E7%AB%AF/awesome/go/gomock

# 常见问题
## 执行路径问题
> https://zhuanlan.zhihu.com/p/363714760
```text
Python很容易确定该执行文件所在目录，但Go在go run下执行和go build下执行程序，获取到的执行路径不一定一致！
```

# 优秀源码案例学习
- chromedp
    - 设置默认值
      ```text
      chromedp/allocate.go(line53): DefaultExecAllocatorOptions
      ```

      ```golang
      var DefaultExecAllocatorOptions = [...]ExecAllocatorOption{
        NoFirstRun,
        NoDefaultBrowserCheck,
        Headless,
    
        // After Puppeteer's default behavior.
        Flag("disable-background-networking", true),
        Flag("enable-features", "NetworkService,NetworkServiceInProcess"),
        Flag("disable-background-timer-throttling", true),
        Flag("disable-backgrounding-occluded-windows", true),
        Flag("disable-breakpad", true),
        Flag("disable-client-side-phishing-detection", true),
        Flag("disable-default-apps", true),
        Flag("disable-dev-shm-usage", true),
        Flag("disable-extensions", true),
        Flag("disable-features", "site-per-process,TranslateUI,BlinkGenPropertyTrees"),
        Flag("disable-hang-monitor", true),
        Flag("disable-ipc-flooding-protection", true),
        Flag("disable-popup-blocking", true),
        Flag("disable-prompt-on-repost", true),
        Flag("disable-renderer-backgrounding", true),
        Flag("disable-sync", true),
        Flag("force-color-profile", "srgb"),
        Flag("metrics-recording-only", true),
        Flag("safebrowsing-disable-auto-update", true),
        Flag("enable-automation", true),
        Flag("password-store", "basic"),
        Flag("use-mock-keychain", true),
  }
  ```
    - 初始化默认值参数
    ```text
    chromedp/allocate.go(line37): setupExecAllocator
    ```

    - 修改默认值参数
      ```text
      chromedp/allocate.go(line85): NewExecAllocator的opts可选参数
      ```