# Go Exercises

Practice exercises for learning Go, designed for developers coming from JavaScript/TypeScript.

## How to Use

Each exercise folder contains:
- `*.go` - Exercise file with TODOs to complete
- `*_test.go` - Unit tests to verify your solutions
- `solution.go.txt` - Reference solutions (try before peeking!)

## Running Exercises

```bash
# Run all tests in an exercise
cd exercises/01-basics
go test -v

# Run a specific test
go test -v -run TestGetGreeting

# Run tests with race detection (useful for 06-concurrency)
go test -race -v
```

## Exercise Progression

| # | Topic | Key Concepts |
|---|-------|--------------|
| 01 | Basics | Variables, types, constants, zero values |
| 02 | Functions | Multiple returns, errors, defer, closures |
| 03 | Structs | Types, methods, embedding, tags |
| 04 | Collections | Slices, maps, iteration patterns |
| 05 | Interfaces | Implicit interfaces, type assertions |
| 06 | Concurrency | Goroutines, channels, WaitGroup, select |
| 07 | File Processing | CSV, JSON, bufio, os |
| 08 | Data Processing | Filter, map, reduce, gota |

## Installing Dependencies (Exercise 08)

```bash
cd exercises/08-data-processing
go get github.com/go-gota/gota/dataframe
go get github.com/go-gota/gota/series
```

## Tips

- Read the comments - they compare Go to JS/TS equivalents
- Don't look at solutions until you've tried
- Run `go fmt` to format your code
- Use `go vet` to catch common mistakes

## JS/TS Mental Model Shifts

| JS/TS | Go |
|-------|-----|
| `try/catch` | Return `error` as value |
| `async/await` | Goroutines + channels |
| `Promise.all()` | `sync.WaitGroup` |
| `class` | `struct` + methods |
| `implements` | Implicit (duck typing) |
| `array.map()` | Write a loop |
| `null/undefined` | Zero values |
