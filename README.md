# Middleware Engine

A lightweight and flexible middleware engine for Go web servers.  
This package allows you to easily chain multiple HTTP middleware functions around a base HTTP handler.

## Features

- **Simple Middleware Chaining:**  
  Add middleware functions that wrap around your base HTTP handler.
- **Ordered Execution:**  
  Middlewares are executed in the order they are added (the first added middleware is the outermost).

- **Easy Integration:**  
  Implements the `http.Handler` interface for seamless integration with Go's standard library.

## Installation

Install the package using `go get`:

```bash
go get github.com/ind-exe/middlengine
```

# Middlengine Usage Examples

## Middleware Structure

A Middlengine middleware should follow this pattern:

```go
// Timer middleware example
func timer(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)  // Pass request to next handler
        fmt.Printf("Request took: %v\n", time.Since(start))
    })
}
```

Key characteristics:

1. Accepts `http.Handler` as input parameter
2. Returns a new `http.Handler`
3. Calls `next.ServeHTTP()` to pass control down the chain
4. Can execute logic before and/or after the handler

---

## Complete Usage Example

```go
package main

import (
    "fmt"
    "net/http"
    "time"
    "github.com/yourusername/middlengine"
)

func main() {
    // Create router
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello with timing!"))
    })

    // Initialize middleware engine
    engine := middlengine.NewEngine(mux)

    // Add middleware to measure request duration
    engine.Use(timer)

    // Build the handler chain
    engine.CreateChain()

    // Start server with middleware chain
    fmt.Println("Server listening on :8080")
    http.ListenAndServe(":8080", engine)
}

// Timer middleware implementation
func timer(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        defer func() {
            fmt.Printf("%s %s - %v\n",
                r.Method,
                r.URL.Path,
                time.Since(start),
            )
        }()
        next.ServeHTTP(w, r)
    }
}
```

---

## Execution Flow

When making a request to `/`:

1. `timer` middleware starts clock
2. Request enters your handler
3. Response is sent to client
4. `timer` logs the duration after response completes

Output:

```
GET / - 127.8µs
```

---

## Key Points

- **Middleware Order**: First added = outermost layer
- **Chain Construction**: Must call `CreateChain()` after adding middleware
- **Handler Types**: Works with any `http.Handler` implementation
- **Dependencies**: Remember to import required packages (`time`, `fmt`)

---
