---
title: "Context"
date: 2022-06-17T11:01:34+08:00
draft: false
---

资料参考 https://zhuanlan.zhihu.com/p/68792989

资料参考2: https://www.5axxw.com/questions/content/hu6pqz

WithTimeout例程
```go
func SayHello(name string) (string, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
    defer cancel()
    ch := make(chan string, 1)

    go func() {

        conn, err := NewGrpcConn()
        if err != nil {
            config.Logger.Fatalf("did not connect: %v", err)
        }
        defer conn.Close()

        c := NewGrpcClient(conn)

        r, err := c.SayHello(ctx, &HelloRequest{Name: name})
        if err != nil {
            log.Fatalf("could not greet: %v", err)
        }

        if r.String() != "" {
            log.Printf("Greeting: %s", r.GetMessage())
            ch <- fmt.Sprintf("Greeting: %s", r.GetMessage())
        }
    }()

    select {
    case s :=<-ch:
        return s, nil
    case <-ctx.Done():
        err := errors.New("执行超时(15秒)")
        return "", err
    }
}
```