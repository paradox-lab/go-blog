/*
https://github.com/otiai10/gosseract

Windows 安装tesseract: https://blog.csdn.net/u013269298/article/details/86679091
*/

package gosseract

import (
	"github.com/gocolly/colly/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("pkg.go.dev"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("body > main > header > div > div.go-Main-headerDetails > span:nth-child(1) > a", func(e *colly.HTMLElement) {
		assert.Equal(t, e.Text, "Version: v2.3.1")
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://pkg.go.dev/github.com/otiai10/gosseract/v2")
}
