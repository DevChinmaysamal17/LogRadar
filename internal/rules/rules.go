package rules

import (
	"LogRadar/internal/parser"
	"time"
)

// Alert represents a detected security threat.
type Alert struct {
	RuleName          string
	Triggered_at_time time.Time
	MatchedEvents     []parser.LogEvent
	Severity          string
}

// Rule defines the common behavior every detection rule must follow.
type Rule interface {
	Check(event parser.LogEvent) *Alert
}
