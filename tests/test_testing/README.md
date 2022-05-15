# 单元测试

:标准库单元测试学习进度: archive/tar/tar_test.go

-----------------------------------------------------

单元测试函数的定义一

* 创建一个名称以 _test.go 结尾的文件
* 文件包含 `TestXxx` 函数

单元测试函数的定义二

* package名以 _test 为结尾的文件
* 这个文件的函数

例子: archive/tar/example_test.go

运行go test命令将运行编写好的单元测试函数

测试函数运行子测试
===============================

使用T.Run方法，参考例子: archive/tar/reader_test.go line:622

并行测试
===============================

测试函数第一行加上 `t.Parallel()`

使用案例: chromedp/eval_test.go的TestEvaluateNumber函数

---------------------------------------------

设计模式
==============================

数据驱动测试，可参考 archive/tar/reader_test.go