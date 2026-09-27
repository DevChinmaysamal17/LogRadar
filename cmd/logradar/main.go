package main

import (
	"LogRadar/internal/parser"
	"LogRadar/internal/tailer"
	"fmt"
)

func main() {

	// Phase 1: Log ingestion
	lines := make(chan string)
	stop := make(chan struct{})

	go tailer.Tail("logs/sample.log", lines, stop)

	// Phase 2: Event parsing
	for line := range lines {

		fmt.Printf("Raw log: %s", line)

		event, err := parser.Parse(line)

		if err != nil {
			fmt.Println("Parse error:", err)
			continue
		}

		fmt.Printf("Parsed event: %+v\n", event)
	}
}
