// https://v2.vuepress.vuejs.org/zh/guide/configuration.html#%E9%85%8D%E7%BD%AE%E6%96%87%E4%BB%B6

module.exports = {
  lang: 'zh-CN',
  title: '蝉时雨丶的博客',
  description: '学习和记录Go和Vue',

  // 主题配置
  themeConfig: {
    logo: 'https://vuejs.org/images/logo.png',
    navbar: [
      // NavbarItem
      {
        text: 'Foo',
        link: '/foo/',
      },
      // NavbarGroup
      {
        text: 'Group',
        children: ['/group/foo.md', '/group/bar.md'],
      },
      // 字符串 - 页面文件路径
      '/bar/README.md',
    ],
  },
}