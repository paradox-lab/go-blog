/*
https://pkg.go.dev/net/http
*/
package tests

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	resp, err := http.Get("https://pkg.go.dev/net/http")
	assert.Nil(t, err)
	defer resp.Body.Close()
	assert.Equal(t, resp.StatusCode, 200)
	body, err := io.ReadAll(resp.Body)
	log.Println(string(body))

	_, err = goquery.NewDocumentFromReader(resp.Body)
	assert.Nil(t, err)
}
