/*
接口是一组仅包含方法名、参数、返回值的未具体实现的方法的集合。如果实现了接口的所有方法，则认为实现了该接口，无需在该类型上显示的添加声明。

应用场景

* 抽象。接口可以理解为某一个方面的抽象，可以是多对一的(多个类型实现一个接口)，这也是多态的体现。
* duck typing

*/

package golang

import (
	"fmt"
)

type Phone interface {
	call()
}

type AndroidPhone struct {
}
type IPhone struct {
}

func (a AndroidPhone) call() {
	fmt.Println("我是安卓手机，可以打电话了")
}

func (i IPhone) call() {
	fmt.Println("我是苹果手机，可以打电话了")
}
