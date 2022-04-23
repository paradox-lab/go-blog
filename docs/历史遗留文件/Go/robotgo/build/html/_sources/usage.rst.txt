**********************
usage
**********************

Window
======================

.. code-block:: go
   :linenos:

    package main

    import (
        "fmt"

        "github.com/go-vgo/robotgo"
    )

    func main() {
      fpid, err := robotgo.FindIds("Xshell.exe")
      if err == nil {
        fmt.Println("pids... ", fpid)

        if len(fpid) > 0 {
          robotgo.ActivePID(fpid[0])

          robotgo.Kill(fpid[0])
        }
      }

      robotgo.ActiveName("Xshell.exe")

      isExist, err := robotgo.PidExists(100)
      if err == nil && isExist {
        fmt.Println("pid exists is", isExist)

        robotgo.Kill(100)
      }

      abool := robotgo.ShowAlert("test", "robotgo")
      if abool {
          fmt.Println("ok@@@ ", "ok")
      }

      title := robotgo.GetTitle()
      fmt.Println("title@@@ ", title)
    }

**第10行**  应该提供准确而非模糊的应用程序名称，所有的应用程序名称可通过 `robotgo.FindNames()` 获取。

**第21行**  经测试没有正常激活Windows应用程序，已经有相关 `issue <https://github.com/go-vgo/robotgo/issues/320>`_ ，可持续关注这个
issue的进展情况。

Keyboard
=====================

.. code-block:: go
   :linenos:

   package main

   import (
     "fmt"

     "github.com/go-vgo/robotgo"
   )

   func main() {
     robotgo.TypeStr("Hello World")
     robotgo.TypeStr("だんしゃり", 1.0)
     // robotgo.TypeString("テストする")

     robotgo.TypeStr("Hi galaxy. こんにちは世界.")
     robotgo.Sleep(1)

     // ustr := uint32(robotgo.CharCodeAt("Test", 0))
     // robotgo.UnicodeType(ustr)

     robotgo.KeyTap("enter")
     // robotgo.TypeString("en")
     robotgo.KeyTap("i", "alt", "command")

     arr := []string{"alt", "command"}
     robotgo.KeyTap("i", arr)

     robotgo.WriteAll("Test")
     text, err := robotgo.ReadAll()
     if err == nil {
       fmt.Println(text)
     }
   }

**第22行**  `command` 等同于win键， `"i", "alt", "command"` 等同于 alt + win + i的组合键。
如果要按下组合键shif+alt+N，代码为 `robotgo.KeyTap("N", "shift", "alt")` 。

**第27行**  将数据写入剪切板。

**第28行**  从剪切板获取数据。

