# Project Structure

Go doesn't enforce a strict project layout, but there are widely-adopted conventions.

## TypeScript Comparison

| TypeScript/Node | Go |
|-----------------|-----|
| `src/` | `cmd/`, `internal/`, `pkg/` |
| `dist/` or `build/` | Binary output (no folder needed) |
| `node_modules/` | Global cache (not in project) |
| `index.ts` | `main.go` |
| `lib/` | `pkg/` or `internal/` |
| `__tests__/` or `*.test.ts` | `*_test.go` (same directory) |

## Simple Project (Single Package)

```
myproject/
├── go.mod
├── go.sum
├── main.go
└── main_test.go
```

TypeScript equivalent:
```
myproject/
├── package.json
├── package-lock.json
├── index.ts
└── index.test.ts
```

Good for: CLI tools, small utilities, learning.

## Standard Project Layout

```
myproject/
├── cmd/                    # Application entry points
│   ├── myapp/
│   │   └── main.go
│   └── mycli/
│       └── main.go
├── internal/               # Private packages (cannot be imported by other modules)
│   ├── config/
│   │   └── config.go
│   └── database/
│       └── db.go
├── pkg/                    # Public packages (can be imported by other modules)
│   └── utils/
│       └── helpers.go
├── api/                    # API definitions (OpenAPI, protobuf, etc.)
├── web/                    # Web assets (templates, static files)
├── scripts/                # Build/CI scripts
├── docs/                   # Documentation
├── go.mod
├── go.sum
└── README.md
```

TypeScript equivalent:
```
myproject/
├── src/
│   ├── index.ts            # Entry point
│   ├── config/             # Internal modules
│   └── database/
├── lib/                    # Shared library code
│   └── utils/
├── api/
├── public/
├── scripts/
├── docs/
├── package.json
└── README.md
```

## Key Directories

### `/cmd`

Each subdirectory is a separate executable (like having multiple `"main"` entries in package.json):

```go
// cmd/myapp/main.go
package main

import "github.com/imgarylai/learn-go/internal/config"

func main() {
    cfg := config.Load()
    // ...
}
```

Build with:
```bash
go build ./cmd/myapp
```

TypeScript equivalent:
```json
// package.json with multiple entry points
{
  "bin": {
    "myapp": "./dist/cmd/myapp/index.js",
    "mycli": "./dist/cmd/mycli/index.js"
  }
}
```

### `/internal`

**Special directory**: Go compiler prevents importing `internal` packages from outside the module.

```
myproject/
└── internal/
    └── secret/
        └── secret.go    # Only importable within myproject
```

This is **enforced by the Go compiler**, not just convention.

TypeScript has no equivalent - you'd use:
- Private npm packages
- Not exporting from `index.ts`
- `@internal` JSDoc comments (not enforced)

### `/pkg`

Public library code that other projects can import:

```go
import "github.com/imgarylai/learn-go/pkg/utils"
```

TypeScript equivalent:
```typescript
// Exporting from package entry point
export { utils } from './lib/utils';
```

Note: Some projects skip `/pkg` and put public packages at the root.

## Packages vs Modules

### TypeScript: File = Module

```typescript
// Each file is a module
// utils/strings.ts
export function capitalize(s: string): string { ... }

// utils/numbers.ts
export function add(a: number, b: number): number { ... }

// Import individually
import { capitalize } from './utils/strings';
import { add } from './utils/numbers';
```

### Go: Directory = Package

```go
// All files in a directory share the same package
// utils/strings.go
package utils
func Capitalize(s string) string { ... }

// utils/numbers.go
package utils
func Add(a, b int) int { ... }

// Import the whole package
import "myproject/utils"
utils.Capitalize("hello")
utils.Add(1, 2)
```

All `.go` files in a directory must have the same package name.

### Package Naming

```go
// Good
package http
package user
package config

// Bad
package httpPackage    // Don't repeat "package"
package user_service   // No underscores
package myUtils        // No camelCase
```

Compare to TypeScript:
```typescript
// TypeScript files can be named anything
userService.ts
my-utils.ts
HTTPClient.ts
```

### Import Paths

```go
import (
    // Standard library (like Node built-ins)
    "fmt"
    "net/http"

    // External dependencies (like npm packages)
    "github.com/fatih/color"

    // Internal packages (like relative imports)
    "github.com/imgarylai/learn-go/internal/config"
)
```

TypeScript equivalent:
```typescript
import fs from 'fs';                    // Built-in
import chalk from 'chalk';              // npm package
import { config } from './config';      // Relative
```

Key difference: Go uses **absolute paths** for internal imports, not relative.

## Tests Live With Code

```
# TypeScript (common patterns)
src/
├── utils.ts
└── __tests__/
    └── utils.test.ts

# Or
src/
├── utils.ts
└── utils.test.ts

# Go (always same directory)
utils/
├── utils.go
└── utils_test.go
```

Go convention: tests are always `*_test.go` in the same directory.

## Try It

Create a multi-package structure:

```bash
mkdir -p cmd/hello internal/greeter

# Create files (see examples in this project)
```

## Next

See [04-testing.md](04-testing.md) for writing tests.
