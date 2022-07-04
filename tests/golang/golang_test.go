/*
官网: https://golang.google.cn/

Go语言有l
*/

package golang

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

// TestEnv https://blog.csdn.net/weixin_42871822/article/details/117096038
func TestEnv(t *testing.T) {
	// 更换源: go env -w GOPROXY=https://goproxy.cn,direct
	output, err := exec.Command("go", "env", "GOPROXY").Output()
	assert.Nil(t, err)
	assert.Equal(t, "https://goproxy.cn,direct\n", string(output))
}

// TestInit 初始化go项目
func TestInit(t *testing.T) {
	// go mod init repo
	dir := t.TempDir()
	os.Chdir(dir)
	_, err := exec.Command("go", "mod", "init", "repo").Output()
	if assert.Nil(t, err) == false {
		t.Fatal(err.Error())
	}
}

// TestVar 常见变量声明形式
// 未初始化变量的默认值有如下特点：
// - 整形和浮点型变量默认值：0
// - 字符串默认值未空字符串
// - 布尔型默认值未false
// - 函数、指针变量、切片默认值为nil
func TestVar(t *testing.T) {
	var a int32
	var i = 13
	var f = float32(3.14)
	// 短变量声明方式
	n := 17
	var (
		crlf       = []byte("\r\n")
		colonSpace = []byte(": ")
	)
	assert.Equal(t, a, int32(0))
	assert.Equal(t, i, 13)
	assert.Equal(t, f, float32(3.14))
	assert.Equal(t, n, 17)
	assert.Equal(t, string(crlf), "\r\n")
	assert.Equal(t, string(colonSpace), ": ")

	// 匿名变量, 不占用命名空间，也不会分配内存
	_ = a
}

// TestSwitchVar 交换变量
func TestSwitchVar(t *testing.T) {
	var a = 10
	var b = 20
	a = a ^ b
	b = b ^ a
	a = a ^ b
	assert.Equal(t, a, 20)
	assert.Equal(t, b, 10)

	// 第二种方法(推荐)
	var c = 10
	var d = 20
	c, d = d, c
	assert.Equal(t, c, 20)
	assert.Equal(t, d, 10)
}

// TestConst
func TestConst(t *testing.T) {
	// 显式类型
	const B string = "Steven"
	// 隐式类型
	const C = "Steven"
	// 多个相同类型的声明可以简写如下
	const WIDTH, HEIGHT = 1980, 1080
	// 使用常量组模拟枚举
	const (
		Unkonwn = 0
		Female  = 1
		Male    = 2
	)
	//常量组如果不指定类型和初始值，则与上一行费空常量的值相同
	const (
		a = 10
		b
		c
	)
	// iota可以被用作枚举值, 特殊常量值，是一个系统定义的可以被编译器修改的常量值。iota只能被用在常量的赋值中，在每一个
	// const关键字出现时，被重置为0，然后每出现一个常量，iota所代表的数值会自动加1。不论该常量的值是什么，
	// 只要有一个，那么iota就加1。
	const (
		e = iota
		f
		g
	)

	assert.Equal(t, B, C)
	assert.Greater(t, WIDTH, HEIGHT)
	assert.Equal(t, Unkonwn, 0)
	assert.Equal(t, Female, 1)
	assert.Equal(t, Male, 2)
	assert.Equal(t, a, b)
	assert.Equal(t, b, c)
	assert.Equal(t, e, 0)
	assert.Equal(t, f, 1)
	assert.Equal(t, g, 2)
}

// TestSlice 切片
func TestSlice(t *testing.T) {
	// 动态扩容:
	// append会根据切片的需要, 在当前底层数组容量无法满足的情况下, 动态分配新的数组,
	// 新数组长度会按一定算法扩展。
	// 新数组简历h
	var s []int
	s = append(s, 11)
	assert.Equal(t, len(s), 1)
	assert.Equal(t, cap(s), 1)
	s = append(s, 12)
	assert.Equal(t, len(s), 2)
	assert.Equal(t, cap(s), 2)
	s = append(s, 13)
	assert.Equal(t, len(s), 3)
	assert.Equal(t, cap(s), 4)
	s = append(s, 14)
	assert.Equal(t, len(s), 4)
	assert.Equal(t, cap(s), 4)
	s = append(s, 15)
	assert.Equal(t, len(s), 5)
	assert.Equal(t, cap(s), 8)
}

// 连接字符串的基准测试 go test -bench=. -benchmem ./golang_test.go =======================================================
// 测试结果:
// 做了预初始化的strings.Builder连接的构建字符串效率最高;
// 带有预初始化的bytes.Buffer和strings.Join效率相近
// 未做预初始化的strings.Builder、bytes.Buffer和操作符连接排中间
// fmt.Sprintf性能最差, 排在末尾
var s1 = []string{
	"Rob Pike ",
	"Robert Griesemer ",
	"Ken Thompson ",
}

func concatStringByOperator(s1 []string) string {
	var s string
	for _, v := range s1 {
		s += v
	}
	return s
}

func concatStringBySprintf(s1 []string) string {
	var s string
	for _, v := range s1 {
		s = fmt.Sprintf("%s%s", s, v)
	}
	return s
}

func concatStringByJoin(s1 []string) string {
	return strings.Join(s1, "")
}

func concatStringByStringsBuilder(s1 []string) string {
	var b strings.Builder
	for _, v := range s1 {
		b.WriteString(v)
	}
	return b.String()
}

func concatStringByStringsBuilderWithInitSize(s1 []string) string {
	var b strings.Builder
	b.Grow(64)
	for _, v := range s1 {
		b.WriteString(v)
	}
	return b.String()
}

func concatStringByBytesBuffer(s1 []string) string {
	var b bytes.Buffer
	for _, v := range s1 {
		b.WriteString(v)
	}
	return b.String()
}

func concatStringByBytesBufferWithInitSize(s1 []string) string {
	buf := make([]byte, 0, 64)
	b := bytes.NewBuffer(buf)
	for _, v := range s1 {
		b.WriteString(v)
	}
	return b.String()
}

func BenchmarkConcatStringByOperator(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByOperator(s1)
	}
}

func BenchmarkConcatStringBySprintf(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringBySprintf(s1)
	}
}

func BenchmarkConcatStringByJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByJoin(s1)
	}
}

func BenchmarkConcatStringByStringsBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByStringsBuilder(s1)
	}
}

func BenchmarkConcatStringByStringsBuilderWithInitSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByStringsBuilderWithInitSize(s1)
	}
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByBytesBuffer(s1)
	}
}

func BenchmarkConcatStringByBytesBufferWithInitSize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concatStringByBytesBufferWithInitSize(s1)
	}
}

// =====================================================================================================================

var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func f() int {
	d++
	return d
}

// TestEvaluationOrder 包级别变量声明语句中的表达式求值顺序
// 如果某个变量(如变量a)的初始化表达式中直接或间接依赖其他变量(如变量b), 那么a的初始化顺序排在变量b后面
func TestEvaluationOrder(t *testing.T) {
	assert.Equal(t, a, 9)
	assert.Equal(t, b, 4)
	assert.Equal(t, c, 5)
	assert.Equal(t, d, 5)
}

// TestFor for循环
// for是Go语言中唯一的循环语句，Go没有while、do…while循环。
func TestFor(t *testing.T) {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d", i)
	}

	// 省略初始语句
	i := 0
	for ; i <= 10; i++ {
		fmt.Printf("%d", i)
	}

	// 省略条件表达式
	i = 0
	for ; ; i++ {
		if i > 20 { //使用break跳出循环
			break
		}
		fmt.Printf("%d", i)
	}

	// 类while循环
	for i <= 10 {
		fmt.Print(i)
		i++
	}

	// 无表达式
	for {
		if i > 10 {
			break
		}
		fmt.Print(i)
		i++
	}

	//遍历字符串
	str := "字符串"
	for i, value := range str {
		fmt.Printf("第%d的ASCII值=%d，字符是%c \n", i, value, value)
	}
}

// TestIf
// if语句        if语句由一个布尔表达式后紧跟一个或多个语句组成
// if...else语句 if语句后可以使用可选的else语句，else语句中的表达式再布尔表达式为false时执行
// if嵌套语句     可以在if或else if语句中潜入一个或多个if或else if语句
func TestIf(t *testing.T) {
	num := rand.Int()
	if num%2 == 0 {
		fmt.Println(num, "偶数")
	} else {
		fmt.Println(num, "奇数")
	}

	score := rand.Int31()
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

	if num := rand.Int(); num%2 == 0 {
		fmt.Println(num, "偶数")
	} else {
		fmt.Println(num, "奇数")
	}
}

// TestSwitch switch语句用于基于不同条件执行不同动作。
// select语句类似于switch语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
// switch语句执行的过程自上而下，直到找到case匹配项，匹配项中无须使用break，因为Go语言 中的switch默认给每个case自带break。因此匹配成功后不会向下执行其他的case分支，而是跳出 整个switch。可以添加fallthrough，强制执行后面的case分支。fallthrough必须放在case分支的 最后一行。如果它出现在中间的某个地方，编译器就会报错。
// 变量var1可以是任何类型，而val1和val2则可以是同类型的任意值。类型不局限于常量或整数，但 必须是相同类型或最终结果为相同类型的表达式。
// case后的值不能重复，但可以同时测试多个符合条件的值，也就是说case后可以有多个值，这些值 之间使用逗号分隔，例如：case val1,val2,val3。
// switch后的表达式可以省略，默认是switch true:
func TestSwitch(t *testing.T) {
	/* 定义局部变量 */
	grade := ""
	score := rand.Int()
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
	year := rand.Int()
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
}

func TestError(t *testing.T) {
	// 返回自定义异常
	// 第一种方法(推荐)
	fmt.Println(errors.New("自定义error"))
	// 第二种方法
	fmt.Println(fmt.Errorf("%s", "error message"))

	// 定义全局异常
	var ValueError error
	ValueError = errors.New("error message")
	fmt.Println(ValueError)
}

// TestString 字符串
func TestString(t *testing.T) {
	s := "世界你好HelloWorld"
	// 截取字符串
	// string类型中，一个字符串占3~4个字节，因此使用len和截取字符串是不准确的， 需要先转换为rune类型再处理，例如截取 s 变量的前4位:

	assert.Equal(t, string([]rune(s)[:4]), "世界你好")
	// 判断是否包含子串
	assert.True(t, strings.Contains("seafood", "foo"))
	// 判断字符串是否包含另一字符串中的任一字符
	assert.True(t, strings.ContainsAny("team", "ie"))
	// 判断开头的字符串
	assert.True(t, strings.HasPrefix(s, "世界你好"))
	// 判断结尾的字符串
	assert.True(t, strings.HasSuffix(s, "World"))
	// 返回字符串中另一字符串首次出现的位置
	assert.Equal(t, strings.Index("chicken", "ken"), 4)
	// 将字符串以空白字符分割, 并返回一个切片
	assert.Equal(t, strings.Fields("abc 123 ABC xyz XYZ"), []string{"abc", "123", "ABC", "xyz", "XYZ"})
}

// TestInt 整型
func TestInt(t *testing.T) {
	// int转string
	var num int = 100
	assert.Equal(t, strconv.Itoa(num), "100")
}

// TestMap
func TestMap(t *testing.T) {
	// 列出所有Key
	dict := make(map[string]interface{})
	dict["name"] = "Daniel"
	dict["age"] = 13
	dict["height"] = 1.71

	list := make([]string, 0, len(dict))
	for key, _ := range dict {
		list = append(list, key)
	}
	assert.Contains(t, list, "name")
	assert.Contains(t, list, "age")
	assert.Contains(t, list, "height")

	// 判断键名是否存在字典中
	if _, ok := dict["USA"]; ok {
		t.Fatal("fail")
	}

	// 删除元素
	if _, ok := dict["background"]; ok {
		delete(dict, "background")
	}
}
