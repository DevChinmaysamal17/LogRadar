package rules

import "LogRadar/internal/parser"

// a worker func receives new events and check using Check() func,
// if any alerts, then sends the alert to channel "alerts"
func Worker(rule Rule, events <-chan parser.LogEvent, alerts chan<- Alert) {
	for event := range events {
		alert := rule.Check(event)

		if alert != nil {
			alerts <- *alert
		}
	}
}
