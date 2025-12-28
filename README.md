# 🧅 TorScraper

**TorScraper** is a specialized Cyber Threat Intelligence (CTI) tool built with **Go (Golang)**. It automates data collection from .onion services by routing all traffic anonymously through the Tor network.

## ✨ Features

* **Full-Page Screenshot:** Captures the entire page of the .onion site as a PNG file.
* **HTML Backup:** Downloads and saves the raw source code for offline analysis.
* **Concurrent Scanning:** Uses Go's worker pool pattern to scan multiple targets at once.
* **Anonymous Routing:** All traffic is strictly routed through SOCKS5 proxy (**127.0.0.1:9150**).
* **Target Management:** Reads addresses from a `targets.yaml` file.

## 🛠️ Requirements

- **Go:** 1.20+ recommended
- **Tor Browser:** Connected and running on port 9150
- **Browser:** A Chromium-based browser (Chrome, Chromium, Edge, or Brave)

## 🚀 Install

Clone the repository:

```bash
git clone [https://github.com/Ayb3rk38/TorScraper.git](https://github.com/Ayb3rk38/TorScraper.git)
cd TorScraper 
```

Install dependencies:

```bash
go mod tidy
```

💻 How to Run

Run the tool from the repository root:

```bash
go run . -f targets.yaml -w 5
```

🌍 Browser Selection (Cross-Platform)

If detection fails, set CHROME_PATH to your browser executable before running.

🐧 Linux (bash/zsh)

```bash
export CHROME_PATH="/usr/bin/google-chrome"
go run . -f targets.yaml -w 5
```

🍎 macOS (zsh)

```bash
export CHROME_PATH="/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
go run . -f targets.yaml -w 5
```

🪟 Windows (PowerShell)

```bash
$env:CHROME_PATH = "C:\Program Files\Google\Chrome\Application\chrome.exe"
go run . -f targets.yaml -w 5
```

📊 Output

The tool generates the following files in the project directory:
    screenshots/: Full-page PNG captures of the visited sites.
    results.json: Structured data containing the scraped HTML and URLs.
    scan_report.log: Detailed log file including active/passive status

Created for educational purposes and Cyber Threat Intelligence research.
⚠️ Disclaimer

This project is intended for learning and local testing. Always respect a website's Terms of Service and legal boundaries.
