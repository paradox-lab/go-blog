***************************
vue-router
***************************

:pure相关代码存放目录: ``src/router`` 全局路由模块

:入口: main.ts

在 ``src/router/index.ts`` 中定义了以下路由:

* homeRouter
* errorRouter
* externalLink
* remainingRouter

router-view
=================================

`router-view<https://next.router.vuejs.org/zh/guide/#router-view>`_ 将显示与 url 对应的组件。你可以把它放在任何地方，以适应你的布局。

*应用*

``src/App.vue``

history模式
================================

有Hash和HTML5两种模式。vue-admin-pro由配置文件的VITE_ROUTER_HISTORY决定history模式，默认是hash模式，因此url能看到 ``#`` 这个符号，对SEO有影响。

router二次开发指导
============================

* 增加新菜单时，在 ``src/router/modules`` 目录下创建新的ts文件，编写路由的相关内容
* 将新路由注册到 ``src/router/modules/index`` 的routes变量，页面即可显示

