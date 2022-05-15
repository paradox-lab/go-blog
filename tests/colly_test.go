/*
https://github.com/gocolly/colly
*/

package tests

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
		assert.Equal(t, e.Text, "Version: v2.1.0")
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://pkg.go.dev/github.com/gocolly/colly/v2")
}
