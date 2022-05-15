/*
TODO docker api
*/
package tests

import (
	"fmt"
	"os/exec"
	"testing"
)

// TestStartDocker 启动Docker
func TestStartDocker(t *testing.T) {
	cmd := exec.Command("systemctl", "start", "docker")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
