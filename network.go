package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/proxy"
)

const (
	TorProxyAddr = "127.0.0.1:9150"
	DefaultUA    = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"
)

func GetTorClient() (*http.Client, error) {
	dialer, err := proxy.SOCKS5("tcp", TorProxyAddr, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Tor proxy: %v", err)
	}

	transport := &http.Transport{
		Dial:            dialer.Dial,
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}, nil
}

func IsTorActive() (bool, error) {
	client, err := GetTorClient()
	if err != nil {
		return false, err
	}

	req, _ := http.NewRequest("GET", "https://check.torproject.org/", nil)
	req.Header.Set("User-Agent", DefaultUA)

	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("Tor service is not reachable: %v", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return false, err
	}

	status := doc.Find("title").Text()
	if strings.Contains(status, "Congratulations") {
		fmt.Println("[SUCCESS]", status)
		return true, nil
	} else {
		fmt.Println("[WARNING]", status)
		return false, nil
	}
}
