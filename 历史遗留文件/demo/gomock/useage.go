package main

import "fmt"

type Foo interface {
	Bar(x int) int
}

func SUT(f Foo) {
	i := f.Bar(99)
	fmt.Println(i)
}

type C struct {
}

func (c *C) Bar(x int) int {
	fmt.Println(x)
	return x + 1
}

func main() {
	c := &C{}
	SUT(c)
}
