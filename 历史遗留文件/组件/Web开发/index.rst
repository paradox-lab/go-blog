*******************************
Web开发
*******************************

.. warning::

    学习Go，其实重点学习并且学精Gin就可以了，Gin是我赚钱吃饭的主力工具!!!

:restful api: https://www.cnblogs.com/liuzhongchao/p/9244516.html

.. highlight:: go

依赖的库

* html/templates 模板语法的资料参考text/templates
* `gin`_
    * 官网 - https://gin-gonic.com/zh-cn/

.. _gin: https://github.com/gin-gonic/gin

学习资料:

:Go程序到底需不需要docker: https://studygolang.com/articles/9751

.. tip::

    用配置文件解决路径难以定位的问题

config.ini

.. code-block:: ini

    # possible values : production, development
    app_mode = development

    [paths]
    chdir = D:\myproject\mygin

.. code-block:: go

    func main() {
        cfg, err:=ini.Load("config.ini")
        if err != nil{
            fmt.Println(err)
            return
        }
        appMode := cfg.Section("").Key("app_mode").String()
        if appMode == "production"{
            gin.SetMode(gin.ReleaseMode)
        } else if appMode == "development" {
            gin.SetMode(gin.DebugMode)
        } else {
            fmt.Println("未获取到配置文件app_mode的值!")
            return
        }
        // 根据配置文件判断生产环境和测试环境，测试环境加载静态资源，生产环境用nginx配置静态资源
        r := zhenhua.SetupRouter()
        if appMode == "development"{
            r.Static("/images", "./upload")
        }

        //Listen and Server in 0.0.0.0:8080
        r.Run("0.0.0.0:8080")
    }

.. toctree::
   :maxdepth: 2

   模板语法
   前端
   示例
   单元测试

HTML渲染
====================================

使用 LoadHTMLGlob() 或者 LoadHTMLFiles()::

    func main() {
        router := gin.Default()
        router.LoadHTMLGlob("templates/*")
        //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
        router.GET("/index", func(c *gin.Context) {
            c.HTML(http.StatusOK, "index.tmpl", gin.H{
                "title": "Main website",
            })
        })
        router.Run(":8080")
    }

templates/index.tmpl

.. code-block:: html

    <html>
        <h1>
            {{ .title }}
        </h1>
    </html>

路由组
========================================

.. code-block:: go

    func main() {
        router := gin.Default()

        // 简单的路由组: v1
        v1 := router.Group("/v1")
        {
            v1.POST("/login", loginEndpoint)
            v1.POST("/submit", submitEndpoint)
            v1.POST("/read", readEndpoint)
        }

        // 简单的路由组: v2
        v2 := router.Group("/v2")
        {
            v2.POST("/login", loginEndpoint)
            v2.POST("/submit", submitEndpoint)
            v2.POST("/read", readEndpoint)
        }

        router.Run(":8080")
    }

解决跨域问题
=====================================

https://www.cnblogs.com/zyndev/p/13658550.html

https://www.cnblogs.com/-xuzhankun/p/11145772.html

.. code-block:: go

    package middlewares

    import (
        "github.com/gin-gonic/gin"
        "net/http"
    )

    func Cors() gin.HandlerFunc {
        return func(c *gin.Context) {
            method := c.Request.Method

            c.Header("Access-Control-Allow-Origin", "*")
            c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
            c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
            c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
            c.Header("Access-Control-Allow-Credentials", "true")

            //放行所有OPTIONS方法
            if method == "OPTIONS" {
                c.AbortWithStatus(http.StatusNoContent)
            }
            // 处理请求
            c.Next()
        }
    }

然后在路由使用这个中间件即可

插件
===========================

.. toctree::
   :maxdepth: 2

   后台管理系统