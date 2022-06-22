*******************************
RPC
*******************************

依赖的第三方库

* `gRPC`_ The Go implementation of gRPC: A high performance, open source, general RPC framework that puts mobile and HTTP/2 first.

.. _gRPC: https://www.grpc.io/docs/languages/go/quickstart/

:API文档: https://pkg.go.dev/google.golang.org/grpc

官方入门级Demo

* client - https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
* server - https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go

高级一点的 - https://github.com/grpc/grpc-go/tree/master/examples/route_guide

根据参数类型和复杂程度分类，有以下4种

* 简单的RPC，参数类型字符串、整型等
* 服务端流式RPC。
* 客户端流式RPC
* 双向流RPC

**SetUp**

protoc工具要单独`下载 <https://github.com/protocolbuffers/protobuf/releases/>`_ , 下载完把protoc.exe移动至%GOPATH%\bin\

生成gRPC代码
================================

.. code-block:: shell

    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative routeguide/route_guide.proto

routeguide/route_guide.proto为.proto文件的相对路径/文件名称。如果route_guide.proto在当前目录：

.. code-block:: shell

    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative route_guide.proto

将生成以下文件:

* route_guide.pb.go，其中包含用于填充、序列化和检索请求和响应消息类型的所有协议缓冲区代码。
* route_guide_grpc.pb.go，它包含以下内容:
* 客户端使用RouteGuide服务中定义的方法调用的接口类型(或存根)。
* 服务器要实现的接口类型，同样使用RouteGuide服务中定义的方法。

