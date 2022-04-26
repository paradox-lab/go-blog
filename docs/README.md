# Vue

* 官网: https://staging-cn.vuejs.org/
* 在线敲码: https://sfc.vuejs.org/

**理论知识**

![Vue.js 3.0从入门到精通](/images/img.png)

**技术栈确定**

Vue技术体系在同一场景下提供了不同的技术点和工具选择, 例如项目可以用javascript或者typescript编写, API风格可以选择选项式或者组合式等等。
对于我这种非专业前端，应该走专精路线, 不要全学, 提前确定好一种选择, 避免写vue时产生各种纠结, 以及无用功学习。

* API风格偏好
    - 选项式 - Vue2.js写法
    - 组合式 - Vue3.js写法 ✔
* 安装方式
    - CDN
    - NPM
    - Vue Cli
    - Vite ✔
* 编辑器
    - vs code
    - JetBrains ✔
* 网络请求
    - axios
    - fetch ✔
* 测试框架
    - vitest ✔
    - Peeky
    - Jest

**调试工具**

https://devtools.vuejs.org/guide/installation.html#standalone

因为打不开google商店，所以要么在github下载然后导入到chrome, 要么用npm下载, 然后在html增加一条script标签导入。

在Pycharm对vue代码进行调试: https://www.jetbrains.com/help/pycharm/2022.1/debugging-javascript-in-chrome.html#debugging_js_on_local_host_development_mode

注意点:

1. 点击localhost那个链接
2. 要按下快捷键Shift + Ctrl来点击链接

**使用社区模板启动项目**

* https://github.com/pohunchn/vite-ts-quick
* https://github.com/xiaoxian521/vue-pure-admin

## awesome

- [富文本编辑器](https://ckeditor.com/) - ckeditor5

### ckeditor5

https://ckeditor.com/docs/ckeditor5/latest/installation/getting-started/frameworks/vuejs-v3.html

### vite-ts-quick

快速启动vite+ts体系的vue项目

### fetch向后端发送api请求

1. url地址要以"/"结尾, 原因未明
2. 后端增加跨域处理

TODO 未来增加跨域的api请求的单元测试验证

## Vuepress

文档: [https://v2.vuepress.vuejs.org/zh/](https://v2.vuepress.vuejs.org/zh/)

::: tip 
在windows安装的vuepress模块不能在linux上使用
:::

