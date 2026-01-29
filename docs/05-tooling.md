# Go Tooling

Go comes with powerful built-in tools. Here's what you need to know.

## TypeScript Comparison

| Task | TypeScript/Node | Go |
|------|-----------------|-----|
| Format | Prettier | `go fmt` (built-in) |
| Lint | ESLint | `go vet` + golangci-lint |
| Compile | `tsc` | `go build` (built-in) |
| Run | `ts-node`, `tsx` | `go run` |
| Test | Jest, Vitest | `go test` (built-in) |
| Package manager | npm, yarn, pnpm | `go mod` (built-in) |
| Language server | tsserver | gopls |
| Debugger | Node inspector | dlv (Delve) |

**Key difference**: Most Go tools are built-in. No `devDependencies` needed.

## Built-in Tools

### `go fmt` - Code Formatting

Go has ONE official style. No configuration, no debates.

```bash
# TypeScript
npx prettier --write .
# Or configure .prettierrc, debate tabs vs spaces...

# Go - just run it
go fmt ./...
```

No config files. No options. Everyone's Go code looks the same.

```bash
# Format a file
go fmt main.go

# Format all files
go fmt ./...

# Alternative: gofmt with more options
gofmt -w -s .  # -w writes, -s simplifies
```

**Run before every commit.** Most editors do this automatically on save.

### `go vet` - Static Analysis

Catches common mistakes the compiler doesn't:

```bash
# TypeScript
npx eslint .

# Go
go vet ./...
```

Detects:
- Printf format errors
- Unreachable code
- Invalid struct tags
- And more...

### `go build` - Compilation

```bash
# TypeScript
tsc                              # Compile to JS
npm run build                    # Usually runs tsc + bundler

# Go
go build                         # Produces binary directly
```

```bash
# Build current package
go build

# Build and name the output
go build -o myapp

# Build for different OS/arch (cross-compile!)
GOOS=linux GOARCH=amd64 go build    # Build Linux binary on Mac
GOOS=windows GOARCH=amd64 go build  # Build Windows .exe on Mac

# Build with optimizations stripped (smaller binary)
go build -ldflags="-s -w"
```

Cross-compilation is built-in. No extra tools needed.

### `go run` - Build and Run

```bash
# TypeScript
npx ts-node app.ts
npx tsx app.ts

# Go
go run .
go run main.go
```

```bash
# Run a file
go run main.go

# Run a package
go run .

# Run with arguments
go run . --config=dev
```

### `go install` - Install Binaries

```bash
# TypeScript (global install)
npm install -g typescript

# Go
go install golang.org/x/tools/gopls@latest
```

Installs to `$GOPATH/bin` (usually `~/go/bin`). Add this to your PATH.

```bash
# Install a CLI tool
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Now you can run it
golangci-lint run
```

### `go doc` - Documentation

```bash
# View package docs in terminal
go doc fmt
go doc fmt.Println
go doc -all fmt
```

TypeScript equivalent: Reading types in VS Code or checking npm docs.

### `go generate` - Code Generation

Runs commands from special comments in code:

```go
//go:generate stringer -type=Status
//go:generate mockgen -source=interface.go
```

```bash
go generate ./...
```

TypeScript equivalent: npm scripts that run before build.

## Essential Third-Party Tools

### golangci-lint

The meta-linter. Runs many linters in parallel.

```bash
# TypeScript
npx eslint . --fix

# Go
golangci-lint run --fix
```

Install and run:

```bash
# Install
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run
golangci-lint run

# Run with auto-fix (where possible)
golangci-lint run --fix
```

Create `.golangci.yml` to configure (like `.eslintrc`):

```yaml
linters:
  enable:
    - gofmt
    - govet
    - errcheck
    - staticcheck
    - unused
```

### gopls - Language Server

The official Go language server. Powers IDE features.

```bash
go install golang.org/x/tools/gopls@latest
```

TypeScript equivalent: `tsserver` (built into TypeScript).

Your editor (VS Code, Neovim, etc.) uses this for:

- Autocomplete
- Go to definition
- Find references
- Rename refactoring
- Hover documentation

### dlv - Debugger

```bash
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug
dlv debug

# Debug a test
dlv test
```

TypeScript equivalent: Node.js inspector / Chrome DevTools.

## Environment Variables

```bash
# Show all Go environment
go env

# Important ones:
go env GOPATH      # Where go install puts binaries
go env GOROOT      # Go installation directory
go env GOPROXY     # Module proxy URL
go env GOMODCACHE  # Module cache location
```

TypeScript equivalent: `node -e "console.log(process.env)"` or checking `.nvmrc`.

### Setting Environment

```bash
# Temporary (for one command)
GOOS=windows go build

# Permanent
go env -w GOPROXY=https://proxy.golang.org,direct
```

## Build Tags

Conditional compilation (like `#ifdef` in C, or platform-specific code):

```typescript
// TypeScript - runtime check
if (process.platform === 'linux') {
    // Linux-specific code
}
```

```go
// Go - compile-time separation
// file: server_linux.go
//go:build linux

package main

// This file only compiles on Linux
```

```bash
# Build with custom tags
go build -tags=integration
```

## Workspace Mode (Multi-Module)

For developing multiple modules together (like npm workspaces / yarn workspaces):

```bash
# TypeScript
# Configure workspaces in package.json

# Go
go work init ./module1 ./module2
go work use ./module3
```

Creates `go.work` file (don't commit to version control - it's for local development).

## Quick Reference

| Task | Go Command | TypeScript Equivalent |
|------|------------|----------------------|
| Format code | `go fmt ./...` | `prettier --write .` |
| Find bugs | `go vet ./...` | `eslint .` |
| Run linters | `golangci-lint run` | `eslint .` |
| Build binary | `go build -o app` | `tsc && pkg .` |
| Run program | `go run .` | `tsx app.ts` |
| Run tests | `go test ./...` | `npm test` |
| Test coverage | `go test -cover ./...` | `npm test -- --coverage` |
| View docs | `go doc fmt` | Check npm/types |
| Install tool | `go install pkg@latest` | `npm i -g pkg` |
| Update deps | `go get -u ./...` | `npm update` |
| Clean cache | `go clean -cache` | `rm -rf node_modules` |

## What You Don't Need

Coming from TypeScript, you might look for these. You don't need them in Go:

| TypeScript | Go |
|------------|-----|
| Babel | Not needed (Go compiles directly) |
| Webpack/Vite/esbuild | Not needed (single binary output) |
| ts-node/tsx | `go run` is built-in |
| Jest/Vitest | `go test` is built-in |
| Prettier config | `go fmt` has no config |
| ESLint config | `go vet` + golangci-lint |
| package.json scripts | Just run commands directly |
| node_modules | Global cache, not per-project |

## Try It

```bash
# Format your code
go fmt ./...

# Run static analysis
go vet ./...

# Build a binary
go build -o learn-go

# Run it
./learn-go
```
