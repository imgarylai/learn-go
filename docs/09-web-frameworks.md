# Web Frameworks in Go

Coming from Express, Fastify, or Koa? Here's how Go web frameworks compare.

## The Standard Library: `net/http`

Unlike Node.js, Go's standard library has a production-ready HTTP server.

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            fmt.Fprintf(w, "List users")
        } else if r.Method == http.MethodPost {
            fmt.Fprintf(w, "Create user")
        }
    })

    http.ListenAndServe(":8080", nil)
}
```

Compare to Express:

```javascript
const express = require('express');
const app = express();

app.get('/', (req, res) => {
    res.send('Hello, World!');
});

app.get('/users', (req, res) => res.send('List users'));
app.post('/users', (req, res) => res.send('Create user'));

app.listen(8080);
```

## Popular Frameworks Comparison

| Framework | Node.js Equivalent | Style | Best For |
|-----------|-------------------|-------|----------|
| `net/http` | Node's `http` module | Minimal | Simple APIs, learning |
| Chi | Express (minimal) | Router-focused | REST APIs |
| Gin | Express/Fastify | Full-featured | Production APIs |
| Echo | Fastify | Performance | High-traffic APIs |
| Fiber | Express (API-compatible) | Express-like | Express migrants |

## Chi - Lightweight Router

Chi is like a minimal Express with just routing. Very idiomatic Go.

```go
package main

import (
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    r := chi.NewRouter()

    // Middleware (like Express middleware)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    // Routes
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello!"))
    })

    // Route groups (like Express Router)
    r.Route("/api/users", func(r chi.Router) {
        r.Get("/", listUsers)
        r.Post("/", createUser)
        r.Get("/{id}", getUser)      // URL params with {id}
        r.Put("/{id}", updateUser)
        r.Delete("/{id}", deleteUser)
    })

    http.ListenAndServe(":8080", r)
}

func getUser(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")  // Get URL parameter
    w.Write([]byte("User: " + id))
}
```

## Gin - Full-Featured Framework

Gin is like Express with batteries included. Most popular Go framework.

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type User struct {
    ID   string `json:"id"`
    Name string `json:"name" binding:"required"`
}

func main() {
    r := gin.Default()  // Includes logger and recovery middleware

    // Simple route
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Hello!"})
    })

    // Route group
    api := r.Group("/api")
    {
        users := api.Group("/users")
        {
            users.GET("", listUsers)
            users.POST("", createUser)
            users.GET("/:id", getUser)     // URL params with :id
            users.PUT("/:id", updateUser)
            users.DELETE("/:id", deleteUser)
        }
    }

    r.Run(":8080")
}

func getUser(c *gin.Context) {
    id := c.Param("id")  // Get URL parameter
    c.JSON(http.StatusOK, gin.H{"id": id})
}

func createUser(c *gin.Context) {
    var user User

    // Bind and validate JSON (like express-validator)
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, user)
}
```

Compare to Express:

```javascript
const express = require('express');
const app = express();
app.use(express.json());

app.get('/', (req, res) => {
    res.json({ message: 'Hello!' });
});

app.get('/api/users/:id', (req, res) => {
    res.json({ id: req.params.id });
});

app.post('/api/users', (req, res) => {
    const { name } = req.body;
    if (!name) {
        return res.status(400).json({ error: 'name required' });
    }
    res.status(201).json({ name });
});
```

## Echo - High Performance

Echo emphasizes performance and has a clean API.

```go
package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Routes
    e.GET("/", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{"message": "Hello!"})
    })

    // Group
    api := e.Group("/api")
    api.GET("/users/:id", getUser)

    e.Start(":8080")
}

func getUser(c echo.Context) error {
    id := c.Param("id")
    return c.JSON(http.StatusOK, map[string]string{"id": id})
}
```

## Fiber - Express-Like API

Fiber is designed to feel like Express. Easiest transition for Node devs.

```go
package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    // Looks just like Express!
    app.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"message": "Hello!"})
    })

    app.Get("/users/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")
        return c.JSON(fiber.Map{"id": id})
    })

    app.Post("/users", func(c *fiber.Ctx) error {
        user := new(User)
        if err := c.BodyParser(user); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": err.Error()})
        }
        return c.Status(201).JSON(user)
    })

    app.Listen(":8080")
}
```

## Middleware Patterns

### Express Middleware

```javascript
// Express
app.use((req, res, next) => {
    console.log(`${req.method} ${req.path}`);
    next();
});

app.use('/admin', authMiddleware);
```

### Go Middleware (Chi/standard)

```go
// Middleware signature: func(next http.Handler) http.Handler
func logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)  // Call next handler
    })
}

// Apply globally
r.Use(logger)

// Apply to group
r.Route("/admin", func(r chi.Router) {
    r.Use(authMiddleware)
    r.Get("/dashboard", dashboard)
})
```

### Gin Middleware

```go
func logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        log.Printf("%s %s", c.Request.Method, c.Request.URL.Path)
        c.Next()  // Call next handler
    }
}

r.Use(logger())
```

## JSON Handling

### Express

```javascript
app.use(express.json());

app.post('/users', (req, res) => {
    const { name, email } = req.body;
    // ...
});
```

### Go (Gin)

```go
type CreateUserRequest struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
}

func createUser(c *gin.Context) {
    var req CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    // req.Name, req.Email available
}
```

## Query Parameters

### Express

```javascript
// GET /search?q=golang&page=1
app.get('/search', (req, res) => {
    const { q, page = '1' } = req.query;
});
```

### Go (Gin)

```go
// GET /search?q=golang&page=1
func search(c *gin.Context) {
    q := c.Query("q")
    page := c.DefaultQuery("page", "1")
}
```

## Which Framework to Choose?

| If you want... | Choose |
|---------------|--------|
| Minimal dependencies | `net/http` + Chi |
| Most resources/tutorials | Gin |
| Best performance | Echo or Fiber |
| Easiest Express transition | Fiber |
| Most idiomatic Go | Chi or `net/http` |

## Getting Started

```bash
# Chi
go get github.com/go-chi/chi/v5

# Gin
go get github.com/gin-gonic/gin

# Echo
go get github.com/labstack/echo/v4

# Fiber
go get github.com/gofiber/fiber/v2
```

## Next Steps

1. Start with `net/http` to understand the basics
2. Try Chi for a minimal router
3. Move to Gin or Echo for production APIs
4. Consider Fiber if you miss Express syntax
