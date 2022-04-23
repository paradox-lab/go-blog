# gomock
> gomock is a mocking framework for the Go programming language.

> 体验了一轮，感觉只能用于测试函数，不能用于测试结构体方法。
> 关键是抽象成接口，然后把接口传入函数。

**github地址**: https://github.com/golang/mock

**Documentation**: https://pkg.go.dev/github.com/golang/mock/gomock

**其他资料**: https://geektutu.com/post/quick-gomock.html

## 用法
```text
(1) 定义你想mock的接口:
      type MyInterface interface {
        SomeMethod(x int64, y string)
      }
(2) 使用 mockgen工具从接口生成mock代码
(3) 在测试代码使用mock:
      func TestMyThing(t *testing.T) {
        mockCtrl := gomock.NewController(t)
        defer mockCtrl.Finish()

        mockObj := something.NewMockMyInterface(mockCtrl)
        mockObj.EXPECT().SomeMethod(4, "blah")
        // pass mockObj to a real object and play with it.
      }    
```


# 运行mockgen
> mockgen有两种操作模式:source和reflect

## Source模式
> Source模式从源文件生成模拟接口。它通过使用-source标志来启用。
> 在这种模式下可能有用的其他标志是-imports和-aux_files。

Example:
```commandline
mockgen -source=foo.go [other options]
```
> Windows系统把“=”替换为空格

## Flags
> mockgen命令用于为给定的包含要模拟的接口的Go源文件生成模拟类的源代码。它支持以下标志:

- -source: 包含要模拟的接口的文件。
- -destination: 要在其中写入结果源代码的文件。如果不设置此参数，代码将打印到标准输出。
- -package: 用于生成的模拟类源代码的包。如果不设置此参数，则包名为mock_与输入文件的包连接。
- -imports: 应该在生成的源代码中使用的显式导入列表，以逗号分隔的元素列表foo=bar/baz的形式指定，其中bar/baz是被导入的包，foo是生成的源代码中用于包的标识符。
- -aux_files: 应参考的附加文件列表，以解决问题，例如，在另一个文件中定义的嵌入式接口。
  这是一个以逗号分隔的元素列表，foo=bar/baz.go,其中bar/baz.go是源文件，foo是-source文件所使用的文件的包名。
- -build_flags: (仅relect模式可用)Flags逐个传递到go构建。  
- -mock_names: 生成的模拟的自定义名称列表。这被指定为以逗号分隔的表单元素列表 Repository=MockSensorRepository,EndPoint=MockSensorEndpoint,
  Respository是接口名称和MockSensorRepository所需的mock名称(模拟工厂方法和模拟记录器将模拟的名字命名)。
  如果其中一个接口没有指定自定义名称，则将使用默认命名约定。
- -self_package: 生成代码的完整包导入路径。该标志的目的是通过试图包含自己的包来防止在生成的代码中导入循环。
  如果mock的包被设置为它的一个输入(通常是主输入)，并且输出是stdio，因此mockgen无法检测到最终的输出包，就会发生这种情况。
  设置这个标志将告诉mockgen要排除哪个导入。
- -copyright_file: 版权文件，用于将版权头添加到生成的源代码中。  