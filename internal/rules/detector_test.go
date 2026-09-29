package rules

// go test scans every file ending in _test.go/
// and the function which we have test must start with Test

import (
	"LogRadar/internal/parser"
	"testing"
	"time"
)

func Test_BruteForceDetector(t *testing.T) {
	detector := NewBruteForceDetector(5, 30*time.Second)

	baseTime := time.Date(2026, time.September, 27, 10, 0, 0, 0, time.UTC)

	makeEvent := func(ip string, secondsOffset int) parser.LogEvent {
		return parser.LogEvent{
			Timestamp: baseTime.Add(time.Duration(secondsOffset) * time.Second),
			SourceIP:  ip,
			EventType: "login_failed",
			Username:  "admin",
			Status:    "failed",
		}
	}

	for i := 0; i < detector.Threshold-1; i++ {
		alert := detector.Check(makeEvent("192.168.1.10", i))

		if alert != nil {
			t.Fatalf("expected no alert at attempt %d, got alert", i+1)
		}
	}
	alert := detector.Check(
		makeEvent("192.168.1.10", detector.Threshold-1),
	)

	if alert == nil {
		t.Fatalf("expected alert after %d failed logins", detector.Threshold)
	}

	for i := detector.Threshold; i < detector.Threshold+3; i++ {

		alert := detector.Check(
			makeEvent("192.168.1.10", i),
		)
		if alert == nil {
			t.Fatalf(
				"unexpected repeated alert at attempt %d",
				i+1,
			)
		}
	}

}
