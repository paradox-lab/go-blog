# sftp
> The sftp package provides support for file system operations on remote ssh servers using the SFTP subsystem. 

**github地址**: https://github.com/pkg/sftp

**sftp手册**: https://pkg.go.dev/github.com/pkg/sftp

**ssh文档**: https://pkg.go.dev/golang.org/x/crypto@v0.0.0-20201221181555-eec23a3978ad/ssh

## version
### 1.13.0(2021.3.7)
Client:
- 增加 Client.Extensiosn 成员方法以列出支持的server拓展。
- 增加对fsync@openssh.com的支持

## Usage
[demo](demo/example.go)


