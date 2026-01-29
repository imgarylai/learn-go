# Go for JavaScript/TypeScript Developers

A side-by-side comparison to help you transition from the JS/TS ecosystem.

## Ecosystem Comparison

| Concept | JavaScript/TypeScript | Go |
|---------|----------------------|-----|
| Package manager | npm / yarn / pnpm | `go mod` (built-in) |
| Package registry | npmjs.com | proxy.golang.org (mirrors GitHub, etc.) |
| Manifest file | `package.json` | `go.mod` |
| Lock file | `package-lock.json` / `yarn.lock` | `go.sum` |
| Install deps | `npm install` | `go mod download` (automatic) |
| Add dependency | `npm install pkg` | `go get pkg` |
| Run script | `npm run dev` | `go run .` |
| Build | `npm run build` | `go build` |
| Test | `npm test` (jest, vitest) | `go test ./...` (built-in) |
| Linter | eslint | `golangci-lint` / `go vet` |
| Formatter | prettier | `go fmt` (built-in) |
| Types | TypeScript compiler | Go compiler (built-in) |
| Version manager | nvm | gvm, goenv, or asdf |

## Key Differences

### No `node_modules`

Go caches dependencies globally in `$GOPATH/pkg/mod`. No massive `node_modules` folders.

```bash
# JS: installs to ./node_modules
npm install

# Go: downloads to global cache
go mod download
# Your project stays clean!
```

### No Build Step for Types

In TypeScript, you compile `.ts` → `.js`. In Go, types are native.

```typescript
// TypeScript: needs compilation
const add = (a: number, b: number): number => a + b;
```

```go
// Go: compiles directly
func add(a int, b int) int {
    return a + b
}
```

### No `package.json` Scripts

Go doesn't have npm scripts. You run commands directly:

```json
// package.json
{
  "scripts": {
    "dev": "nodemon server.js",
    "build": "tsc",
    "test": "jest"
  }
}
```

```bash
# Go equivalents - just run them directly
go run .
go build
go test ./...

# Or use Makefile for complex scripts
```

### Imports Work Differently

```typescript
// JS/TS: relative paths, file extensions, or bare specifiers
import { helper } from './utils/helper';
import express from 'express';
```

```go
// Go: always full module paths
import (
    "github.com/imgarylai/learn-go/internal/helper"
    "github.com/gin-gonic/gin"
)
```

### No Default Exports

Go doesn't have default exports. Everything is named.

```typescript
// JS/TS
export default function handler() {}
export const PORT = 3000;
```

```go
// Go: capitalize to export
func Handler() {}  // exported (public)
func helper() {}   // unexported (private)
const Port = 3000  // exported
const port = 3000  // unexported
```

## Common Translations

### Variables

```typescript
// TypeScript
const name: string = "Go";
let count: number = 0;
const config = { port: 3000 };  // type inferred
```

```go
// Go
var name string = "Go"
name := "Go"           // shorthand (type inferred)
var count int = 0
count := 0             // shorthand
config := struct{ port int }{port: 3000}
```

### Functions

```typescript
// TypeScript
function greet(name: string): string {
    return `Hello, ${name}!`;
}

const add = (a: number, b: number): number => a + b;
```

```go
// Go
func greet(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}

// No arrow functions, but functions are first-class
add := func(a, b int) int { return a + b }
```

### Async/Await vs Goroutines

```typescript
// TypeScript
async function fetchData(): Promise<Data> {
    const response = await fetch(url);
    return response.json();
}

// Parallel
const [users, posts] = await Promise.all([
    fetchUsers(),
    fetchPosts()
]);
```

```go
// Go: goroutines + channels
func fetchData() (Data, error) {
    resp, err := http.Get(url)
    if err != nil {
        return Data{}, err
    }
    // ...
}

// Parallel with goroutines
var wg sync.WaitGroup
wg.Add(2)
go func() { users = fetchUsers(); wg.Done() }()
go func() { posts = fetchPosts(); wg.Done() }()
wg.Wait()
```

### Error Handling

```typescript
// TypeScript
try {
    const data = await fetchData();
} catch (error) {
    console.error(error);
}
```

```go
// Go: explicit error returns (no exceptions)
data, err := fetchData()
if err != nil {
    log.Println(err)
    return
}
```

### Arrays/Slices

```typescript
// TypeScript
const numbers: number[] = [1, 2, 3];
numbers.push(4);
const doubled = numbers.map(n => n * 2);
const filtered = numbers.filter(n => n > 2);
```

```go
// Go: slices (no map/filter built-in)
numbers := []int{1, 2, 3}
numbers = append(numbers, 4)

// Manual map
doubled := make([]int, len(numbers))
for i, n := range numbers {
    doubled[i] = n * 2
}
```

### Objects vs Structs

```typescript
// TypeScript
interface User {
    id: number;
    name: string;
    email?: string;  // optional
}

const user: User = { id: 1, name: "Alice" };
```

```go
// Go
type User struct {
    ID    int
    Name  string
    Email string  // use pointer for optional: *string
}

user := User{ID: 1, Name: "Alice"}
```

### Classes vs Structs + Methods

```typescript
// TypeScript
class UserService {
    private db: Database;

    constructor(db: Database) {
        this.db = db;
    }

    async getUser(id: number): Promise<User> {
        return this.db.findUser(id);
    }
}
```

```go
// Go: no classes, use structs + methods
type UserService struct {
    db *Database
}

func NewUserService(db *Database) *UserService {
    return &UserService{db: db}
}

func (s *UserService) GetUser(id int) (*User, error) {
    return s.db.FindUser(id)
}
```

## Popular Package Equivalents

| Use Case | Node.js | Go |
|----------|---------|-----|
| HTTP server | Express, Fastify | net/http, Gin, Chi, Echo |
| HTTP client | axios, fetch | net/http, resty |
| ORM | Prisma, TypeORM | GORM, sqlx, Ent |
| Validation | zod, joi | go-playground/validator |
| Config | dotenv | viper, envconfig |
| CLI | commander, yargs | cobra, urfave/cli |
| Logging | winston, pino | zerolog, zap, slog |
| Testing | jest, vitest | testing (built-in) |
| Mocking | jest.mock | testify/mock, gomock |

## What You'll Miss

- No `map`, `filter`, `reduce` (you write loops)
- No optional chaining (`user?.profile?.name`)
- No nullish coalescing (`value ?? default`)
- No spread operator (`...array`)
- No destructuring (`const { name } = user`)

## What You'll Gain

- Single binary deployment (no `node_modules`)
- Faster cold starts
- Built-in concurrency (goroutines)
- Built-in testing and formatting
- Explicit error handling (no surprise exceptions)
- Static typing without compilation step

## Next Steps

1. Read [01-go-modules.md](01-go-modules.md) - the Go equivalent of npm
2. Try adding a dependency with `go get`
3. Create a simple HTTP server

Good luck with your Go journey!
