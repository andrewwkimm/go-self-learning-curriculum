## Philosophy

Every module follows the same loop: **concept intro → read code → write code → knowledge check.** Exercises are small and targeted until end-of-phase projects consolidate everything. The course skews toward idiomatic Go and real industry patterns — not toy examples, not deep theory.

Learning happens in Claude Code or Copilot: paste exercises, get feedback on your actual code in your editor.

---

## Learning Methodology

This curriculum applies four evidence-backed techniques. Every module is structured around all four — they are not optional enrichment, they are the mechanism by which things stick.

### 1. The Five-Stage Exercise Protocol

Every coding exercise follows this sequence, in order, without skipping:

```
ATTEMPT   →  no references, no notes, genuinely stuck ≥10 min before asking for help
SUBMIT    →  paste your attempt — broken, incomplete, or wrong is fine and expected
FEEDBACK  →  critique against correctness + Uber style guide + idiomatic Go
REWRITE   →  close the feedback, rewrite from scratch without referencing it
CARD      →  add one Anki card capturing the core insight from this exercise
```

The 10-minute stuck rule is not arbitrary. Struggling with a problem before instruction produces better long-term retention than instruction-first, even when the attempt fails completely (Kapur, productive failure research). The discomfort of not remembering is the mechanism, not a sign something is wrong.

For knowledge check questions (conceptual, not write-code): attempt your own answer in writing → get feedback → restate the corrected understanding in your own words without looking at the feedback. Same principle, adapted format.

### 2. Anki — Spaced Repetition

After every module, make Anki cards for each core concept. The protocol:

- **You propose cards first.** I review whether they're well-formed, comprehensive, and targeting the right thing.
    
- **I suggest additional cards** you may have missed, especially for edge cases and concept relationships.
    
- **Card quality rules:** cards must force production, not recognition. "What is the syntax for X" is a bad card — it tests recognition of a fact. "Write a function that does X using Y and Z" is a good card — it tests retrieval of a skill. Cards that combine concepts from multiple modules are the highest value.
    
- **Review daily.** 15 minutes of Anki beats 2 hours of re-reading every time.
    

### 3. Interleaving — Integration Exercises

Each module's exercises are split into two types:

- **Core exercises** — isolate and practice only the new concept. Used early in a module before the concept is solid.
    
- **Integration exercises** — combine the new concept with 2-3 prior modules. These are the primary learning mechanism, not a bonus. They force the brain to retrieve earlier concepts while applying new ones, which is how isolated knowledge nodes become a connected mental model.
    

The ratio shifts as modules accumulate: Modules 1-3 are mostly core (nothing to integrate yet). By Module 8, roughly 50/50. By Phase 2, integration exercises dominate. Each module in this curriculum marks exercises with their type and lists which prior modules are required.

### 4. Feynman Assessment — Two-Point Evaluation Per Module

At two points per module — mid-module and end-of-module — you explain the concepts covered so far in your own words. No notes, no references.

- **Mid-module check:** diagnostic. Identifies gaps before they compound. Lighter in scope.
    
- **End-of-module check:** evaluative. Determines readiness to proceed to the next module.
    

Assessment intensity scales by phase:

**L1 — Modules 1-8 (Language Basics)** Explain the concept. Handle one "what if" variant. Identify one edge case. Pass/fail is clear. Goal is to confirm the mental model is correct, not to stress-test it.

**L2 — Modules 9-21 (Concurrency, I/O, Patterns, gRPC)** Explain + justify design choices + handle failure modes + compare to the alternative approach. Closer to a code review conversation. I'll push on "why did you choose X over Y" and "what breaks under Z condition."

**L3 — Phase Projects + Capstone** Full technical interview simulation. Defend your design under pressure. Handle hostile edge cases. Articulate tradeoffs with evidence. I'll push until you either hold your position with a sound argument or revise it with a better one. Silence and "I'm not sure" are not passing answers at this level — "I'm not sure, but my best reasoning is X because Y" is.

If you don't pass the end-of-module Feynman check, you don't move to the next module. You identify the specific gap, address it, and re-check. Moving forward with a broken mental model compounds into larger confusion later.

---

---

## Module Delivery Standard

Every module session must conform to this structure, in this order. This is the template any session uses to deliver a module — not a loose guideline. A module that skips or shortens any section is not being delivered correctly.

### 1. Mental Model (concept intro)

- Brief prose, not bullet lists of syntax
    
- Anchored to what the learner already knows (Python comparisons are appropriate for Modules 1-8)
    
- Focused on _what is different or non-obvious_ about Go — not a rehash of general programming concepts
    
- Ends with a clear statement of what you'll be able to do by the end of the module
    

### 2. Worked Example (read + predict)

- A complete, runnable Go program that demonstrates the module's core concepts
    
- Must include 3-5 prediction questions the learner answers _before_ running the code
    
- Questions target zero values, type inference, output order, compile behavior, or scope — things that are easy to get wrong
    
- Learner runs the program after predicting and reconciles any surprises
    

### 3. Exercises

Labeled explicitly as **[Core]** or **[Integration]** with prior modules listed for integration exercises.

Every module must include all four exercise types:

- **Predict** — given code, state what it does and why (no running until after prediction written)
    
- **Write** — produce a program or function from memory using the five-stage protocol
    
- **Vary** — modify a working solution to satisfy a new constraint
    
- **Debug** — broken code with multiple errors; identify every error and explain why it's wrong before fixing
    

The five-stage protocol applies to every Write and Debug exercise:

```
ATTEMPT   →  no references, ≥10 min genuine struggle before asking
SUBMIT    →  paste attempt even if broken or incomplete
FEEDBACK  →  correctness + Uber style guide + idiomatic Go
REWRITE   →  close feedback, rewrite from scratch
CARD      →  one Anki card capturing the core insight
```

### 4. Mid-Module Feynman Check

- Happens after core concepts are introduced, before integration exercises
    
- Learner explains covered concepts in their own words, no notes
    
- Graded at the appropriate intensity level (L1/L2/L3 per phase)
    
- Diagnostic — if gaps surface, address them before proceeding
    
- Not skippable even if exercises went well
    

### 5. Integration Exercises (Module 4 onwards)

- Combine the current module's concepts with 2-3 prior modules explicitly listed
    
- These are not optional — they are the primary mechanism for building a connected mental model
    
- Ratio grows with module count: by Phase 2, integration exercises outnumber core exercises
    

### 6. End-of-Module Feynman Check

- Evaluative — determines readiness to proceed
    
- Covers the full module including integration concepts
    
- Graded at L1/L2/L3 per phase with explicit pass/fail
    
- Failing means: identify the gap, address it, re-check before moving on
    

### 7. Anki Card Session

- Learner proposes cards first
    
- Cards are reviewed for quality: must test retrieval of a skill, not recognition of a fact
    
- Additional cards suggested for edge cases and concept relationships missed
    
- Cards combining multiple module concepts are flagged as highest priority
    

---

## Suggested Sequencing

Run these three tracks in parallel for the first stretch, then converge:

```
Track A (Go Language)      Track B (Infrastructure)     Track C (CS Foundations)
    Modules 1–8                Linux + Git (deep)           SQL (deep)
    Phase 1 Project            Docker (deep)
    Modules 9–12               Terraform basics
    Phase 2 Project
    Modules 13–16              Kubernetes (conceptual)      Redis
    Phase 3 Project            CI/CD (GitHub Actions)
    Modules 17–21
    Phase 4 Project
                    ↘               ↓                ↙
                         CAPSTONE PROJECT
```

Don't start Kubernetes until you're comfortable with Docker. Don't start Kafka until you have Redis. Everything else in Track B and C can run alongside the Go modules.

---

## Part 1: The Go Language

### Module 1 — Go Basics & Mental Model

**Goal:** Understand what makes Go different before writing a line.

Topics: what Go is and isn't (compiled, statically typed, no classes, no exceptions), the toolchain (`go run`, `go build`, `go fmt`, `go vet`), packages and `main`, variables (`var` vs `:=`, zero values), basic types (`int`, `string`, `bool`, `float64`), constants and `iota`.

Exercises:

- **Knowledge check:** "In Python you can do `x = 5` anywhere. What's the Go equivalent inside a function? Outside a function?"
    
- **Read:** Given a snippet, identify what each variable declaration style is doing and why.
    
- **Write:** Declare variables 5 different ways, print them, observe zero values.
    

---

### Module 2 — Control Flow

**Goal:** Write programs that branch and loop.

Topics: `if/else` (including the init statement: `if x := f(); x > 0`), `for` (Go's only loop — covers while, do-while, range patterns), `switch` (no fallthrough by default, type switches preview), `defer` — what it does, when it runs, stack order.

Exercises:

- **Knowledge check:** "Python has `while`, `for`, and do-while (sort of). Go only has `for`. How do you write a while loop in Go?"
    
- **Read:** A snippet using `defer` in a function — predict the output order.
    
- **Write:** FizzBuzz in Go.
    
- **Write:** A function that uses `defer` to print "done" after any return path.
    

---

### Module 3 — Functions

**Goal:** Understand Go's function model — this is where Go diverges most from Python.

Topics: function signatures, multiple return values, named return values, variadic functions (`...`), functions as first-class values, anonymous functions and closures, `init()` functions.

Exercises:

- **Knowledge check:** "Python functions return one value (or a tuple). What does Go do differently, and why does it matter for error handling?"
    
- **Read:** A function with named returns — what does a bare `return` do?
    
- **Write:** A function `divide(a, b float64) (float64, error)` that returns an error if `b` is 0.
    
- **Write:** A closure that acts as a counter — each call increments and returns the count.
    

---

### Module 4 — Types: Structs, Methods, Interfaces

**Goal:** Understand Go's type system — no classes, but powerful composition.

Topics: structs (definition, initialization, embedding), methods (value receivers vs pointer receivers — critical distinction), interfaces (implicit satisfaction, the `error` interface), type assertions and type switches, the empty interface `any` / `interface{}`, `fmt.Stringer`.

Exercises:

- **Knowledge check:** "Python uses `__init__` for constructors. Go has no constructors — what's the idiomatic pattern instead?"
    
- **Knowledge check:** "When would you use a pointer receiver vs a value receiver?"
    
- **Read:** A struct with both value and pointer receiver methods — predict which ones mutate state.
    
- **Write:** A `Shape` interface with `Area() float64` and `Perimeter() float64`, implemented by `Rectangle` and `Circle`.
    
- **Write:** A `String()` method on a struct so it prints nicely with `fmt.Println`.
    

---

### Module 5 — Collections: Arrays, Slices, Maps

**Goal:** Use Go's core data structures correctly and avoid common gotchas.

Topics: arrays (fixed size, rarely used directly), slices (backing arrays, `len`, `cap`, `append`, `copy`), slice gotchas (shared backing arrays, nil vs empty slice), maps (declaration, zero values, checking key existence), iterating with `range`, structs as map values.

Exercises:

- **Knowledge check:** "What's the difference between `var s []int` and `s := []int{}`? Does it matter?"
    
- **Read:** A function that appends to a slice passed as a parameter — does the caller see the change? Why or why not?
    
- **Write:** A function that takes a `[]string` and returns a `map[string]int` of word frequencies.
    
- **Write:** Remove duplicates from a `[]int` using a map.
    

---

### Module 6 — Pointers

**Goal:** Understand pointers well enough to use them correctly, not fear them.

Topics: `&` and `*` operators, when Go passes by value vs reference, pointers to structs, `new()` vs `&T{}`, when not to use pointers.

Exercises:

- **Knowledge check:** "Python passes objects by reference and primitives by value (sort of). How does Go handle this?"
    
- **Read:** Two functions — one takes `*int`, one takes `int` — predict which modifies the caller's variable.
    
- **Write:** A function `increment(n *int)` and demonstrate calling it correctly.
    
- **Write:** Rewrite a struct method to use a pointer receiver and show the difference.
    

---

### Module 7 — Error Handling

**Goal:** Handle errors the Go way — this is non-negotiable for industry code.

Topics: the `error` interface, `errors.New`, `fmt.Errorf`, wrapping errors with `%w` and `errors.Is` / `errors.As`, sentinel errors, custom error types, when to `panic` vs return error, `recover`.

Exercises:

- **Knowledge check:** "Python uses try/except. Go has no exceptions. What's the tradeoff?"
    
- **Read:** A chain of functions passing errors up — trace the error path.
    
- **Write:** A custom error type `ValidationError` with a field and a message.
    
- **Write:** A function that wraps an error with context using `fmt.Errorf("%w")`.
    

---

### Module 8 — Packages & Modules

**Goal:** Structure real Go code the way industry does.

Topics: `go mod init`, `go.mod`, `go.sum`, package naming conventions, exported vs unexported identifiers, internal packages, common stdlib packages: `fmt`, `os`, `io`, `strings`, `strconv`, `math`, `sort`, `time`.

Exercises:

- **Knowledge check:** "What makes an identifier exported in Go? How is this different from Python's `_` convention?"
    
- **Write:** Split a single-file program into two packages and import one from the other.
    
- **Write:** Use `strings`, `strconv`, and `os` together to read a command-line argument, parse it as an int, and print its square.
    

---

### Phase 1 Project — CLI Task Manager

Build a CLI task manager: `tasks add "buy milk"`, `tasks list`, `tasks done 1`. Uses structs, slices, maps, file I/O, error handling, packages, `os.Args`. No frameworks — stdlib only.

---

### Module 9 — Goroutines

**Goal:** Understand Go's concurrency primitive.

Topics: `go` keyword, what a goroutine is, why goroutines are cheap (vs threads), the scheduler (M:N threading, brief), race conditions and why they're silent.

Exercises:

- **Knowledge check:** "Python has threads and the GIL. What does Go do differently?"
    
- **Read:** A program with a goroutine — predict whether the output is deterministic.
    
- **Write:** Launch 5 goroutines that each print their index — observe the chaos, then fix it.
    

---

### Module 10 — Channels

**Goal:** Coordinate goroutines using channels.

Topics: unbuffered vs buffered channels, send/receive/close, `range` over a channel, `select` statement, directional channels (`chan<-`, `<-chan`), common patterns (fan-out, fan-in, done channels).

Exercises:

- **Knowledge check:** "What happens if you send to a closed channel? What about receive?"
    
- **Read:** A pipeline pattern — trace data flow through 3 channels.
    
- **Write:** A worker pool: N goroutines processing jobs from a channel.
    
- **Write:** A timeout using `select` and `time.After`.
    

---

### Module 11 — Sync Primitives

**Goal:** Know when channels aren't the right tool.

Topics: `sync.Mutex` and `sync.RWMutex`, `sync.WaitGroup`, `sync.Once`, `sync.Map` (and when not to use it), `atomic` package basics, `errgroup` (the idiomatic extension of `WaitGroup` — cancel the group on first error).

Exercises:

- **Knowledge check:** "When would you use a mutex instead of a channel?"
    
- **Read:** A concurrent counter — identify the race condition, then the fix.
    
- **Write:** A safe concurrent cache using `sync.RWMutex`.
    
- **Write:** Use `sync.WaitGroup` to wait for N goroutines to finish.
    

---

### Module 12 — Context

**Goal:** Propagate cancellation and deadlines through call chains.

Topics: `context.Background()`, `context.TODO()`, `context.WithCancel`, `context.WithTimeout`, `context.WithDeadline`, passing context through call chains, checking `ctx.Done()`, context values (and why to use them sparingly).

Exercises:

- **Knowledge check:** "Why does almost every Go function in industry take a `context.Context` as its first argument?"
    
- **Read:** A function chain passing context — trace what happens when the context is cancelled.
    
- **Write:** A function that does work in a goroutine and respects cancellation via context.
    

---

### Phase 2 Project — Concurrent URL Checker

Takes a list of URLs, checks them in parallel with a worker pool, respects a timeout per request, prints results with status codes and latency. Uses goroutines, channels, WaitGroup, context, HTTP client.

---

### Module 13 — I/O & Files

Topics: `os.Open`, `os.Create`, `os.ReadFile`, `os.WriteFile`, `bufio.Scanner` for line-by-line reading, `io.Reader` and `io.Writer` interfaces, `io.Copy`, `io.TeeReader`, working with stdin/stdout/stderr.

---

### Module 14 — JSON & Encoding

Topics: `encoding/json` (`Marshal`, `Unmarshal`), struct tags (`json:"name"`, `omitempty`, `-`), streaming (`json.Decoder`, `json.Encoder`), custom `MarshalJSON` / `UnmarshalJSON`, brief: `encoding/csv`.

Exercises:

- **Write:** Parse a JSON API response into a struct, extract a field, print it.
    
- **Read:** A struct with tags — predict the JSON output.
    

---

### Module 15 — HTTP Client & Server

Topics: `net/http` for GET/POST requests, reading response bodies correctly (always close), building a simple HTTP server with `http.HandleFunc`, `http.ServeMux`, middleware pattern, JSON request/response handling.

Exercises:

- **Write:** A server with two routes: `GET /health` and `POST /echo`.
    
- **Write:** A client that calls a public API and parses the response.
    

---

### Module 16 — Databases

Topics: `database/sql` interface, drivers (pgx for Postgres, sqlite for practice), `Query`, `QueryRow`, `Exec`, scanning rows into structs, transactions, prepared statements. Then: `sqlx` (thin wrapper, struct scanning, widely adopted), then `sqlc` (code generation from `.sql` files — understand what it generates and why). Brief: GORM tradeoffs.

The stack: write raw `database/sql` first → use `sqlx` to reduce boilerplate → see how `sqlc` generates the equivalent code automatically.

Exercises:

- **Write:** Create a table, insert rows, query them back — all in Go with `database/sql`.
    
- **Knowledge check:** "Why does `database/sql` use an interface rather than a concrete type?"
    
- **Write:** The same queries using `sqlx` struct scanning — compare the difference in code volume.
    

---

### Phase 3 Project — REST API with Persistence

A simple CRUD REST API (e.g., notes or todos) backed by SQLite, with JSON in/out, proper error handling, and context propagation on every DB call. Uses HTTP server, JSON encoding, `database/sql` + `sqlx`, error handling, context.

---

### Module 17 — Testing

Topics: `testing` package (`TestXxx`, `t.Error`, `t.Fatal`), table-driven tests (the Go standard), subtests (`t.Run`), benchmarks (`BenchmarkXxx`), `testify` library (`assert`, `require`, `mock`), mocking interfaces, `httptest` for testing HTTP handlers, the race detector (`go test -race`).

Exercises:

- **Write:** Table-driven tests for your `divide` function from Module 3.
    
- **Write:** A test for an HTTP handler using `httptest.NewRecorder`.
    
- **Write:** A benchmark for a string manipulation function and read the output.
    

---

### Module 18 — Generics (Go 1.18+)

Topics: type parameters (`func Map[T, U any](s []T, f func(T) U) []U`), constraints (`comparable`, `any`, custom constraints), generic data structures (brief), when to use generics vs interfaces.

Exercises:

- **Knowledge check:** "Before generics, how did Go handle functions that needed to work on multiple types?"
    
- **Write:** A generic `Filter[T]` function.
    
- **Write:** A generic `Stack[T]` with `Push`, `Pop`, `Peek`.
    

---

### Module 19 — Common Patterns & Idioms

**Style & conventions anchor for this curriculum:** all patterns here are validated against three sources in order of authority — [Effective Go](https://go.dev/doc/effective_go) (official baseline), [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md) (community codification), and the Go standard library source (ground truth). If all three agree, it's de facto standard. If only Uber says it, it's strong convention. Anything from only one source is flagged as one opinion. The builder pattern is explicitly excluded — functional options replaced it in idiomatic Go and none of the three sources endorse it.

For Module 11 sync primitives: the stance followed here is mutexes for protecting shared state in structs, channels for transferring ownership or signaling — consistent with Uber's guide and how K8s, Docker, Terraform, and LiveKit all behave in practice. Rob Pike's "share memory by communicating" is a design philosophy, not a rule to always prefer channels.

Topics:

- **Functional options pattern** — the idiomatic way to handle optional constructor parameters in Go. Endorsed by Uber (dedicated section), Dave Cheney's original blog post, and used throughout the stdlib. Zero controversy. _(Source: Effective Go + Uber + stdlib)_
    
- **Embedding for composition** — add behavior via embedding, never to satisfy an interface, never embed mutable types carelessly. Replaces inheritance from other languages. _(Source: Effective Go + Uber)_
    
- **Interface design** — keep interfaces small, define them at the point of use (in the consuming package, not the implementing package). The `io.Reader`/`io.Writer` pattern is the canonical example. _(Source: Effective Go + Uber)_
    
- `io.Reader`**/**`io.Writer` **composition** — composing behavior through standard interfaces rather than concrete types. Core to Go's design, not a style preference. _(Source: Effective Go + stdlib)_
    
- **Error wrapping patterns** — `fmt.Errorf("%w")`, `errors.Is`, `errors.As`, sentinel errors, custom error types with context. _(Source: Effective Go + Uber)_
    
- **Table-driven design** — applies beyond tests; structuring logic as data-driven tables reduces branching and makes behavior explicit. The Go team uses this pattern throughout the stdlib. _(Source: Uber + stdlib)_
    
- `slog` **for structured logging** — stdlib since Go 1.21, the standard answer for structured logs. No third-party logging library needed for new projects. _(Source: stdlib)_
    

Study each with a read-then-replicate exercise: find an example in a real OSS repo (LiveKit, cobra, or the stdlib), read it, then write your own version from scratch without referencing the original.

---

### Module 20 — Tooling & Production Habits

Topics: `go vet`, `staticcheck`, `golangci-lint`, race detector, profiling (`pprof` basics), build tags, Makefile patterns for Go projects, Docker multi-stage builds for Go, environment config patterns (`os.Getenv`, `godotenv`).

---

### Module 21 — Protobuf & gRPC

**Why it's here:** LiveKit's entire signaling layer is protobuf + gRPC. This topic also unlocks etcd, Kubernetes internals, gRPC-gateway, and most modern distributed systems codebases.

Topics: what protobuf is and why Go projects use it instead of JSON for internal comms, defining `.proto` files, generating Go code with `protoc`, a simple gRPC server and client, streaming RPCs (brief), reading generated protobuf code without confusion, schema evolution concepts (backward/forward compatibility).

Exercises:

- **Write:** Define a `.proto` file for a simple `UserService` with `GetUser` and `CreateUser` RPCs. Generate the Go code and implement the server.
    
- **Knowledge check:** "What's the difference between unary and streaming gRPC? When would you choose streaming?"
    
- **Read:** A section of LiveKit's generated protobuf Go code — identify the generated types, methods, and how they map back to the `.proto` file.
    

---

### Phase 4 Project — gRPC Metrics Service

A small gRPC service that accepts event reports (e.g., session start, session end) via protobuf, stores them in SQLite, and exposes a query endpoint. Tested with table-driven tests and `httptest`. Includes structured logging with `slog`.

---

## Part 2: Infrastructure & Supporting Skills

### SQL (Deep)

Run this in parallel with Go Phase 1-2. Backend SQL is different from data engineering SQL.

Topics: transactions and isolation levels, indexes and query planning (`EXPLAIN ANALYZE`), schema migrations (Flyway, golang-migrate), connection pooling (pgBouncer concepts), window functions, CTEs, `UPSERT` and conflict handling, normalization tradeoffs.

Exercises include reading a slow query's `EXPLAIN` output, designing a schema with appropriate indexes, and writing a migration script.

---

### Linux & the Command Line

Topics: processes and signals (`SIGTERM`, `SIGKILL`, `SIGHUP`), file descriptors, permissions and `chmod`/`chown`, `systemd` unit files, `cron`, `grep`/`awk`/`sed` pipelines, `strace` and `lsof` for debugging, `/proc` filesystem basics.

**Why it matters:** Most backend/platform work runs on Linux and you debug it from a shell. Every deployment, every container, every Kubernetes pod runs Linux.

---

### Git (Deep)

Topics: rebasing (interactive rebase, rebase vs merge), cherry-picking, `git bisect`, `git reflog` (your safety net), merge strategies, upstream sync workflow (fork → remote → rebase), writing clean commits (conventional commits), squashing for PR cleanliness.

OSS contribution specifically requires keeping a fork in sync with upstream and navigating PR review cycles cleanly.

---

### Docker (Deep)

Topics: multi-stage builds (the Go pattern — build stage + scratch/distroless final stage), layer caching optimization, `ENTRYPOINT` vs `CMD`, health checks, networking between containers, Docker Compose for local dev, image size optimization, `.dockerignore`.

Go multi-stage build pattern:

```dockerfile
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/server ./cmd/server

FROM gcr.io/distroless/static
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
```

---

### Terraform

Topics: providers and resources, `terraform init` / `plan` / `apply` / `destroy`, state files and remote state (S3, GCS), variables and outputs, modules (writing and consuming), workspaces, `data` sources, `lifecycle` rules, import existing resources, `terraform fmt` and `terraform validate`.

Relevant to the LiveKit-type role: provisioning cloud infrastructure (VMs, databases, object storage, DNS) as code so you can reproduce environments deterministically.

Exercises:

- **Write:** A Terraform module that provisions a GCS bucket with versioning and lifecycle rules.
    
- **Write:** A module that provisions a Postgres instance (Cloud SQL or RDS) with a database and user, outputting the connection string.
    
- **Write:** Use remote state in GCS so two team members can share state.
    

---

### Kubernetes (Operational)

Topics: pods, deployments, services, ingress, ConfigMaps/Secrets, resource limits and requests, liveness/readiness probes, namespaces, RBAC, `kubectl` fluency, Helm basics (consuming charts, not writing them yet), horizontal pod autoscaling.

You already use GKE at Ovative. The goal here is to go from "I know it exists" to "I can debug a failing deployment and write the manifests from scratch."

---

### Observability: Metrics, Logs, Traces

The three pillars — and one of the first things a new backend engineer is asked to add to their service.

Topics: Prometheus metrics (counters, gauges, histograms), Grafana dashboards, structured logging with `slog` (already in Module 19, but applied to infrastructure context here), OpenTelemetry (OTel) for distributed tracing — spans, traces, context propagation, exporters, instrumenting an HTTP server and a DB client.

The LiveKit JD explicitly lists OpenTelemetry as a nice-to-have. Every production backend role expects instrumented code.

---

### CI/CD (GitHub Actions)

Topics: workflow triggers, jobs and steps, caching Go module downloads, building and pushing Docker images (using OIDC, not static secrets), deploying to Kubernetes from Actions, reusable workflows, branch protection rules, required status checks.

Write a workflow that: lints with `golangci-lint` → runs `go test -race` → builds a Docker image → pushes to a registry → applies a Kubernetes manifest.

---

### Redis

Topics: data structures (strings, hashes, lists, sorted sets, sets), caching patterns (cache-aside, write-through), pub/sub, distributed locks (`SET NX PX`), rate limiting with sliding window, TTLs, persistence options (RDB vs AOF), connection pooling in Go (using `go-redis`).

LiveKit uses Redis for distributed room state across nodes — the exact use case for distributed locks and pub/sub.

---

### Kafka (Concepts + Basic Use)

Topics: topics, partitions, consumer groups, offsets, delivery semantics (at-least-once vs exactly-once vs at-most-once), producers and consumers in Go (`confluent-kafka-go` or `segmentio/kafka-go`), dead letter queues, compaction vs deletion retention.

The LiveKit JD lists Kafka/Pulsar as a nice-to-have. More broadly, Kafka is the default event streaming backbone for data-intensive backend systems.

---

### System Design (Ongoing Study)

Not a module with a finish line — a reading and reasoning practice that runs throughout the curriculum.

Core concepts: CAP theorem, consistency vs availability, idempotency, backpressure, circuit breakers, rate limiting, sharding and partitioning, read replicas, event sourcing vs CRUD, schema evolution (the LiveKit JD calls this out explicitly).

Primary resource: _Designing Data-Intensive Applications_ by Martin Kleppmann. Read one chapter per week alongside the technical modules. It's the one book that's genuinely irreplaceable for this career path.

Secondary: Read production engineering blog posts from companies like Cloudflare, Discord, Shopify, and — relevant to the LiveKit role — any WebRTC or media infrastructure postmortems.

---

## Capstone Project — Distributed Usage Metering Service

Inspired by the LiveKit data engineering role. Build a production-grade metering pipeline from scratch.

### What It Does

Accepts session events (start, end, usage ticks) via a gRPC API, processes them concurrently, stores aggregated metrics in a Postgres database, exposes a REST query API for dashboards, and ships with full infrastructure-as-code.

### Components

**gRPC ingestion service (Go)**

- Accepts `SessionEvent` protobufs (session ID, user ID, event type, timestamp, bytes transferred)
    
- Validates and deduplicates events (idempotency key)
    
- Publishes valid events to a Kafka topic
    

**Event processor (Go)**

- Consumes from Kafka with a worker pool
    
- Aggregates per-session usage (start time, end time, total bytes)
    
- Writes to Postgres with `sqlc`-generated queries
    
- Implements dead letter queue for failed events
    
- Instrumented with OpenTelemetry traces + Prometheus metrics
    

**REST query API (Go)**

- `GET /usage?user_id=X&start=Y&end=Z` — returns aggregated usage
    
- `GET /health` — liveness probe
    
- Backed by the same Postgres DB
    

**Infrastructure (Terraform)**

- GCS bucket for Terraform remote state
    
- Cloud SQL Postgres instance
    
- GKE cluster (or local Kind for cost)
    
- All resources versioned and reproducible
    

**Kubernetes manifests**

- Deployments for all three Go services
    
- ConfigMaps for config, Secrets for credentials
    
- Liveness/readiness probes on all pods
    
- Horizontal pod autoscaling on the processor
    

**CI/CD (GitHub Actions)**

- Lint → test with race detector → build Docker images → push to registry → apply Kubernetes manifests
    
- Reusable workflow for lint + test
    

**Observability**

- Prometheus metrics on the processor (events processed, lag, error rate)
    
- Structured logging with `slog` across all services
    
- OTel traces from gRPC ingestion through to DB write
    

**Tests**

- Table-driven unit tests for business logic
    
- Integration test for the gRPC handler using `httptest`-equivalent
    
- Race detector enabled in CI
    

### What This Demonstrates to Hiring Teams

- Go proficiency across the full language (concurrency, error handling, interfaces, generics, testing)
    
- Protobuf + gRPC (the LiveKit stack and most distributed systems)
    
- Kafka event streaming with delivery guarantees
    
- SQL at production depth (schema design, migrations, idempotency)
    
- Redis for distributed coordination (add a distributed lock on the event processor)
    
- Docker multi-stage builds + Kubernetes deployment
    
- Terraform for reproducible infrastructure
    
- Observability (OTel + Prometheus + structured logs)
    
- CI/CD as code
    

---

## LeetCode Readiness Appendix

A focused reference for Go-specific patterns in coding problems — work through these alongside Phases 2-3.

- Slice manipulation tricks (two pointers, in-place modification)
    
- Map patterns (frequency counting, "seen" set, grouping)
    
- String ↔ `[]byte` ↔ `[]rune` conversions (critical for string problems)
    
- Two pointers and sliding window in Go
    
- Sorting with `sort.Slice` and `sort.SliceStable`
    
- Stack and queue with slices
    
- Recursion + memoization with maps
    
- `math.MaxInt`, `math.MinInt` for boundary initialization
    

---

## What's Intentionally Left Out

- CGo
    
- Reflection (beyond reading generated code)
    
- Go plugin system
    
- Assembly or compiler internals
    
- Deep WebRTC internals (domain-specific, not a Go topic)
    
- ClickHouse internals (the LiveKit JD mentions it; learn it on the job once the foundations are solid)