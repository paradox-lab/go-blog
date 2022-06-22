---
title: "Home"
date: 2022-05-24T17:41:32+08:00
draft: false
---

## hugo

创建文章

```text
hugo new post/first.zh.md
```

https://github.com/gohugoio/gohugoioTheme

本地运行hugo

```text
hugo server --theme=hugo-theme-learn --buildDrafts
```

编译打包
```text
hugo --theme=hyde --baseUrl="http://coderzh.github.io/"
```

自动部署
```text
https://gohugo.io/hosting-and-deployment/hosting-on-github/#build-hugo-with-github-action
```

## theme

使用的是[hugo-theme-learn](https://learn.netlify.app/en/)

## Home的问题

点击Home按钮链接到 `http://localhost:1313/` , 需要映射到 `http://localhost:1313/go-blog` , 但找不到配置方案。

已有人提出这个问题: https://github.com/matcornic/hugo-theme-learn/issues/507