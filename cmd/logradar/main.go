package main

import (
	"LogRadar/internal/parser"
	"LogRadar/internal/rules"
	"LogRadar/internal/tailer"
	"fmt"
	"time"
)

func main() {

	// Phase 1: Log ingestion
	lines := make(chan string)
	stop := make(chan struct{})

	go tailer.Tail("logs/sample.log", lines, stop)

	detector := rules.NewBruteForceDetector(5, 30*time.Second)

	// Phase 2: Event parsing
	for line := range lines {

		fmt.Printf("Raw log: %s", line)

		event, err := parser.Parse(line)
		if err != nil {
			fmt.Println("Parse error: ", err)
		}

		fmt.Printf("Parsed event: %+v\n", event)

		alert := detector.Check(event)
		if alert != nil {
			fmt.Printf("Alert: %s| IP: %s | Severity: %s\n", alert.RuleName, event.SourceIP, alert.Severity)
		}

	}
}
