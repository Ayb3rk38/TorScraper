🧅 TorScraper

TorScraper is a Cyber Threat Intelligence (CTI) tool built with Go (Golang) to automate data collection from .onion services. It routes traffic anonymously through the Tor network and processes multiple targets simultaneously using Go's concurrency model.
✨ Features

    🛡️ Anonymous Routing: Prevents IP leaks by directing all HTTP traffic through the local SOCKS5 proxy (127.0.0.1:9150).

    ⚡ Concurrent Scanning: Achieves high efficiency by scanning multiple URLs at once using the worker pool pattern.

    📸 Evidence Collection: Automatically captures site status, raw HTML content, and full-page screenshots.

    📂 Target Management: Reads, cleanses, and processes onion addresses directly from a targets.yaml file.

    📊 Structured Output: Saves all gathered data into results.json and logs the entire process in scan_report.log.

🛠️ Prerequisites

    Go 1.20+

    Tor Browser (Active on port 9150)

    Chrome/Chromium (Required for screenshot capture)

🚀 Usage

    Install dependencies:
    Bash

go mod tidy

Run the tool:
Bash

go run . -f targets.yaml -w 5

Build the binary (Optional):
Bash

    go build -o tor-scraper .

⚠️ Disclaimer

This project is developed for educational purposes and authorized security research only. Any misuse of this tool is the responsibility of the user.
