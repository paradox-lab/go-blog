---
title: "interface"
date: 2022-06-17T09:48:05+08:00
draft: false
---

# 接口

接口是一组仅包含方法名、参数、返回值的未具体实现的方法的集合。如果实现了接口的所有方法，则认为实现了该接口，无需在该类型上显示的添加声明。

应用场景

* 抽象。接口可以理解为某一个方面的抽象，可以是多对一的(多个类型实现一个接口)，这也是多态的体现。
* duck typing。

## 开发阶段判断是否实现了接口

```go
var _ Phone = new(AndroidPhone)
// 等价于
var _ Phone = (*AndroidPhone)(nil)
```
