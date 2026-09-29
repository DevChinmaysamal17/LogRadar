package parser

import (
	"testing"
	"time"
)

// go test scans every file ending in _test.go/
// and the function which we have test must start with Test

func Test_parser(t *testing.T) {

	sample_input := `{"timestamp":"2026-09-27T10:00:00Z","source_ip":"192.168.1.10","event_type":"login_failed","username":"admin","status":"failed"}`

	event, err := Parse(sample_input)

	if err != nil {
		t.Fatalf("Parse failed entirely: %v", err)
	}

	expected_timestamp := time.Date(2026, time.September, 27, 10, 0, 0, 0, time.UTC)
	if event.Timestamp != expected_timestamp {
		t.Errorf("Timestamp is wrong! wanted %v, and got this %v", expected_timestamp, event.Timestamp)
	}

	if event.SourceIP != "192.168.1.10" {
		t.Errorf("IP is wrong! wanted '192.168.1.10', and got this %v", event.SourceIP)
	}

	if event.EventType != "login_failed" {
		t.Errorf("EventType is wrong! wanted 'login_failed', and got this %v", event.EventType)
	}

	if event.Username != "admin" {
		t.Errorf("Username is wrong! wanted 'admin', and got this %v", event.Username)
	}

	if event.Status != "failed" {
		t.Errorf("Status is wrong! wanted 'failed', and got this %v", event.Status)
	}
}
