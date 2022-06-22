*******************************
logrus
*******************************

.. highlight:: go

:github: https://github.com/sirupsen/logrus
:其他参考博客: https://segmentfault.com/a/1190000021706728

**Example**

.. code-block:: go

    package main

    import (
      log "github.com/sirupsen/logrus"
    )

    func main() {
      log.WithFields(log.Fields{
        "animal": "walrus",
      }).Info("A walrus appears")
    }

显式调用来源::

    log.SetReportCaller(true)

配置修改::

    package main

    import (
      "os"
      log "github.com/sirupsen/logrus"
    )

    func init() {
      // Log as JSON instead of the default ASCII formatter.
      log.SetFormatter(&log.JSONFormatter{})

      // Output to stdout instead of the default stderr
      // Can be any io.Writer, see below for File example
      log.SetOutput(os.Stdout)

      // Only log the warning severity or above.
      log.SetLevel(log.WarnLevel)
    }

    func main() {
      log.WithFields(log.Fields{
        "animal": "walrus",
        "size":   10,
      }).Info("A group of walrus emerges from the ocean")

      log.WithFields(log.Fields{
        "omg":    true,
        "number": 122,
      }).Warn("The group's number increased tremendously!")

      log.WithFields(log.Fields{
        "omg":    true,
        "number": 100,
      }).Fatal("The ice breaks!")

      // A common pattern is to re-use fields between logging statements by re-using
      // the logrus.Entry returned from WithFields()
      contextLogger := log.WithFields(log.Fields{
        "common": "this is a common field",
        "other": "I also should be logged always",
      })

      contextLogger.Info("I'll be logged with common and other field")
      contextLogger.Info("Me too")
    }

实例对象::

    package main

    import (
      "os"
      "github.com/sirupsen/logrus"
    )

    // Create a new instance of the logger. You can have any number of instances.
    var log = logrus.New()

    func main() {
      // The API for setting attributes is a little different than the package level
      // exported logger. See Godoc.
      log.Out = os.Stdout

      // You could set this to any `io.Writer` such as a file
      // file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
      // if err == nil {
      //  log.Out = file
      // } else {
      //  log.Info("Failed to log to file, using default stderr")
      // }

      log.WithFields(logrus.Fields{
        "animal": "walrus",
        "size":   10,
      }).Info("A group of walrus emerges from the ocean")
    }

输出到文件中
=================================================

.. code-block:: go

    package main

    import (
      "bytes"
      "io"
      "log"
      "os"

      "github.com/sirupsen/logrus"
    )

    func main() {
      writer1 := &bytes.Buffer{}
      writer2 := os.Stdout
      writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
      if err != nil {
        log.Fatalf("create file log.txt failed: %v", err)
      }

      // 同时将日志写到bytes.Buffer、标准输出和文件中
      logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
      logrus.Info("info msg")
    }