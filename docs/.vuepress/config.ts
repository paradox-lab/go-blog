// @ts-ignore
import { defineUserConfig} from "vuepress";
import type {DefaultThemeOptions} from "vuepress";

export default defineUserConfig<DefaultThemeOptions>({
    base: '/goblog/',
    lang: 'zh_CN',
    title: '蝉时雨丶的博客',
    description: '学习和记录Go & Vue',

    // 主题配置
    themeConfig: {
      logo: 'https://vuejs.org/images/logo.png',
      navbar: [
        // NavbarItem
        {
          text: 'TypeScript',
          link: '/TypeScript.md',
        },
        // NavbarGroup
        {
          text: 'Golang',
          children: [
              {
                  text: '第一部分: 基础',
                  children: [
                      '/Golang/base.md'
                  ]
              },
              {
                  text: '第三部分: 工具',
                  children: [
                      '/Golang/Docker.md'
                  ]
              }
          ],
        },
      ],
    },
})