*******************************
UI自动化
*******************************

.. highlight:: go

依赖的第三方库

* `chromedp`_

.. _chromedp: https://github.com/chromedp/chromedp

.. warning::

    在实际使用中，发现chromedp挺多问题的，目前还不够稳定

**演示代码的前置动作**

.. code-block:: go

    import (
        "context"
        "github.com/chromedp/cdproto/cdp"
        "github.com/chromedp/chromedp"
    )
    // --------------------------------------------------------------

    // 以下是代码片段，需要放入到函数中
    opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		// 是否无头模式，默认true
		chromedp.Flag("headless", false),
		// 忽略ERR_CERT_AUTHORITY_INVALID警告
		// 【参考】 https://github.com/chromedp/chromedp/issues/157
		//chromedp.Flag("ignore-certificate-errors", "1"),
	)

    allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
    defer cancel()

    // 创建chrome实例
    ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
    defer cancel()

    // 创建用于超时退出的上下文管理器
    ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
    defer cancel()

与页面交互
===============================================

打开页面
-----------------------------------------------

.. code-block:: go

    err := chromedp.Run(ctx,
        chromedp.Navigate(`http://www.baidu.com`),
    )
    if err != nil {
        log.Fatal(err)
    }

下载文件
----------------------------------------------

.. warning::

    官方样例 `download_file`_ ，测试过下载失败，而且查 `issue#807`_ 可知，目前确实有问题的。

.. _download_file: https://github.com/chromedp/examples/blob/master/download_file/main.go
.. _issue#807: https://github.com/chromedp/chromedp/issues/807

切换frame&iframe窗口并与元素交互
----------------------------------------------

.. code-block:: go

    var frames []*cdp.Node
    if err = chromedp.Run(ctx,
        // 获取frame节点
        chromedp.Nodes(`frame[name="FrameName"]`, &frames, chromedp.ByQuery),
        ); err != nil {
        log.Fatal(err)
    }

    frame := frames[0]
    if err = chromedp.Run(ctx,
        // 通过XPath定位并点击元素
        chromedp.Click(`//div[text()="文本值"]`, chromedp.BySearch, chromedp.FromNode(frame)),
    ); err != nil {
        log.Fatal(err)
    }

.. seealso::

    chromedp官方的 `example`_ 仓库并没有放出如何操作frame&iframe的代码样例，可以参阅的相关资料是 `issue#72`_
    和nav的 `单元测试代码`_ (在TestQueryIframe函数)。

.. _example: https://github.com/chromedp/examples
.. _issue#72: https://github.com/chromedp/chromedp/issues/72
.. _单元测试代码: https://github.com/chromedp/chromedp/blob/master/nav_test.go

.. warning::

    frame元素定位和操作frame里面的元素必要拆分出来分开run，放在一起，frame会提前去定位，但此时frame并没有定位成功。
    同理，在chromedp.Tasks也是提前定位了frame。

元素定位
====================================================

* chromedp.ByQuery - CSS选择器
    - "p" - 获取p标签的第一个元素
    - ".example" - 获取class="example"的第一个元素
    - "p.example" - 获取class="example"的p标签的第一个元素
    - "a[target]" - 获取第一个含有target属性的a标签元素
    - "#div_button_excel > a" - 获取ID为div_button_excel的子标签 `a`
* chromedp.ByQueryAll - CSS选择器，返回所有符合条件的元素
* chromedp.BySearch - Xpath或者CSS selector查询, Query函数的默认查找方法
* chromedp.ByID - 根据ID查询，`#` + ID

.. warning::

    使用 `.BySearch` 时, 如果查找元素是 `section[class="el-container app-body"]`，
    `sel="section[class="app-body"]"` 是行不通的，必须填入所有的class name

.. warning::

    实践中，经常发现元素定位失败，如果selenium或者js定位成功，改为调用js语句去执行动作吧。

元素取值与赋值
================================================

**setup**

.. code-block:: go

    var res string

获取元素的text
-----------------------------------------------

.. code-block:: go

    chromedp.Run(ctx,
        // 文本赋值到 `res`
        chromedp.Text(`#pkg-overview`, &res, chromedp.NodeVisible, chromedp.ByID),
    )

获取input元素的value
-----------------------------------------------

.. code-block:: go

    // input元素xpath路径
    input := `//div[@class="wide-input el-input"]/input`
    chromedp.Run(ctx,
        // 文本赋值到res
        chromedp.Evaluate(fmt.Sprintf(`document.evaluate('%v', document).iterateNext().value`, input), &res),
    )

变更input元素的value
------------------------------------------------

在selenium，习惯是先使用clear方法清空value，然后使用send_keys赋值。

在chromedp，也有Clear方法和SendKeys方法，不过Clear方法尝试过定位不了（可能对定位方法有限制）, SendKeys倒没有问题。

其实还有一个SetValue方法，可以从2个步骤直接一步到位搞定。

.. code-block:: go

    chromedp.SetValue(`#id`, "value", chromedp.ByID),

------------------------------------------------

Faq
=================================================

有类似selenium WebDriverWait的等待超时设置吗?
---------------------------------------------------------

很遗憾，没有。有一个解决办法是设置一个timeout的context，传入到Run中。

.. seealso::

    `issue#647 <https://github.com/chromedp/chromedp/issues/647>`_

如何变更element-ui框架时间选择器组件的值?
----------------------------------------------------------

**SetUp**

.. code-block:: go

    var res interface{}

4个步骤

1. 更改input框的value属性::

    chromedp.SetValue(`input[placeholder="请选择起始时间"]`, "2021/07/15")

2. 使用Click函数点击日期时间选择器，这一步是为了第4步运行成功

3. 触发input事件::

    chromedp.Evaluate(`document.querySelector("input[placeholder='请选择结束时间']").dispatchEvent(new Event('input'))`, res),

.. seealso::

    这一步是触发了input中的v-model，修改了绑定变量的数据

4. 触发change事件::

    chromedp.Evaluate(`document.querySelector("input[placeholder='请选择结束时间']").dispatchEvent(new Event('change'))`, res),

.. seealso::

    这一步执行后，时间选择器的当前选中值将显示正确

如何等待手工关闭网页后才退出程序?
-------------------------------------------

在main应用程序的最后位置加上一行代码::

    <- ctx.Done()

这行代码在自动化完成后阻塞整个程序，直到网页被关闭

忽略ERR_CERT_AUTHORITY_INVALID警告
-------------------------------------------

增加参数 `chromedp.Flag("ignore-certificate-errors", "1")`

.. code-block:: go

    opts := append(
        chromedp.DefaultExecAllocatorOptions[:],
        chromedp.Flag("ignore-certificate-errors", "1"),
    )

参考自 https://github.com/chromedp/chromedp/issues/157

忽略NET::ERR_SSL_OBSOLETE_VERSION警告
---------------------------------------------

也是增加参数 `chromedp.Flag("ignore-certificate-errors", "1")`