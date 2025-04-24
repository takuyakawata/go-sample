# Go Development Best Practices Guide

This document summarizes the key best practices for developing applications in Go (Golang). Feel free to share it with Junie.

---

## 1. Project Structure

```
myapp/
├── cmd/                        // Application entry points
│   └── myapp/
│       └── main.go
├── internal/                   // App-specific packages (not importable by other modules)
│   └── product/                // “商品”ドメイン固有パッケージ
│       ├── domain/             // Domain models and business logic
│       ├── usecase/            // Use case / service layer
│       ├── interface/          // Adapters for HTTP, gRPC, etc.
│       └── infrastructure/     // Implementations for DB, external APIs, etc.
├── test/                       // テスト用パッケージ
│   └── product/                // product ドメインのテスト
│       ├── domain/
│       ├── usecase/
│       ├── interface/
│       └── infrastructure/
├── pkg/                        // Generic libraries intended for reuse across projects
├── api/                        // OpenAPI definitions, JSON schemas, etc.
├── configs/                    // Configuration files
├── scripts/                    // Build and deployment scripts
├── Dockerfile                  // Container definition
├── go.mod                      // Go modules dependencies
└── go.sum                      // Checksum of dependencies

```

- **`cmd/`**: Contains small `main` packages for launching each app.
- **`internal/`**: Hosts code private to this application.
- **`pkg/`**: Houses reusable packages that may be imported by other modules.

---

## 2. Coding Style

- **`gofmt` / `go fmt`**: Use the standard formatter; enforce via CI or pre-commit hooks.
- **`golint` / `staticcheck`**: Run linters to maintain code quality.
- **Naming Conventions**: Package names in lowercase; exported types and functions in PascalCase; variables in camelCase.
- **Comments**: Document public APIs with full sentences, e.g. `// Package foo ...` or `// Foo does ...`.

---

## 3. Dependency Management

- **Go Modules**: Manage dependencies in `go.mod` and `go.sum` with precise versioning.
- Minimize use of `replace` directives.
- **Semantic Versioning**: Follow SemVer (e.g. v1.2.3) for all public releases.

---

## 4. Error Handling

```go
if err := doSomething(); err != nil {
    return fmt.Errorf("doSomething failed: %w", err)
}
```

- **Wrapping errors**: Use `%w` to preserve error chains.
- **Custom error types**: Define custom types and allow assertion via `errors.Is` / `errors.As`.
- Reserve **`panic`** for unrecoverable conditions; handle panics with a top-level `recover`.

---

## 5. Testing

- **Table-driven tests**: Consolidate test cases into tables for clarity.
- **`go test`**: Run with `-coverprofile` and `-race` flags in CI pipelines.
- **Mocking**: Use `gomock` or `testify/mock` to isolate external dependencies.

---

## 6. Concurrency

```go
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // perform work
}()
wg.Wait()
```

- **`context.Context`**: Pass through goroutines for cancellation and timeouts.
- **`sync.WaitGroup`**: Wait for multiple goroutines to finish.
- **Channels**: Use buffered or unbuffered channels as required.
- **Error groups**: Leverage `errgroup.Group` to manage parallel tasks and collect errors.

---

## 7. Logging

- **Structured logging**: Use libraries like `zap` or `logrus`.
- **Log levels**: Appropriately use DEBUG, INFO, WARN, ERROR.
- **Context fields**: Include request IDs, user IDs, and other context in log entries.

---

## 8. Performance & Metrics

- **Profiling**: Integrate `pprof` for CPU and memory profiling.
- **Metrics**: Expose application metrics via a Prometheus client.
- **Benchmarking**: Use `go test -bench` to identify and optimize hot paths.

---

## 9. Security

- **Dependency scanning**: Run vulnerability checks after `go mod tidy`.
- **Input validation**: Validate external input, e.g., with the `validator` package.
- **Secret management**: Store secrets in Vault, AWS Secrets Manager, or similar solutions.

---

## 10. CI/CD

- **Automation**: Use GitHub Actions or GitLab CI to automate:
    - Build, test, and lint steps
    - Container image builds and pushes
    - Deployments (Staging → Production)
- **Branch protection**: Prevent direct merges to `main`; require code reviews.

---

These guidelines provide a solid foundation for Go development. Adapt them to fit your project's needs and share them with your team to ensure consistency and high quality.
