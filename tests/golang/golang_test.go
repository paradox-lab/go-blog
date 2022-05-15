/*
官网: https://golang.google.cn/

TODO 处理掉vuepress跟go有关的md文件
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
