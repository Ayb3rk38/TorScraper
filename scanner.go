package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

type ScanResult struct {
	URL            string `json:"url"`
	HTML           string `json:"html"`
	Status         string `json:"status"`
	ScreenshotFile string `json:"screenshot_file"`
}

func Worker(id int, jobs <-chan string, results chan<- ScanResult, wg *sync.WaitGroup, browserCfg BrowserConfig) {
	defer wg.Done()

	for url := range jobs {
		log.Printf("[Worker %d] Processing: %s", id, url)

		res := ScanResult{URL: url, Status: "Down"}

		client, err := GetTorClient()
		if err != nil {
			log.Printf("[Worker %d] Error creating client: %v", id, err)
			results <- res
			continue
		}

		resp, err := client.Get(url)
		if err != nil {
			log.Printf("[Worker %d] Failed to reach %s: %v", id, url, err)
			results <- res
			continue
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Printf("[Worker %d] Failed to parse HTML for %s", id, url)
			res.Status = "Up (HTML Error)"
			results <- res
			continue
		}

		html, _ := doc.Html()
		res.HTML = html
		res.Status = "Up"

		ctx, cancel, err := NewChromeDPContext(browserCfg)
		if err == nil {
			safeName := url
			safeName = strings.TrimPrefix(safeName, "http://")
			safeName = strings.TrimPrefix(safeName, "https://")
			safeName = strings.TrimSuffix(safeName, "/")
			safeName = strings.ReplaceAll(safeName, ".", "_")

			screenshotPath := fmt.Sprintf("screenshots/%s.png", safeName)

			var buf []byte
			err = chromedp.Run(ctx,
				chromedp.Navigate(url),
				chromedp.Sleep(2*time.Second),
				chromedp.FullScreenshot(&buf, 90),
			)
			if err == nil {
				SaveDataToFile(screenshotPath, buf)
				res.ScreenshotFile = screenshotPath
			}
			cancel()
		}

		results <- res
	}
}
