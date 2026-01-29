# Structs in Go

Structs are Go's way to group related data. They replace classes, interfaces, and objects from JS/TS.

## Defining a Struct

```go
type User struct {
    ID    int
    Name  string
    Email string
}
```

## Creating Instances

```go
// Named fields (recommended)
user := User{
    ID:    1,
    Name:  "Alice",
    Email: "alice@example.com",
}

// Positional (avoid - breaks if fields change)
user := User{1, "Alice", "alice@example.com"}

// Partial (unset fields get zero values)
user := User{Name: "Bob"}  // ID=0, Email=""

// Empty (all zero values)
var user User  // ID=0, Name="", Email=""
```

## Accessing Fields

```go
user := User{ID: 1, Name: "Alice"}

// Read
fmt.Println(user.Name)  // "Alice"

// Write
user.Email = "alice@new.com"
```

## Capitalization = Visibility

**This is fundamental to Go**: the first letter determines if something is public or private.

| First Letter | Visibility | JS/TS Equivalent |
|--------------|------------|------------------|
| `Uppercase` | Exported (public) | `export`, `public` |
| `lowercase` | Unexported (private) | no export, `private` |

This applies to **everything**: structs, fields, functions, variables, constants.

```go
type User struct {
    ID       int      // Exported - other packages can access
    Name     string   // Exported
    password string   // Unexported - only this package can access
}

func PublicFunc() {}   // Exported
func privateFunc() {}  // Unexported

var GlobalVar = 1      // Exported
var localVar = 2       // Unexported
```

### Why?

No keywords like `public`/`private`. Go uses naming convention instead - you can tell visibility at a glance.

```go
// In package "user"
type User struct {
    Name     string   // Other packages can read/write
    password string   // Only user package can access
}
```

```go
// In package "main"
import "myapp/user"

u := user.User{Name: "Alice"}
fmt.Println(u.Name)      // ✓ Works
fmt.Println(u.password)  // ✗ Compile error: unexported field
```

## Methods on Structs

Go has no classes, but structs can have methods.

```go
type User struct {
    Name  string
    Email string
}

// Method with receiver
func (u User) Greet() string {
    return "Hello, " + u.Name
}

// Usage
user := User{Name: "Alice"}
fmt.Println(user.Greet())  // "Hello, Alice"
```

### Value vs Pointer Receiver

```go
// Value receiver - gets a COPY
func (u User) GetName() string {
    return u.Name
}

// Pointer receiver - can MODIFY original
func (u *User) SetName(name string) {
    u.Name = name  // Modifies the original
}
```

See [07-pointers.md](07-pointers.md) for more on pointers.

## Comparison with JS/TS

### Object Literal → Struct

```typescript
// TypeScript
const user = { id: 1, name: "Alice", email: "a@test.com" };
```

```go
// Go
user := User{ID: 1, Name: "Alice", Email: "a@test.com"}
```

### Interface → Struct

```typescript
// TypeScript
interface User {
    id: number;
    name: string;
}
```

```go
// Go
type User struct {
    ID   int
    Name string
}
```

### Class → Struct + Methods

```typescript
// TypeScript
class UserService {
    private db: Database;

    constructor(db: Database) {
        this.db = db;
    }

    getUser(id: number): User {
        return this.db.find(id);
    }
}

const svc = new UserService(db);
svc.getUser(1);
```

```go
// Go
type UserService struct {
    db *Database
}

// Constructor function (convention: NewXxx)
func NewUserService(db *Database) *UserService {
    return &UserService{db: db}
}

// Method
func (s *UserService) GetUser(id int) (*User, error) {
    return s.db.Find(id)
}

// Usage
svc := NewUserService(db)
svc.GetUser(1)
```

## Nested Structs

```go
type Address struct {
    Street  string
    City    string
    Country string
}

type User struct {
    Name    string
    Address Address
}

user := User{
    Name: "Alice",
    Address: Address{
        Street:  "123 Main St",
        City:    "Tokyo",
        Country: "Japan",
    },
}

fmt.Println(user.Address.City)  // "Tokyo"
```

## Embedding (Composition)

Go prefers composition over inheritance.

```go
type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person         // Embedded (no field name)
    EmployeeID int
}

emp := Employee{
    Person:     Person{Name: "Alice", Age: 30},
    EmployeeID: 12345,
}

// Fields are "promoted" - access directly
fmt.Println(emp.Name)       // "Alice" (promoted from Person)
fmt.Println(emp.Age)        // 30
fmt.Println(emp.EmployeeID) // 12345
```

## Struct Tags (Metadata)

Tags add metadata for encoding/decoding (JSON, database, etc.).

```go
type User struct {
    ID        int    `json:"id" db:"user_id"`
    Name      string `json:"name"`
    Email     string `json:"email,omitempty"`
    Password  string `json:"-"`  // Exclude from JSON
}
```

```go
user := User{ID: 1, Name: "Alice", Email: ""}
data, _ := json.Marshal(user)
// {"id":1,"name":"Alice"}
// Note: email omitted (omitempty), password excluded (-)
```

### Common Tags

| Tag | Library | Purpose |
|-----|---------|---------|
| `json:"name"` | encoding/json | JSON field name |
| `db:"column"` | sqlx, gorm | Database column name |
| `yaml:"name"` | yaml.v3 | YAML field name |
| `validate:"required"` | go-playground/validator | Validation rules |

## Zero Values

Every type has a zero value. No `undefined` or `null` surprises.

```go
var user User
// user.ID = 0
// user.Name = ""
// user.Email = ""
```

| Type | Zero Value |
|------|------------|
| `int`, `float64` | `0` |
| `string` | `""` |
| `bool` | `false` |
| `pointer`, `slice`, `map` | `nil` |

## Anonymous Structs

For one-off use (like inline object literals in JS).

```go
// No type definition needed
config := struct {
    Host string
    Port int
}{
    Host: "localhost",
    Port: 8080,
}

fmt.Println(config.Host)  // "localhost"
```

## Summary

| Concept | Go |
|---------|-----|
| Define | `type Name struct { ... }` |
| Create | `Name{Field: value}` |
| Public field | `Uppercase` |
| Private field | `lowercase` |
| Method | `func (r Type) Name() {}` |
| Constructor | `func NewType() *Type {}` (convention) |
| Inheritance | Embedding (composition) |
| Metadata | Struct tags |

## Next

See [07-pointers.md](07-pointers.md) to understand pointers and when to use `*User` vs `User`.
