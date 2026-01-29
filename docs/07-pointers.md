# Pointers in Go

Yes, Go has pointers like C/C++, but **safer** - no pointer arithmetic.

## Why Should JS/TS Devs Care?

In JavaScript, objects are always passed by reference:

```javascript
function updateUser(user) {
    user.name = "Bob";  // Modifies original
}
```

In Go, structs are passed by **value** (copied) by default:

```go
func updateUser(user User) {
    user.Name = "Bob"  // Modifies a COPY, original unchanged!
}
```

Pointers let you modify the original.

## The Basics

| Symbol | Meaning | Example |
|--------|---------|---------|
| `*T` | Pointer to type T | `*User` = pointer to User |
| `&x` | Get address of x | `&user` = pointer to user |
| `*p` | Dereference (get value) | `*ptr` = the User it points to |

```go
user := User{Name: "Alice"}

ptr := &user        // ptr is *User (pointer to user)
fmt.Println(*ptr)   // Dereference: {Name: "Alice"}

(*ptr).Name = "Bob" // Modify via pointer
// Or simply:
ptr.Name = "Bob"    // Go auto-dereferences for you
```

## Value vs Pointer

### Without Pointer (Copy)

```go
func birthday(u User) {
    u.Age++  // Modifies the COPY
}

user := User{Name: "Alice", Age: 30}
birthday(user)
fmt.Println(user.Age)  // Still 30! Original unchanged
```

### With Pointer (Reference)

```go
func birthday(u *User) {
    u.Age++  // Modifies the ORIGINAL
}

user := User{Name: "Alice", Age: 30}
birthday(&user)        // Pass pointer
fmt.Println(user.Age)  // 31! Modified
```

## JS/TS Mental Model

```typescript
// JavaScript - objects are always references
let user = { name: "Alice" };
let ref = user;        // ref points to same object
ref.name = "Bob";
console.log(user.name); // "Bob"
```

```go
// Go - explicit choice
user := User{Name: "Alice"}

// Copy (like spreading in JS)
copy := user           // Independent copy
copy.Name = "Bob"
fmt.Println(user.Name) // "Alice" (unchanged)

// Reference (like JS default)
ptr := &user           // Points to same data
ptr.Name = "Bob"
fmt.Println(user.Name) // "Bob" (changed)
```

## Creating Pointers

```go
// Method 1: & operator
user := User{Name: "Alice"}
ptr := &user

// Method 2: new() - allocates and returns pointer
ptr := new(User)       // *User pointing to zero-value User
ptr.Name = "Alice"

// Method 3: Composite literal with &
ptr := &User{Name: "Alice"}  // Most common for structs
```

## nil Pointers

Pointers can be `nil` (like `null` in JS).

```go
var ptr *User          // nil by default
fmt.Println(ptr)       // <nil>

if ptr != nil {
    fmt.Println(ptr.Name)
}

// Dereferencing nil = panic (runtime crash)
fmt.Println(ptr.Name)  // PANIC: nil pointer dereference
```

### Safe Pattern

```go
func getUser(id int) *User {
    if id <= 0 {
        return nil  // Not found
    }
    return &User{ID: id, Name: "Alice"}
}

user := getUser(0)
if user == nil {
    fmt.Println("User not found")
    return
}
fmt.Println(user.Name)  // Safe to use
```

## When to Use Pointers

### Use Pointers When:

1. **You need to modify the original**
   ```go
   func (u *User) SetName(name string) {
       u.Name = name
   }
   ```

2. **Struct is large** (avoid copying overhead)
   ```go
   func process(data *LargeStruct) { ... }
   ```

3. **Value might be absent** (`nil` = not present)
   ```go
   func findUser(id int) *User {
       // Returns nil if not found
   }
   ```

4. **Consistency** - if some methods need pointers, use pointers for all

### Use Values When:

1. **Small structs** (copying is cheap)
   ```go
   type Point struct { X, Y int }
   func distance(a, b Point) float64 { ... }
   ```

2. **You want immutability** (function can't modify original)
   ```go
   func format(u User) string {
       return u.Name  // Can't accidentally modify
   }
   ```

3. **Maps and slices** (already reference types internally)

## Method Receivers

```go
type Counter struct {
    value int
}

// Value receiver - works on copy
func (c Counter) Get() int {
    return c.value
}

// Pointer receiver - modifies original
func (c *Counter) Increment() {
    c.value++
}
```

**Convention**: If one method needs a pointer receiver, use pointers for all methods on that type.

## Pointers and Slices/Maps

Slices and maps are already "reference-like" internally.

```go
func addItem(items []string) {
    items[0] = "modified"  // This DOES modify original!
}

names := []string{"Alice", "Bob"}
addItem(names)
fmt.Println(names[0])  // "modified"
```

But to append:

```go
func addItem(items []string) []string {
    return append(items, "Charlie")  // Must return new slice
}

names = addItem(names)
```

## Common Patterns

### Constructor Returns Pointer

```go
func NewUser(name string) *User {
    return &User{Name: name}
}

user := NewUser("Alice")  // *User
```

### Optional Fields with Pointers

```go
type Config struct {
    Port    int
    Timeout *int  // nil = use default
}

func getTimeout(c Config) int {
    if c.Timeout == nil {
        return 30  // default
    }
    return *c.Timeout
}
```

### Avoiding nil Checks

```go
// Helper function
func intPtr(i int) *int {
    return &i
}

config := Config{
    Port:    8080,
    Timeout: intPtr(60),
}
```

## Pointer Comparison with C/C++

| Feature | C/C++ | Go |
|---------|-------|-----|
| Pointer arithmetic | Yes (`ptr++`) | No |
| Manual memory management | Yes (`malloc/free`) | No (garbage collected) |
| Dangling pointers | Common bug | Not possible |
| Null pointer access | Undefined behavior | Panic (predictable) |

Go pointers are **safe** - you can't corrupt memory.

## Summary

| Concept | Syntax | Example |
|---------|--------|---------|
| Pointer type | `*T` | `var p *User` |
| Get address | `&x` | `ptr := &user` |
| Dereference | `*p` | `name := (*ptr).Name` |
| Auto-deref | `p.field` | `ptr.Name` (same as above) |
| nil check | `p == nil` | `if ptr != nil { }` |
| Create with new | `new(T)` | `ptr := new(User)` |
| Create inline | `&T{}` | `ptr := &User{Name: "A"}` |

## Quick Decision Guide

```
Need to modify original?
├─ Yes → Use pointer (*T)
└─ No
    ├─ Large struct? → Use pointer (performance)
    └─ Small struct? → Use value (simpler, safer)
```
