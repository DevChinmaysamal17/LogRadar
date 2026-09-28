# LogRadar

A Go-based security monitoring service that continuously tails application logs, parses events, detects suspicious activity using configurable rules, and exposes metrics for Prometheus.

## Purpose

LogRadar is a **learning-depth portfolio project**, not a novel security product.

The goal is to build hands-on understanding of:

- Go concurrency — goroutines, channels, worker patterns
- Continuous log ingestion and parsing
- Rule-based security detection
- Configuration-driven applications
- Observability with Prometheus
- Testing and containerized deployment

---

## Status

🚧 **In Progress**

### Phase 1 — Log Ingestion ✅
Continuously tails the log file and sends newly added lines through a Go channel.

### Phase 2 — Event Parsing ✅
Converts raw JSON log lines into structured `LogEvent` objects.

### Phase 3 — Rule Engine ✅
Processes parsed events through detection rules. Currently includes configurable brute-force detection.

### Phase 4 — Alerting ✅
Generates alerts when a rule is triggered and sends them through an alert channel to the terminal printer.

### Phase 5 — Prometheus Metrics ✅
Tracks processed events and triggered alerts and exposes them through a `/metrics` HTTP endpoint.

### Phase 6 — Configuration ✅
Moves detection settings such as threshold and time window into `configs/configs.yml`, so rules can be changed without modifying Go code.

### Phase 7 — Testing ⬜
Add unit tests for the parser, configuration loader, rules, and other core components.

### Phase 8 — Docker & Deployment ⬜
Containerize LogRadar and prepare it for deployment with a production-style runtime setup.

---

## Architecture

```text
                    configs/configs.yml
                            │
                            ▼
                     Config Loader
                            │
                            ▼
                      Rule Config
                            │
                            ▼
┌─────────────┐      ┌──────────────┐      ┌──────────────┐
│  Log File   │ ───► │    Tailer    │ ───► │ lines channel│
└─────────────┘      │  Goroutine   │      └──────┬───────┘
                     └──────────────┘             │
                                                 ▼
                                          ┌──────────────┐
                                          │    Parser    │
                                          └──────┬───────┘
                                                 │
                                                 ▼
                                          events channel
                                                 │
                                                 ▼
                                      ┌───────────────────┐
                                      │    Rule Worker    │
                                      │ Brute Force Rule  │
                                      └─────────┬─────────┘
                                                │
                                                ▼
                                         alerts channel
                                                │
                                                ▼
                                      ┌───────────────────┐
                                      │   Alert Printer   │
                                      └───────────────────┘


                  ┌──────────────────────────┐
                  │     Prometheus Metrics  │
                  │ Events + Alerts counters│
                  └────────────┬─────────────┘
                               │
                               ▼
                         /metrics :9000
```

### Core data flow

```text
Log File
   ↓
Tailer
   ↓
Raw log channel
   ↓
JSON Parser
   ↓
LogEvent
   ↓
Rule Worker
   ↓
Detection Rule
   ↓
Alert
   ↓
Alert Printer
```

---

## Tech Stack

- **Go** — Core language; goroutines, channels, structs, interfaces, error handling and concurrency
- **`bufio` + `os`** — Continuously read and tail the application log file
- **`encoding/json`** — Convert JSON log lines into structured events
- **Go Channels** — Move logs, events and alerts between concurrent components
- **`time`** — Handle configurable detection time windows
- **YAML (`gopkg.in/yaml.v3`)** — Load external application/rule configuration
- **Prometheus `client_golang`** — Instrument LogRadar and expose application metrics
- **Prometheus** *(currently used)* — Scrape and store LogRadar metrics
- **Grafana** *(future)* — Visualize Prometheus metrics through dashboards
- **Docker** *(future)* — Containerize LogRadar
- **GitHub Actions** *(future)* — Automate testing and CI/CD

---

## Project Structure

```text
LogRadar/
├── cmd/
│   └── logradar/
│       └── main.go              # Application entry point
│
├── internal/
│   ├── alerting/
│   │   └── printer.go           # Prints triggered alerts
│   │
│   ├── config/
│   │   └── config.go            # Loads YAML configuration
│   │
│   ├── metrics/
│   │   └── metrics.go           # Prometheus counters + /metrics
│   │
│   ├── parser/
│   │   └── parser.go            # JSON log -> LogEvent
│   │
│   ├── rules/
│   │   ├── brute_force.go       # Brute-force detection
│   │   ├── rules.go             # Rule and Alert definitions
│   │   └── workers.go           # Concurrent rule worker
│   │
│   └── tailer/
│       └── tailer.go            # Continuous log tailing
│
├── configs/
│   └── configs.yml              # External rule configuration
│
├── logs/
│   └── sample.log               # Sample application logs
│
├── tests/                       # Test files/data
├── prometheus.yml               # Prometheus scrape configuration
├── go.mod
├── go.sum
└── README.md
```

---

## Configuration

Detection settings are stored outside the Go source code:

```yaml
brute_force:
  threshold: 5
  window: 30s
```

This means LogRadar triggers a brute-force alert when the configured threshold of failed login attempts occurs within the configured time window.

Changing:

```yaml
threshold: 5
```

to:

```yaml
threshold: 10
```

does not require changing the Go detection code.

## Prometheus Metrics

LogRadar exposes metrics at:

```text
http://localhost:9000/metrics
```

Currently tracked metrics include:

```text
logradar_events_processed_total
logradar_alerts_triggered_total
```

These can be scraped by Prometheus for monitoring and analysis.

Grafana dashboards are planned for a future phase.

## Running Locally

Clone the repository:

```bash
git clone https://github.com/DevChinmaysamal17/LogRadar.git
cd LogRadar
```

Run the application:

```bash
go run ./cmd/logradar
```

In another terminal, append a new log:

```bash
echo '{"timestamp":"2026-09-27T10:00:00Z","source_ip":"192.168.1.10","event_type":"login_failed","username":"admin","status":"failed"}' >> logs/sample.log
```

Multiple failed login events from the same IP within the configured window can trigger the brute-force detection rule.

## Scope

### Current

- Single log-file input
- JSON log parsing
- Configurable brute-force detection
- Concurrent rule worker
- Terminal alerts
- Prometheus metrics
- YAML-based configuration

### Future

- Unit and integration testing
- Multiple log sources
- Additional detection rules
- Grafana dashboards
- Docker deployment
- CI/CD with GitHub Actions
- Webhook/Slack notifications
- ML-based anomaly detection
- Web dashboard
- Configuration UI

## Learning Focus

LogRadar is primarily designed to demonstrate practical understanding of:

```text
Go
├── Goroutines
├── Channels
├── Worker patterns
├── Structs
├── Interfaces
├── Error handling
└── Concurrency

Backend / Systems
├── Log ingestion
├── Event processing
├── Rule engines
├── Configuration management
└── Observability

DevOps
├── Prometheus, Grafana
├── Docker (planned)
└── CI/CD (planned)
```