/*
官网: https://golang.google.cn/
*/

package golang

import (
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
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
