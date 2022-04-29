/*******************************************************************************
 * Copyright (c) 2021
 * @Description：Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
 * @Date: 2021/3/12 下午5:25
 *
 * @github地址: https://github.com/golang/go
 * @标准库文档 : https://studygolang.com/pkgdoc
 * @Go指南: http://tour.studygolang.com/welcome/1
 * @awesome库: https://github.com/avelino/awesome-go
 * @awesome库(中文版): https://github.com/jobbole/awesome-go-cn
 *
 * @use version: v1.15.8
 ******************************************************************************

 - Put: 上传文件到sftp,实现Python中paramiko.sftp_client.SFTPClient的put功能
 - Get: 从sftp下载文件,实现Python中paramiko.sftp_client.SFTPClient的get功能

 - Example1: 公钥登录,并校验服务器
 - Example2: 公钥登录,不校验服务器
 - Example3: 密码登录,官方Demo

 ******************************************************************************/

package main

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/sftp"
)

func main() {

}

// Put 上传文件到sftp
// 实现Python中paramiko.sftp_client.SFTPClient的put功能
func Put(c *sftp.Client,localpath, remotepath string) error {
	srcFile, err := os.Open(localpath)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	dstFile, err := c.Create(remotepath)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	buf := make([]byte, 1024)

	for {

		n, _ := srcFile.Read(buf)

		if n == 0 {
			break
		}

		if n != 1024 {
			dstFile.Write(buf[:n])
		} else {
			dstFile.Write(buf)
		}

	}

	return err
}

// Get 从sftp下载文件
// 实现Python中paramiko.sftp_client.SFTPClient的get功能
func Get(c *sftp.Client, remotepath, localpath string) error {
	srcFile, err := c.Open(remotepath)
	if err != nil {
		log.Fatalf("remotepath: %v", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(localpath)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	if _, err = srcFile.WriteTo(dstFile); err != nil {
		log.Fatal(err)
	}
	return err
}

func Example1() {
	var hostKey ssh.PublicKey
	// A public key may be used to authenticate against the remote
	// server by using an unencrypted PEM-encoded private key file.
	//
	// 密钥文件
	key, err := ioutil.ReadFile("/home/user/.ssh/id_rsa")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: "user",
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		// 【注意】这里是要验证服务端的公钥，因此先给hostKey赋值，这里的hostKey=nil。
		// 【output】ssh: handshake failed: ssh: required host key was nil
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	// Connect to the remote server and perform the SSH handshake.
	client, err := ssh.Dial("tcp", "host.com:22", config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer client.Close()
}

func Example2(){
	// 【修改】密钥文件路径
	key, err := ioutil.ReadFile("/home/user/.ssh/id_rsa")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: "user",
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		// 【注意】不做任何验证传入ssh.InsecureIgnoreHostKey()
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 【修改】主机地址和端口
	client, err := ssh.Dial("tcp", "host.com:22", config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer client.Close()
}

func Example3()  {
	//【参考】https://pkg.go.dev/golang.org/x/crypto@v0.0.0-20201221181555-eec23a3978ad/ssh#Dial
	//【核心依赖库】"golang.org/x/crypto/ssh"

	var hostKey ssh.PublicKey
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.
	config := &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{
			ssh.Password("yourpassword"),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}
	client, err := ssh.Dial("tcp", "yourserver.com:22", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}
	defer client.Close()

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
	
}
