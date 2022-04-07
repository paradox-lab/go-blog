********************
Vue
********************

该章节以vue-admin-pro作为应用案例来记录学习笔记。

二次开发指导
============================

如何在完整版将需要的vue组件搬到精简版上?以完整版的系统管理菜单下的字典管理为例子做说明:

* 安装第三方库vxe-table
* 同步main.ts代码, 加载vxe-table插件
* 同步 ``src/plugins/vex-table/index.ts`` 代码，按需加载组件，在main.ts中使用
* 在 ``src/views`` 创建自定义的Vue组件，从完整版COPY或者参考后自己编写代码
* 在 ``src/router`` 注册路由

.. toctree::
   :maxdepth: 2
   :caption: Contents:

   vue-router
