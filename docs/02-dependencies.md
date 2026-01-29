# Dependency Management

## TypeScript Comparison

| Task | TypeScript/Node | Go |
|------|-----------------|-----|
| Add package | `npm install pkg` | `go get pkg` |
| Add specific version | `npm install pkg@1.2.3` | `go get pkg@v1.2.3` |
| Update package | `npm update pkg` | `go get -u pkg` |
| Update all | `npm update` | `go get -u ./...` |
| Remove unused | `npm prune` | `go mod tidy` |
| Lock file | `package-lock.json` | `go.sum` |

## Adding Dependencies

### Using `go get`

```bash
# Add latest version
go get github.com/fatih/color           # Like: npm install color

# Add specific version
go get github.com/fatih/color@v1.15.0   # Like: npm install color@1.15.0

# Add specific commit
go get github.com/fatih/color@abc1234   # Like: npm install user/repo#abc1234

# Add latest from a branch
go get github.com/fatih/color@main      # Like: npm install user/repo#main

# Upgrade to latest
go get -u github.com/fatih/color        # Like: npm update color

# Upgrade all dependencies
go get -u ./...                          # Like: npm update
```

### Automatic Detection

Unlike npm, Go can auto-detect dependencies from your imports:

```typescript
// TypeScript: Must install first
// npm install chalk
import chalk from 'chalk';
```

```go
// Go: Just import it
package main

import "github.com/fatih/color"  // Just import it

func main() {
    color.Cyan("Hello!")
}
```

Then run:
```bash
go mod tidy  # Automatically adds the dependency
```

## The go.sum File

`go.sum` is the lockfile containing cryptographic checksums of dependencies.

```
github.com/fatih/color v1.15.0 h1:abc123...
github.com/fatih/color v1.15.0/go.mod h1:def456...
```

### Comparison with package-lock.json

| Aspect | package-lock.json | go.sum |
|--------|-------------------|--------|
| Contains | Exact versions + resolved URLs | Checksums only |
| Versions defined in | Lock file | go.mod |
| Security | Integrity hashes | Cryptographic checksums |
| Edit manually? | No | No |
| Commit to git? | Yes | Yes |

## Semantic Versioning

Both ecosystems use semver:

```
v1.2.3
│ │ │
│ │ └── Patch: bug fixes
│ └──── Minor: new features (backward compatible)
└────── Major: breaking changes
```

### Major Version Suffix (v2+)

This is different from npm! For major versions 2+, the import path changes:

```typescript
// TypeScript: same import, different version in package.json
import something from 'some-package';  // Could be v1, v2, v3...
```

```go
// Go: version is part of the import path
import "github.com/user/repo"     // v0 or v1
import "github.com/user/repo/v2"  // v2
import "github.com/user/repo/v3"  // v3
```

This allows importing multiple major versions simultaneously.

## Version Selection

Go uses **Minimal Version Selection (MVS)**:

```
Your app requires:  pkg v1.2.0
Dependency A requires: pkg v1.3.0
Dependency B requires: pkg v1.1.0

npm: Uses latest (v1.3.0) or complex resolution
Go:  Uses minimum that satisfies all (v1.3.0)
```

- Always uses the minimum version that satisfies all requirements
- Predictable and reproducible
- No "phantom updates" - explicit upgrades only

## Useful Commands

```bash
# List all dependencies
go list -m all                              # Like: npm ls

# List available versions of a module
go list -m -versions github.com/fatih/color # Like: npm view color versions

# Check for available updates
go list -m -u all                           # Like: npm outdated

# Why is this module in my build?
go mod why github.com/some/module           # Like: npm explain

# Download all dependencies
go mod download                             # Like: npm ci (but to global cache)
```

## Private Packages

```bash
# TypeScript: .npmrc
@mycompany:registry=https://npm.mycompany.com

# Go: GOPRIVATE environment variable
go env -w GOPRIVATE=github.com/mycompany/*
```

## Try It

```bash
# Add a dependency
go get github.com/fatih/color

# Check your files
cat go.mod
cat go.sum

# See the dependency graph
go mod graph
```

## Next

See [03-project-structure.md](03-project-structure.md) for organizing your code.
