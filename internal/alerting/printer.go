package alerting

import (
	"LogRadar/internal/rules"
	"fmt"
)

// receives signal from alerts(workers.go) and print alert

func Printer(alerts <-chan rules.Alert) {
	for receive_alert := range alerts {
		fmt.Println("\nALERT: ")
		fmt.Println("Rule: ", receive_alert.RuleName)
		fmt.Println("Trigger Time: ", receive_alert.Triggered_at_time)
		fmt.Println("Severity: ", receive_alert.Severity)
		fmt.Println("Matched Events:", len(receive_alert.MatchedEvents))
		fmt.Println("Source IP:", receive_alert.MatchedEvents[0].SourceIP)
		fmt.Println("Username:", receive_alert.MatchedEvents[0].Username)
		fmt.Println("Event:", receive_alert.MatchedEvents[0].EventType)
	}
}
