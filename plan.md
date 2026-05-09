# Go Project Progression Guide

A curated learning path for Go, oriented around systems programming, networking, and infrastructure tooling. Assumes prior programming experience (C, shell, web) — beginner hand-holding is skipped.

---

## Stage 1 — Get Go-Fluent
*Goal: Get comfortable with Go's project structure, error handling, stdlib, and idioms before tackling meatier stuff. A few days each.*

### Task Manager CLI
Build a CLI to add, list, complete, and delete tasks persisted to a JSON file.

**What you'll learn:** structs, JSON marshalling/unmarshalling, file I/O, `os.Args`, the `flag` package, Go error handling patterns (`if err != nil`).

**Stretch:** add due dates, priorities, filtering.

---

### Static Site Generator
Read a directory of Markdown files, apply an HTML template to each, write rendered HTML output files.

**What you'll learn:** `filepath.Walk`, `text/template` and `html/template`, string processing, os/io packages.

**Stretch:** watch mode that rebuilds on file change using `fsnotify`.

---

## Stage 2 — Networking Fundamentals
*Goal: Understand Go's networking primitives and how the protocols you use daily actually work. 1–3 weeks each.*

### HTTP Server from Scratch
Implement a basic HTTP/1.1 server using raw TCP — no `net/http`. Parse the request line, headers, and body manually. Send a valid HTTP response.

**What you'll learn:** `net.Listen`, `net.Conn`, reading/writing raw bytes, HTTP wire format, why the standard library exists.

**Stretch:** serve static files, handle multiple concurrent connections with goroutines.

---

### DNS Resolver
Send raw DNS queries over UDP, parse the binary response format (wire format) yourself — no resolver libraries.

**What you'll learn:** `net.UDPConn`, binary protocol parsing, bit manipulation, Go's `encoding/binary`, how DNS actually works (questions, answers, RR types).

**Stretch:** support recursive resolution, implement a local cache with TTL expiry.

**Reference:** RFC 1035 is the spec. Reading it directly is worth it.

---

### Port Scanner
Scan a range of ports on a host concurrently, report which are open.

**What you'll learn:** goroutines, `sync.WaitGroup`, worker pool pattern, `net.DialTimeout`, context cancellation.

**Stretch:** service fingerprinting, CIDR range input, rate limiting.

---

## Stage 3 — Systems Level
*Goal: Get into process management, Linux internals, and low-level Go. 2–4 weeks each.*

### Shell
Implement a Unix shell: parse input, fork and exec commands, handle pipes (`|`), I/O redirection (`>`, `<`), and built-ins like `cd` and `exit`.

**What you'll learn:** `os/exec`, `syscall`, process groups, signal handling, `io.Pipe`, stdin/stdout/stderr wiring.

**Reference:** [Write a Shell in C](https://brennan.io/2015/01/16/write-a-shell-in-c/) — translate the concepts to Go.

**Stretch:** job control (`fg`, `bg`, `&`), environment variable expansion, command history.

---

### Key-Value Store
In-memory KV store with a simple TCP protocol (`GET key`, `SET key value`, `DEL key`) and persistence via a write-ahead log (WAL).

**What you'll learn:** TCP server design, custom protocol parsing, concurrent map access (`sync.RWMutex`), append-only log format, crash recovery on startup.

**Stretch:** TTL/expiry, snapshotting, basic replication to a replica node.

---

### Containers from Scratch
Use Linux namespaces and cgroups via syscalls in Go to isolate a process — give it its own PID namespace, UTS namespace (hostname), and filesystem via `chroot`. This is what Docker actually does.

**What you'll learn:** `syscall.CLONE_NEWPID`, `syscall.CLONE_NEWNS`, `syscall.CLONE_NEWUTS`, cgroups via `/sys/fs/cgroup`, `chroot`, how container runtimes work under the hood.

**Reference:** Liz Rice's "Containers from Scratch" talk (GopherCon) — uses Go specifically.

**Stretch:** network namespace isolation, overlay filesystem layering.

---

## Stage 4 — Distributed Systems & Infra Tooling
*Goal: Build the kinds of tools that underpin real infrastructure. 3–6 weeks each.*

### Persistent Message Queue
A pub/sub message queue with topics, multiple consumers, and durability — messages survive restarts via an append-only log on disk.

**What you'll learn:** producer/consumer patterns, WAL design, offset tracking, at-least-once delivery guarantees, why Kafka makes the tradeoffs it does.

**Stretch:** consumer groups, partition-like sharding, a simple HTTP API for producing/consuming.

---

### Metrics Collector & Prometheus Exporter
Scrape system metrics from `/proc` (CPU, memory, disk, network), aggregate them, and expose them in Prometheus exposition format over HTTP.

**What you'll learn:** `/proc` filesystem internals, Prometheus text format, `net/http` properly, goroutine-safe metric aggregation.

**Stretch:** pull metrics from other processes via their HTTP endpoints, alerting rules.

---

### Reverse Proxy / Load Balancer
Route incoming HTTP requests across backend servers with round-robin load balancing, active health checks, circuit breaking, and retry logic.

**What you'll learn:** `httputil.ReverseProxy`, connection pooling, health check loops, circuit breaker pattern, middleware chaining.

**Stretch:** weighted routing, path-based routing, rate limiting per client.

---

### Distributed Key-Value Store with Raft
Extend the basic KV store to run across multiple nodes with consensus — no single point of failure. Implement the Raft algorithm for leader election and log replication.

**What you'll learn:** distributed consensus, leader election, log replication, network partitions, why this is hard.

**Reference:** [MIT 6.5840 (formerly 6.824)](https://pdos.csail.mit.edu/6.824/) — the lab assignments are in Go and are the gold standard for this. Work through them.

---

## Stage 5 — Stretch Goals
*Long-running, open-ended projects. Pick one and go deep.*

### Interpreter (Crafting Interpreters in Go)
Work through Robert Nystrom's [Crafting Interpreters](https://craftinginterpreters.com/) (free online) but implement it in Go instead of Java/C. Build a complete scripting language from scratch.

**What you'll learn:** lexing, parsing, AST design, tree-walk interpretation, bytecode compilation, garbage collection concepts. Touches everything.

---

### Git Implementation in Go
Implement core Git commands properly in Go: `init`, `add`, `commit`, `log`, `diff`, `branch`, `checkout`. Not shell scripts this time — actually handle the object store, pack files, and binary formats.

**What you'll learn:** content-addressable storage, SHA-1 object hashing, tree/blob/commit objects, binary file formats, compression (zlib).

**Reference:** [Git Internals](https://git-scm.com/book/en/v2/Git-Internals-Plumbing-and-Porcelain) chapter of the Pro Git book.

---

### TLS from Scratch
Implement the TLS 1.3 handshake manually against a real server without using `crypto/tls`. Handle key exchange, certificate verification, and symmetric encryption yourself.

**What you'll learn:** public key cryptography, X.509 certificates, ECDH key exchange, AEAD ciphers, why TLS is designed the way it is.

**Reference:** RFC 8446 (TLS 1.3 spec) and [The Illustrated TLS Connection](https://tls13.xargs.org/).

---

## Resources

| Resource | What it's for |
|---|---|
| [Tour of Go](https://go.dev/tour/) | Quick syntax orientation if needed |
| [Effective Go](https://go.dev/doc/effective_go) | Idiomatic Go patterns — read early |
| [Go by Example](https://gobyexample.com/) | Quick reference for stdlib patterns |
| [Beej's Guide to Network Programming](https://beej.us/guide/bgnet/) | Networking concepts (C-based but concepts are universal) |
| [Crafting Interpreters](https://craftinginterpreters.com/) | Interpreter project (free online) |
| [MIT 6.5840](https://pdos.csail.mit.edu/6.824/) | Distributed systems labs in Go |
| [Liz Rice — Containers from Scratch](https://www.youtube.com/watch?v=8fi7uSYlOdc) | Container project starting point |
| RFC 1035 | DNS wire format spec |
| RFC 8446 | TLS 1.3 spec |

---

## Recommended Order

If you're unsure where to start given your background:

1. **DNS Resolver** — concrete, contained, teaches binary protocol parsing
2. **HTTP Server from Scratch** — you know HTTP, now see the wire
3. **Shell** — process management, syscalls, signals
4. **Containers from Scratch** — the Linux internals payoff
5. **Distributed KV Store / Raft** — the deep end

The Raft implementation or the Interpreter are the two projects that will teach you the most per hour invested. Both are serious undertakings — plan weeks, not days.

---

## Books

### Getting Fluent with Go

**Learning Go** — Jon Bodner (2nd ed, 2024)
The best "learn to think in Go" book rather than just learn the syntax. Written for developers with experience in other languages — teaches Go's idioms so you don't end up recreating patterns that don't fit the language. Covers generics properly too.

**100 Go Mistakes and How to Avoid Them** — Teiva Harsanyi
Read this after you've written a bit of Go and some things feel off but you can't articulate why. Covers mistakes across concurrency, error handling, testing, and memory management — the kind of things that trip up even experienced developers.

**Learn Go with Tests** — Chris James *(free online)*
One of the most popular free Go resources. A test-driven approach to learning the language — Go takes testing seriously as a first-class concern and this teaches both simultaneously.

---

### Systems & Networking

**Network Programming with Go** — Jan Newmarch
Directly relevant to the DNS resolver and HTTP server projects. Covers writing secure, readable, production-ready network code in Go.

**Concurrency in Go** — Katherine Cox-Buday
Go's concurrency model is its biggest differentiator from C. This goes deep on it — best practices and patterns for goroutines, channels, and the scheduler. Essential before tackling the distributed systems projects.

**Build an Orchestrator in Go (From Scratch)**
Guides you through building your own orchestration framework using Go and the Docker API, uncovering the inner mechanics of systems like Kubernetes and Nomad. Sits perfectly alongside the containers and distributed systems projects.

---

### Microservices & Web

**Let's Go / Let's Go Further** — Alex Edwards
The most recommended Go web development books. Let's Go builds a web app from scratch; Let's Go Further covers APIs, authentication, rate limiting, and deployment. Solid and practical — the right place to start when you get to web/API territory.

**gRPC Microservices in Go** — Hüseyin Babal
The most targeted book for microservices specifically. Covers designing resilient microservice architecture, inter-service communication with gRPC, backward compatible API design, and hexagonal architecture applied to microservices. gRPC is what Go microservices actually use in production rather than REST between services.

**Microservices with Go** — Alexander Shuiskov (2nd ed)
More comprehensive than the gRPC book — covers the full picture from service design through to production operations. Includes event-driven architecture, Kafka messaging patterns, Kubernetes deployment, CI/CD, and cloud-native observability. Author has background at Uber and Booking.com which shows in the production-grade focus.

**Event-Driven Architecture in Golang**
Covers building systems with asynchronicity and eventual consistency in Go. Directly relevant to the Kafka/event queue side of systems design — pairs well with the Microservices with Go book.

> **Note:** A lot of microservices design knowledge isn't Go-specific. **Building Microservices** (Sam Newman) and **Designing Distributed Systems** (Brendan Burns, free from Microsoft) cover the architectural patterns in a language-agnostic way and are excellent complements to the Go-specific books above.

---

### Compiler / Interpreter (Stage 5 Companion)

**Writing An Interpreter In Go / Writing A Compiler In Go** — Thorsten Ball
The Go-specific companion to the Crafting Interpreters project. You build a complete language called Monkey from scratch — the interpreter book covers lexing, parsing, and tree-walk evaluation; the compiler book follows up by replacing the interpreter with a bytecode compiler and virtual machine. All code in Go, no third-party libraries.

Well regarded for being unusually clear on a topic that usually gets very academic. Ball doesn't assume compiler theory knowledge — you build the thing and the theory becomes obvious from doing it. Shorter and more focused than Crafting Interpreters but Go-native throughout.

---

### Suggested Reading Order

| When | Read |
|---|---|
| Starting out | Learning Go |
| After first real project | 100 Go Mistakes |
| Hitting networking projects | Network Programming with Go |
| Before distributed systems | Concurrency in Go |
| Getting into web/APIs | Let's Go → Let's Go Further |
| Moving into microservices | gRPC Microservices in Go → Microservices with Go |
| Stage 5 interpreter project | Writing An Interpreter In Go → Writing A Compiler In Go |
