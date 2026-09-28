package main

import (
	"LogRadar/internal/alerting"
	"LogRadar/internal/config"
	"LogRadar/internal/metrics"
	"LogRadar/internal/parser"
	"LogRadar/internal/rules"
	"LogRadar/internal/tailer"
	"fmt"
	"time"
)

func main() {

	// PHASE 6: CONFIGURATION
	// Load threshold and time-window settings from YAML.
	conf, err := config.Load("configs/configs.yml")

	if err != nil {
		fmt.Println(err)
		return
	}

	// PHASE 5: PROMETHEUS
	// Start metrics server on port 9000 in a goroutine.
	// Metrics are available at /metrics.
	go metrics.Serve("9000")

	// PHASE 1: LOG INGESTION
	// 'lines' carries raw log strings from the tailer.
	lines := make(chan string)

	// 'stop' is used to signal the tailer to stop.
	stop := make(chan struct{})

	// Start continuously watching the log file.
	go tailer.Tail("logs/sample.log", lines, stop)

	// PHASE 4: CHANNELS
	// events: parsed logs → Rule Worker
	// alerts: Rule Worker → Alert Printer
	events := make(chan parser.LogEvent)
	alerts := make(chan rules.Alert)

	// Convert YAML value "30s" into time.Duration.
	window, err := time.ParseDuration(conf.BruteForce.Window)

	if err != nil {
		fmt.Println("Invalid window:", err)
		return
	}

	// Create the brute-force detector using configured values.
	detector := rules.NewBruteForceDetector(
		conf.BruteForce.Threshold,
		window,
	)

	// Start alert printer to display triggered alerts.
	go alerting.Printer(alerts)

	// Start rule worker to check parsed events.
	go rules.Worker(detector, events, alerts)

	// PHASE 2 + 3: PARSING + RULE ENGINE
	// Receive raw logs, parse them, and send events to the worker.
	for line := range lines {

		// Show the original log line.
		fmt.Printf("\nRaw log: %s", line)

		// Convert raw JSON into a structured LogEvent.
		event, err := parser.Parse(line)

		// Skip invalid/malformed log entries.
		if err != nil {
			fmt.Println("\nParse error: \n", err)
			continue
		}

		// PHASE 5: Count successfully parsed events.
		metrics.EventsProcessed.Inc()

		// Show the parsed event for debugging.
		fmt.Printf("\nParsed event: \n%+v\n", event)

		// Send the parsed event to the Rule Worker.
		events <- event
	}
}
