// Phase 3 - Rule engine — Detect suspicious patterns using configurable rules.

package rules

import (
	"LogRadar/internal/parser"
	"sync"
	"time"
)

type BruteForceDetector struct {
	Threshold int
	Window    time.Duration
	Attempts  map[string][]parser.LogEvent
	mu        sync.Mutex
}

func NewBruteForceDetector(threshold int, window time.Duration) *BruteForceDetector {
	return &BruteForceDetector{
		Attempts:  make(map[string][]parser.LogEvent),
		Threshold: threshold,
		Window:    window,
	}
}

func (d *BruteForceDetector) Check(event parser.LogEvent) *Alert {
	d.mu.Lock()
	defer d.mu.Unlock()
	// 	New login event
	//       |
	// Is it login_failed?
	//       | yes
	// Store the event for that IP
	//       |
	// Remove old attempts
	//       |
	// Count recent attempts
	//       |
	// 	5 or more?
	//    /      \
	//  YES       NO
	//   |        |
	// Alert     nil
	if event.EventType != "login_failed" {
		return nil
	}

	ip := event.SourceIP

	//Store the failed attempts in d
	d.Attempts[ip] = append(d.Attempts[ip], event)

	// Remove attempts outside the time window
	cutoff := event.Timestamp.Add(-d.Window)

	validAttempts := []parser.LogEvent{}

	for _, attempt := range d.Attempts[ip] {
		if attempt.Timestamp.After(cutoff) {
			validAttempts = append(validAttempts, attempt)
		}
	}

	d.Attempts[ip] = validAttempts

	// Check threshold
	if len(validAttempts) >= d.Threshold {
		return &Alert{
			RuleName:          "Brute Force",
			Triggered_at_time: event.Timestamp,
			MatchedEvents:     validAttempts,
			Severity:          "HIGH",
		}
	}
	return nil

}
