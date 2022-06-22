========================
SMTP
========================

:github地址: https://github.com/xhit/go-simple-mail

:documentation: https://pkg.go.dev/github.com/xhit/go-simple-mail/v2

Demo
---------------------------

.. code-block:: go

    package main

    import (
        "crypto/tls"
        mail "github.com/xhit/go-simple-mail/v2"
        "log"
        "time"
    )

    const (
        Host = "host"
        UserName = "username"
        PassWord = "********"
    )

    func main() {
        _, err := NewSTMPClient()
        if err != nil {
            log.Fatal(err)
        }
    }

    func newSTMPClient(host, username, password string) (*mail.SMTPClient, error) {
        server := mail.NewSMTPClient()

        // 【修改】SMTP Server的信息
        server.Host = Host
        server.Port = 587
        server.Username = UserName
        server.Password = PassWord

        server.Encryption = mail.EncryptionTLS

        // Variable to keep alive connection
        server.KeepAlive = false

        // Timeout for connect to SMTP Server
        server.ConnectTimeout = 10 * time.Second

        // Timeout for send the data and wait respond
        server.SendTimeout = 10 * time.Second

        // Set TLSConfig to provide custom TLS configuration. For example,
        // to skip TLS verification (useful for testing):
        server.TLSConfig = &tls.Config{InsecureSkipVerify: true}

        // 外部调用时记得调用Close()
        return server.Connect()
    }

    func NewSTMPClient() (*mail.SMTPClient, error) {
        return newSTMPClient(Host, UserName, PassWord)
    }

    func Example1(){
            smtpClient, _ := NewSTMPClient()

            defer smtpClient.Close()

            email := mail.NewMSG()

            email.SetFrom(UserName).
                AddTo("To@example.com").
                AddCc("Cc@example.com").
                SetSubject("Subject")

            //【Notification】设置正文内容，mail.TextPlain为纯文本类型
            email.SetBody(mail.TextPlain, "text")

            // Call Send and pass the client
            err := email.Send(smtpClient)
            if err != nil {
                log.Println(err)
            } else {
                log.Println("Email Sent")
            }
    }

    func Example2(){
        smtpClient, _ := NewSTMPClient()

        defer smtpClient.Close()

        email := mail.NewMSG()

        email.SetFrom(UserName).
            AddTo("To@example.com").
            // 【Notification】新增收件人，再调用一次.Addto方法即可。
            AddTo("To2@example.com").
            AddCc("Cc@example.com").
            SetSubject("Subject")

        //【Notification】设置正文内容，mail.TextHTML为超文本类型
        email.SetBody(mail.TextHTML, "text")

        // Call Send and pass the client
        err := email.Send(smtpClient)
        if err != nil {
            log.Println(err)
        } else {
            log.Println("Email Sent")
        }
    }

常见问题
---------------------------------------

服务器465/587端口, ssl发送邮件失败，本地正常
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

修改配置，改为25端口，不采用SSL&TSL发送

.. code-block:: go

    server.Port = 25
    server.Encryption = mail.EncryptionNone

    // server.TLSConfig = &tls.Config{InsecureSkipVerify: true}
