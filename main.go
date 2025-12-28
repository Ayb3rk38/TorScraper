package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"sync"
)

func main() {
	filePath := flag.String("f", "targets.yaml", "Path to targets file")
	workers := flag.Int("w", 3, "Number of concurrent workers")
	flag.Parse()

	_ = EnsureDir("screenshots")
	logFile, err := SetupLogger("scan_report.log")
	if err != nil {
		fmt.Printf("Critical Error: %v\n", err)
		return
	}
	defer logFile.Close()

	active, err := IsTorActive()
	if !active || err != nil {
		fmt.Println("CRITICAL: Tor is not active! Please start Tor Browser.")
		return
	}

	links, err := LoadTargets(*filePath)
	if err != nil {
		fmt.Printf("Error loading targets: %v\n", err)
		return
	}

	jobs := make(chan string, len(links))
	results := make(chan ScanResult, len(links))
	var wg sync.WaitGroup

	browserCfg := defaultBrowserConfig()

	for i := 1; i <= *workers; i++ {
		wg.Add(1)
		go Worker(i, jobs, results, &wg, browserCfg)
	}

	for _, link := range links {
		jobs <- link
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	var finalResults []ScanResult
	for res := range results {
		finalResults = append(finalResults, res)
		statusMsg := fmt.Sprintf("[INFO] Finished: %s -> %s\n", res.URL, res.Status)

		fmt.Println(statusMsg)
		log.Println(statusMsg)
	}

	jsonData, err := json.MarshalIndent(finalResults, "", "\t")
	if err == nil {
		_ = SaveDataToFile("results.json", jsonData)
	}

	fmt.Println("Scan completed. Check results.json and scan_report.log")
}
