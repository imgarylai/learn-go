# Testing in Go

Go has built-in testing support. No external frameworks required.

## TypeScript Comparison

| Concept | TypeScript (Jest/Vitest) | Go |
|---------|--------------------------|-----|
| Test runner | Jest, Vitest, Mocha | `go test` (built-in) |
| Test file | `*.test.ts`, `*.spec.ts` | `*_test.go` |
| Test function | `test()`, `it()` | `func TestXxx(t *testing.T)` |
| Assertions | `expect().toBe()` | `if got != want { t.Error() }` |
| Mocking | `jest.mock()` | Interfaces + manual mocks |
| Coverage | `--coverage` flag | `-cover` flag |
| Watch mode | `--watch` | Use external tools |

## Basics

### Test File Naming

```typescript
// TypeScript
math.ts
math.test.ts  // or math.spec.ts
```

```go
// Go
math.go
math_test.go  // Must end with _test.go
```

### Test Function Naming

```typescript
// TypeScript (Jest)
describe('add', () => {
    it('should add two numbers', () => {
        expect(add(2, 3)).toBe(5);
    });
});

// Or
test('add should add two numbers', () => {
    expect(add(2, 3)).toBe(5);
});
```

```go
// Go - must start with Test, takes *testing.T
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}

func TestAdd_WithNegativeNumbers(t *testing.T) {
    // Underscore for subtests in name
}
```

## Writing Tests

### Basic Test

```go
// math.go
package math

func Add(a, b int) int {
    return a + b
}
```

```go
// math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Add(2, 3) = %d; want 5", result)
    }
}
```

### No Built-in Assertions

Go doesn't have `expect()`. You write plain `if` statements:

```typescript
// TypeScript (Jest)
expect(result).toBe(5);
expect(result).toEqual({ name: 'Alice' });
expect(result).toBeTruthy();
expect(() => fn()).toThrow();
```

```go
// Go - plain if statements
if result != 5 {
    t.Errorf("got %d; want 5", result)
}

if !reflect.DeepEqual(result, expected) {
    t.Errorf("got %v; want %v", result, expected)
}

if result == nil {
    t.Error("expected non-nil result")
}
```

Popular assertion libraries if you want them:
- `github.com/stretchr/testify/assert`
- `github.com/matryer/is`

### Table-Driven Tests

The idiomatic Go way to test multiple cases (like `test.each` in Jest):

```typescript
// TypeScript (Jest)
test.each([
    [2, 3, 5],
    [-1, -2, -3],
    [0, 0, 0],
])('add(%i, %i) = %i', (a, b, expected) => {
    expect(add(a, b)).toBe(expected);
});
```

```go
// Go - table-driven tests
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -1, -2, -3},
        {"zero", 0, 0, 0},
        {"mixed", -1, 5, 4},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d",
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

### Subtests

`t.Run()` creates subtests (like `describe` blocks):

```typescript
// TypeScript
describe('Math', () => {
    describe('addition', () => {
        it('adds numbers', () => { ... });
    });
    describe('subtraction', () => {
        it('subtracts numbers', () => { ... });
    });
});
```

```go
// Go
func TestMath(t *testing.T) {
    t.Run("addition", func(t *testing.T) {
        // ...
    })

    t.Run("subtraction", func(t *testing.T) {
        // ...
    })
}
```

## Running Tests

```bash
# Run all tests in current directory
go test                          # Like: npm test

# Run all tests recursively
go test ./...                    # Like: npm test -- --all

# Verbose output
go test -v ./...                 # Like: npm test -- --verbose

# Run specific test
go test -run TestAdd ./...       # Like: npm test -- -t "add"

# Run specific subtest
go test -run TestAdd/positive ./...

# With coverage
go test -cover ./...             # Like: npm test -- --coverage

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out  # Opens in browser
```

## Benchmarks

Go has built-in benchmarking (no separate tool needed):

```typescript
// TypeScript - need separate tool like benchmark.js
import Benchmark from 'benchmark';
const suite = new Benchmark.Suite();
suite.add('add', () => add(2, 3));
```

```go
// Go - built-in
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

Run with:

```bash
go test -bench=. ./...
```

## Test Helpers

### Setup and Teardown

```typescript
// TypeScript (Jest)
beforeAll(() => { setup(); });
afterAll(() => { teardown(); });
beforeEach(() => { ... });
afterEach(() => { ... });
```

```go
// Go - TestMain for package-level setup
func TestMain(m *testing.M) {
    // Setup
    setup()

    code := m.Run()  // Run all tests

    // Teardown
    teardown()

    os.Exit(code)
}

// Per-test setup: just call functions
func TestSomething(t *testing.T) {
    cleanup := setupTest()
    defer cleanup()  // Runs after test
    // ...
}
```

### Helper Functions

```go
func assertEqual(t *testing.T, got, want int) {
    t.Helper()  // Marks this as a helper (error points to caller)
    if got != want {
        t.Errorf("got %d; want %d", got, want)
    }
}
```

### Temporary Directories

```typescript
// TypeScript
import tmp from 'tmp';
const dir = tmp.dirSync();
// Manual cleanup
```

```go
// Go - automatic cleanup
func TestWithTempDir(t *testing.T) {
    dir := t.TempDir()  // Automatically cleaned up
    // Use dir...
}
```

## Testing Private Functions

```go
// In mypackage/private.go
package mypackage

func privateFunc() int { return 42 }
```

```go
// In mypackage/private_test.go
package mypackage  // Same package = access private functions

func TestPrivateFunc(t *testing.T) {
    result := privateFunc()  // Can access!
    // ...
}
```

Or use `_test` suffix for black-box testing:

```go
// In mypackage/public_test.go
package mypackage_test  // Different package, only public API

import "mypackage"

func TestPublicAPI(t *testing.T) {
    // Can only test exported functions
}
```

## Mocking

Go uses interfaces for mocking (no magic like `jest.mock`):

```typescript
// TypeScript - mock modules
jest.mock('./database');
import { db } from './database';
(db.getUser as jest.Mock).mockReturnValue({ name: 'Alice' });
```

```go
// Go - use interfaces
type UserStore interface {
    GetUser(id int) (*User, error)
}

// Real implementation
type Database struct { ... }
func (d *Database) GetUser(id int) (*User, error) { ... }

// Mock for tests
type MockUserStore struct {
    Users map[int]*User
}
func (m *MockUserStore) GetUser(id int) (*User, error) {
    return m.Users[id], nil
}

// Test
func TestService(t *testing.T) {
    mock := &MockUserStore{
        Users: map[int]*User{1: {Name: "Alice"}},
    }
    svc := NewService(mock)  // Inject mock
    // ...
}
```

## Try It

```bash
go test -v ./...
```

## Next

See [05-tooling.md](05-tooling.md) for Go development tools.
