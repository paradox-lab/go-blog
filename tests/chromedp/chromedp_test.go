/*
github地址: https://github.com/chromedp/chromedp
Docs: https://pkg.go.dev/github.com/chromedp/chromedp
Example: https://github.com/chromedp/examples

CentOS 安装chromedp: https://blog.csdn.net/weixin_44439675/article/details/117789187
*/

package chromedp

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestVersion(t *testing.T) {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("pkg.go.dev"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("body > main > header > div > div.go-Main-headerDetails > span:nth-child(1) > a", func(e *colly.HTMLElement) {
		assert.Equal(t, e.Text, "Version: v0.8.1")
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://pkg.go.dev/github.com/chromedp/chromedp")
}

func TestQuickStart(t *testing.T) {
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		// 是否无头模式，默认true
		chromedp.Flag("headless", true),
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
	var text string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`http://www.baidu.com`),
		chromedp.Title(&text),
	)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, text, "百度一下，你就知道")

	// 创建用于超时退出的上下文管理器
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
}

func TestGetCookies(t *testing.T) {
	// TODO 保存cookies
}

func TestSetCookies(t *testing.T) {
	// TODO 设置cookies, 用gva在线demo做实验
	// 验证码识别 收费: 超级鹰平台 http://www.chaojiying.com/
	// 免费: 百度ocr:
	// Python开源库: PyTesseract
}
