/*
目前还没找到怎么交互式执行shell的资料

cmd.Start和cmd.Run的区别
Start 非阻塞, Run 阻塞
*/

package tests

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"testing"
)

func TestCommand(t *testing.T) {
	// 调用默认浏览器打开w
	cmd := exec.Command("explorer", "http://127.0.0.1:8000/")
	err := cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestOutput(t *testing.T) {
	cmd := exec.Command("npm", "init", "vite@latest", "project", "--", "template", "vue")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(stdout)
	errReader := bufio.NewReader(stderr)

	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Print(line)
	}

	for {
		line, err := errReader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Print(line)
	}

	if err = cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
