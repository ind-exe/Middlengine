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
