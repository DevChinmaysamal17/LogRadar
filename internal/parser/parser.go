// Phase - 2
// Event parsing — raw line, LogEvent struct (timestamp, source_ip, event_type, username, status)

package parser

import (
	"encoding/json"
	"time"
)

type LogEvent struct {
	Timestamp time.Time `json:"timestamp"`
	SourceIP  string    `json:"source_ip"`
	EventType string    `json:"event_type"`
	Username  string    `json:"username"`
	Status    string    `json:"status"`
}

func Parse(line string) (LogEvent, error) {
	var event LogEvent
	err := json.Unmarshal([]byte(line), &event)
	return event, err

}
