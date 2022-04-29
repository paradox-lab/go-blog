.. _faq:

常见问题
==============

处理中文乱码
--------------

.. seealso::

    https://studygolang.com/articles/8897

导入github.com/axgle/mahonia库

.. code-block::

    enc := mahonia.NewDecoder("gbk")
    goStr:= enc.ConvertString(robotgo.GetTitle())
    fmt.Println(goStr)

.. note::

    robot.GetTitle()返回C.GoString的值，遇到中文会显示乱码