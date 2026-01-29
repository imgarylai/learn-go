# Go Modules

Go modules are the official dependency management system in Go (since Go 1.11, default since Go 1.16).

## TypeScript Comparison

| Concept | TypeScript/Node | Go |
|---------|-----------------|-----|
| Package manifest | `package.json` | `go.mod` |
| Lock file | `package-lock.json` | `go.sum` |
| Initialize | `npm init` | `go mod init` |
| Package registry | npmjs.com | proxy.golang.org |
| node_modules | Local folder | Global cache (`$GOPATH/pkg/mod`) |

## What is a Module?

A module is a collection of Go packages stored in a directory with a `go.mod` file at its root.

```typescript
// TypeScript: package.json defines your project
{
  "name": "my-project",
  "version": "1.0.0",
  "dependencies": { ... }
}
```

```go
// Go: go.mod defines your module
module github.com/imgarylai/learn-go

go 1.25.6

require (
    github.com/fatih/color v1.15.0
)
```

## Initializing a Module

```bash
# TypeScript
npm init

# Go
go mod init <module-path>
```

Example:
```bash
go mod init github.com/imgarylai/learn-go
```

The module path is typically your repository URL. This becomes the import path for your packages (like your npm package name, but usually includes the full repo URL).

## The go.mod File

```go
module github.com/imgarylai/learn-go

go 1.25.6

require (
    github.com/fatih/color v1.15.0
)
```

### Directives

| Directive | Purpose | package.json Equivalent |
|-----------|---------|-------------------------|
| `module` | Declares the module path | `"name"` |
| `go` | Specifies the Go version | `"engines": { "node": ">=18" }` |
| `require` | Lists dependencies and versions | `"dependencies"` |
| `replace` | Substitutes a module with another | `npm link` or yarn resolutions |
| `exclude` | Prevents a specific version from being used | - |
| `retract` | Marks versions of your own module as not recommended | `npm deprecate` |

## Common Commands

```bash
# Initialize a new module
go mod init github.com/username/project    # Like: npm init

# Add missing and remove unused modules
go mod tidy                                 # Like: npm prune + auto-install

# Download dependencies to local cache
go mod download                             # Like: npm install (but global cache)

# Verify dependencies have expected content
go mod verify                               # Like: npm audit signatures

# Explain why a module is needed
go mod why github.com/some/module           # Like: npm explain

# Show module dependency graph
go mod graph                                # Like: npm ls --all
```

## Key Difference: No node_modules

```bash
# TypeScript: deps in project folder
my-project/
├── node_modules/     # 500MB+ of dependencies
├── package.json
└── src/

# Go: deps in global cache
my-project/
├── go.mod            # Just 2 files!
├── go.sum
└── main.go

# Dependencies stored globally at:
~/go/pkg/mod/
```

This means:
- No duplicate dependencies across projects
- Faster project setup
- Smaller project folders

## Try It

```bash
# In your project directory
go mod tidy
cat go.mod
```

## Next

See [02-dependencies.md](02-dependencies.md) for managing dependencies.
