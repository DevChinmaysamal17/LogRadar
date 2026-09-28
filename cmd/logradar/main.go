package main

import (
	"LogRadar/internal/alerting"
	"LogRadar/internal/metrics"
	"LogRadar/internal/parser"
	"LogRadar/internal/rules"
	"LogRadar/internal/tailer"
	"fmt"
	"time"
)

func main() {

	// Phase 5: Start metrics server on port 9000
	go metrics.Serve("9000")

	// Phase 1: Log ingestion
	lines := make(chan string)
	stop := make(chan struct{})

	go tailer.Tail("logs/sample.log", lines, stop)

	// Channels for Phase 4, making channels for parser and rules
	events := make(chan parser.LogEvent)
	alerts := make(chan rules.Alert)

	detector := rules.NewBruteForceDetector(5, 30*time.Second)

	go alerting.Printer(alerts)

	go rules.Worker(detector, events, alerts)

	// Phase 2 and 3: Event parsing and Rule engine
	for line := range lines {

		fmt.Printf("\nRaw log: %s", line)

		event, err := parser.Parse(line)
		if err != nil {
			fmt.Println("\nParse error: \n", err)
			continue
		}

		// Phase 5: Incrementing error count in prometheus
		metrics.EventsProcessed.Inc()

		fmt.Printf("\nParsed event: \n%+v\n", event)

		events <- event

	}
}
