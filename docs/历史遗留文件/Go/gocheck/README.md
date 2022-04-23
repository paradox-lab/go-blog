# gocheck
> The Go language provides an internal testing library, named testing, which is relatively slim due to the fact that the standard library correctness by itself is verified using it. The check package, on the other hand, expects the standard library from Go to be working correctly, and builds on it to offer a richer testing framework for libraries and applications to use.

**github**: https://github.com/go-check/check/tree/v1

**Docs**: http://labix.org/gocheck

**API文档**: https://github.com/go-check/check/tree/v1

**编译命令**: go test -c

## 断言方法
- Equals: 相等
- ErrorMatches: Error错误信息的正则表达式匹配
- Matches: 正则表达式
- IsNil 断言为空

# demos
## basic sample
```golang
package hello_test

import (
    "testing"
    "io"

    . "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestHelloWorld(c *C) {
    c.Assert(42, Equals, "42")  // 失败
    c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")  // pass
    c.Check(42, Equals, 42)  // pass
}
```

## Using fixtures
- func (s *SuiteType) SetUpSuite(c *C) - 在suite开始运行时运行一次。
- func (s *SuiteType) SetUpTest(c *C) - 在每个测试或基准测试开始运行之前运行。
- func (s *SuiteType) TearDownTest(c *C) - 在每个测试或基准测试运行之后运行。
- func (s *SuiteType) TearDownSuite(c *C) - 在所有测试或基准测试完成运行后运行一次。

```golang
type Suite struct{
    dir string
}

func (s *MySuite) SetUpTest(c *C) {
    s.dir = c.MkDir()
    // Use s.dir to prepare some data.
}

func (s *MySuite) TestWithDir(c *C) {
    // Use the data in s.dir in the test.
}
```

## Adding benchmarks
TODO

## Skipping tests
TODO

## Running tests and output sample
TODO

## Assertions and checks
```text
gocheck使用*C的两种方法来验证测试用例中获得的值的期望:Assert和Check。这两个方法都接受相同的参数，
它们之间的唯一区别是，当Assert失败时，测试立即中断，而Check将失败，返回false，并允许它继续进行进一步的检查。
```
```golang
func (s *S) TestSimpleChecks(c *C) {
    c.Assert(value, Equals, 42)
    c.Assert(s, Matches, "hel.*there")
    c.Assert(err, IsNil)
    c.Assert(foo, Equals, bar, Commentf("#CPUs == %d", runtime.NumCPU())
}
```
> 最后一条语句将在通常的调试信息旁边显示所提供的消息，但只有在检查失败时才会如此。