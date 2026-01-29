# Learn Go

Learning Go with AI guidance, coming from a JavaScript/TypeScript background.

## Getting Started

```bash
go run .
```

## Learning Path

| Topic | Description |
|-------|-------------|
| [00 - From JS/TS](docs/00-from-js-ts.md) | Ecosystem comparison and key differences |
| [01 - Go Modules](docs/01-go-modules.md) | Package management basics |
| [02 - Dependencies](docs/02-dependencies.md) | Adding and managing deps |
| [03 - Project Structure](docs/03-project-structure.md) | Layout conventions |
| [04 - Testing](docs/04-testing.md) | Built-in testing |
| [05 - Tooling](docs/05-tooling.md) | Go toolchain |
| [06 - Structs](docs/06-structs.md) | Types and structs |
| [07 - Pointers](docs/07-pointers.md) | Understanding pointers |
| [08 - Web Frameworks](docs/08-web-frameworks.md) | Chi, Gin, Echo, Fiber |
| [09 - File Processing](docs/09-file-processing.md) | Reading, writing, CSV, JSON |
| [10 - Data Processing](docs/10-data-processing.md) | Slices, generics, gota DataFrame |

## Exercises

Practice exercises with JS/TS comparisons. See [exercises/README.md](exercises/README.md).

```bash
cd exercises/01-basics && go test -v
```

| # | Topic | Focus |
|---|-------|-------|
| 01 | Basics | Variables, types, constants |
| 02 | Functions | Multiple returns, errors, defer |
| 03 | Structs | Methods, embedding, tags |
| 04 | Collections | Slices, maps, iteration |
| 05 | Interfaces | Implicit interfaces, assertions |
| 06 | Concurrency | Goroutines, channels, select |
| 07 | File Processing | CSV, JSON, line-by-line |
| 08 | Data Processing | Filter, map, reduce, gota |

## Quick Reference

```bash
go run .          # Run program
go build          # Compile binary
go test ./...     # Run tests
go fmt ./...      # Format code
go get <pkg>      # Add dependency
```
