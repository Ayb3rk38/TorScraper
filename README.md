🧅 TorScraper

TorScraper is a Go-based Cyber Threat Intelligence (CTI) tool designed to automate data collection from .onion services. It ensures anonymity by routing traffic through the Tor network and utilizes Go's concurrency model (Goroutines) for efficient multi-target scanning.

✨ Key Features

    🛡️ Anonymous Routing: Securely directs all HTTP traffic through the local SOCKS5 proxy (127.0.0.1:9150/9050) to prevent IP leaks.

⚡ Concurrent Scanning: Leverages Goroutines to process multiple onion addresses simultaneously, as recommended for high-performance CTI collection.

📸 Evidence Collection: Automatically captures site status, raw HTML content, and full-page screenshots for intelligence reporting.

📂 Target Management: Cleanses and processes bulk onion URL lists provided via a targets.yaml file.

📊 Structured Output: Exports gathered data into a structured format and generates a scan_report.log for status tracking.

🛠️ Prerequisites

    Go 1.20+ 

Tor Service / Tor Browser (Running on port 9150 or 9050)

    Chrome/Chromium (Required for screenshot capture functionality)

🚀 Installation & Usage

    Clone the repository:
    Bash

git clone https://github.com/Ayb3rk38/TorScraper.git
cd TorScraper

Install dependencies:
Bash

go mod tidy

Run the tool: The program accepts a target file and a worker count as flags:

Bash

go run . -f targets.yaml -w 5

Build the binary: To generate a compiled executable as required by the project specifications:

Bash

    go build -o TorScraper .

📊 Outputs

Following a successful scan, the tool generates the following somatized outputs:

    screenshots/: Directory containing PNG captures of the visited onion sites.

results.json: A structured file containing URLs and their associated HTML content.

scan_report.log: A log file detailing which URLs were active or unreachable.

⚠️ Disclaimer

This project is developed for educational purposes and authorized security research only. The developer is not responsible for any misuse of this tool.
