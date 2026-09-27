# LogRadar

A Go-based security monitoring service that tails application logs, detects suspicious activity using configurable rules, and exposes metrics for Prometheus/Grafana.

## Purpose

This is a learning-depth project, not a novel product. Log-based security monitoring is already solved at scale by tools like Splunk, Wazuh, Datadog, and fail2ban. LogRadar exists to build real, hands-on understanding of:
- Go concurrency (goroutines, channels, worker pools, mutexes)
- Log ingestion and parsing
- Rule-based detection logic
- Observability tooling (Prometheus/Grafana)

## Status: 🚧 In Progress

- [x] Phase 1 — Log ingestion (file tailer)
- [x] Phase 2 — Event parsing (JSON log lines → structured events)
- [x] Phase 3 — Brute-force login detection rule
- [ ] Phase 4 — Worker pool (concurrent rule evaluation) — in progress
- [ ] Phase 5 — Prometheus `/metrics` endpoint
- [ ] Phase 6 — Config file support (rules become configurable)
- [ ] Phase 7 — Dockerize + CI/CD via GitHub Actions

## Architecture

[Log File] --tail--> [Ingest goroutine] --channel--> [Rule Engine (worker pool)] --> [Alert Sink (stdout)]
|
v
[Metrics Store] --> [/metrics HTTP endpoint] --> Prometheus


## Tech Stack

- **Go** — goroutines, channels, worker pools, `sync.Mutex`
- **`bufio`/`os`** — file tailing
- **`encoding/json`** — log event parsing
- **`prometheus/client_golang`** *(planned)* — metrics exposition
- **Prometheus + Grafana** *(planned)* — monitoring/dashboards, self-hosted
- **Docker + GitHub Actions** *(planned)* — containerization and CI/CD

## Project Structure

logradar/
├── cmd/logradar/main.go
├── internal/
│ ├── tailer/ # file tailing
│ ├── parser/ # log line → LogEvent
│ ├── rules/ # detection rules (brute force, rate spike)
│ ├── worker/ # concurrent rule evaluation (WIP)
│ ├── metrics/ # Prometheus metrics (planned)
│ └── config/ # rule configuration (planned)
├── configs/config.yaml
├── logs/sample.jsonl
└── tests/testdata/


## Running Locally

```bash
git clone https://github.com/DevChinmaysamal17/LogRadar.git
cd LogRadar
go run cmd/logradar/main.go
```

In a separate terminal, append log lines to trigger detection:
```bash
echo '{"timestamp":"2026-09-27T10:00:00Z","source_ip":"192.168.1.10","event_type":"login_failed","username":"admin","status":"failed"}' >> logs/sample.jsonl
```

5 failed logins from the same IP within 30 seconds triggers a brute-force alert.

## Scope (v1)

**In:** single log file input, hardcoded brute-force + rate-spike rules, stdout alerts, Prometheus metrics endpoint.

**Out (future):** multi-source ingestion, web dashboard, ML anomaly detection, Slack/webhook alerts, config UI.