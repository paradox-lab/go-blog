package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	devtoolsWsURL := flag.String("devtools-ws-url", "", "DevTools WebSsocket URL")
	flag.Parse()
	if *devtoolsWsURL == "" {
		log.Fatal("must specify -devtools-ws-url")
	}

	// create allocator context for use with creating a browser context later
	allocatorContext, cancel := chromedp.NewRemoteAllocator(context.Background(), *devtoolsWsURL)
	defer cancel()

	// create context
	ctxt, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()

	// run task list
	//var body string
	if err := chromedp.Run(ctxt,
		chromedp.Navigate("https://www.baidu.com"),
		//chromedp.WaitVisible("#logo_homepage_link"),
		//chromedp.OuterHTML("html", &body),
	); err != nil {
		log.Fatalf("Failed getting body of duckduckgo.com: %v", err)
	}

	log.Println("Body of duckduckgo.com starts with:")
	//log.Println(body[0:100])
	time.Sleep(1 * time.Hour)
}
