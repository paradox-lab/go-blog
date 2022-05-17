/*
github地址: https://github.com/chromedp/chromedp
Docs: https://pkg.go.dev/github.com/chromedp/chromedp
Example: https://github.com/chromedp/examples
知乎: https://zhuanlan.zhihu.com/p/515734676?

CentOS 安装chromedp: https://blog.csdn.net/weixin_44439675/article/details/117789187
*/

package chromedp

import (
	"context"
	"encoding/json"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

var allocCtx context.Context

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

func TestMain(m *testing.M) {
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		// 是否无头模式，默认true
		chromedp.Flag("headless", true),
		// 忽略ERR_CERT_AUTHORITY_INVALID警告
		// 【参考】 https://github.com/chromedp/chromedp/issues/157
		//chromedp.Flag("ignore-certificate-errors", "1"),
	)
	var cancel context.CancelFunc
	allocCtx, cancel = chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	code := m.Run()
	os.Exit(code)
}

func TestQuickStart(t *testing.T) {
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

// TestCookies 保存Cookies和使用Cookies免密码登录
func TestCookies(t *testing.T) {
	// https://github.com/chromedp/examples/blob/16788b67ed42809ca1f23294f80e9b89f67662e8/cookie/main.go#L74-L78
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	filename := testGetAllCookies(t, ctx)
	// 遍历设置Cookie
	testSetCookie(t, filename)
	// 一次性设置Cookies
	testSetCookies(t, filename)

	// 创建用于超时退出的上下文管理器
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
}

func testGetAllCookies(t *testing.T, ctx context.Context) string {
	file, err := os.OpenFile(t.TempDir()+"/cookies", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer file.Close()
	err = chromedp.Run(ctx,
		chromedp.Navigate(`https://luzhenxiong.top/admin/`),
		chromedp.SetValue(`input[name="username"]`, "admin", chromedp.ByQuery),
		chromedp.SetValue(`input[name="password"]`, "password123", chromedp.ByQuery),
		chromedp.Click(`button`, chromedp.ByQuery),
		chromedp.WaitVisible(`main`, chromedp.ByID),
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies, err := network.GetAllCookies().Do(ctx)
			if err != nil {
				return err
			}

			// TODO do json test
			j, err := json.Marshal(cookies)
			if err != nil {
				return err
			}

			//【write】写入[]byte类型数据
			_, err = file.Write(j)
			if err != nil {
				return err
			}

			return nil
		}),
	)

	if err != nil {
		log.Fatal(err)
	}
	return file.Name()
}

func testSetCookie(t *testing.T, cookiesFile string) {
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
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
		chromedp.Navigate(`https://luzhenxiong.top/admin/`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			file, err := os.Open(cookiesFile)
			if err != nil {
				return err
			}

			defer file.Close()
			// 读取文件数据
			jsonBlob, _ := ioutil.ReadAll(file)
			var cookies []network.CookieParam
			// Json解码
			err = json.Unmarshal(jsonBlob, &cookies)
			if err != nil {
				return err
			}
			for _, cookie := range cookies {
				err := network.SetCookie(cookie.Name, cookie.Value).
					WithExpires(cookie.Expires).
					WithDomain(cookie.Domain).
					WithHTTPOnly(true).
					Do(ctx)
				if err != nil {
					return err
				}
			}

			return nil
		}),
		chromedp.Reload(),
		chromedp.Title(&text),
	)

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, text, "Django 站点管理员")
	// 创建用于超时退出的上下文管理器
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
}

func testSetCookies(t *testing.T, cookiesFile string) {
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
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
		chromedp.Navigate(`https://luzhenxiong.top/admin/`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			file, err := os.Open(cookiesFile)
			if err != nil {
				return err
			}

			defer file.Close()
			// 读取文件数据
			jsonBlob, _ := ioutil.ReadAll(file)
			var cookies []*network.CookieParam
			// Json解码
			err = json.Unmarshal(jsonBlob, &cookies)
			if err != nil {
				return err
			}

			err = network.SetCookies(cookies).Do(ctx)
			if err != nil {
				return err
			}

			return nil
		}),
		chromedp.Reload(),
		chromedp.Title(&text),
	)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, text, "Django 站点管理员")

	// 创建用于超时退出的上下文管理器
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
}
