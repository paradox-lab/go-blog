/*
官网: https://golang.google.cn/

TODO 处理掉vuepress跟go有关的md文件, 然后再看精进之路的书
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
}
