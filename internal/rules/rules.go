package rules

import (
	"LogRadar/internal/parser"
	"time"
)

// making a alert struct

type Alert struct {
	RuleName          string
	Triggered_at_time time.Time
	MatchedEvents     []parser.LogEvent
	Severity          string
}

type Rule interface {
	Check(event parser.LogEvent) *Alert
}
