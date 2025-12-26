# AuthKit

`authkit` is a **reusable Go library** for authentication in Go applications. It provides **JWT token generation & verification**, **password hashing**, and **middleware support**, enabling secure, stateless authentication across multiple Go services.


## Features

* JWT token generation & verification with configurable claims
* Password hashing & verification (bcrypt)
* Middleware support for **Gin** and **net/http**


## Installation

```bash
go get github.com/hardikm9850/authkit@v1.2.1
```


## Usage

### JWT Manager

```go
import (
    "time"
    "github.com/hardikm9850/authkit/jwt"
)

// Configure JWT
cfg := jwt.Config{
    Algorithm:      jwt.HS256,
    Secret:         "your-secret-key",
    Issuer:         "your-app",
    Audience:       "your-users",
    AccessTokenTTL: time.Hour * 24,
}

// Initialize manager
manager, err := jwt.NewManager(cfg)
if err != nil {
    panic(err)
}

// Generate token
claims := jwt.Claims{
    UserID: "123",
    Roles:  []string{"user"},
}
token, _ := manager.Generate(claims)

// Verify token
verifiedClaims, _ := manager.Verify(token)
```


### Password Hashing

```go
import "github.com/hardikm9850/authkit/password"

// Hash a password
hash, _ := password.HashPassword("mySecret123")

// Verify a password
err := password.VerifyPassword("mySecret123", hash) // nil if matches
```


### Middleware

#### Gin

```go
import (
    "github.com/gin-gonic/gin"
    "github.com/hardikm9850/authkit/jwt"
    "github.com/hardikm9850/authkit/middleware"
)

r := gin.Default()
r.Use(middleware.JWTAuth(manager))

r.GET("/protected", func(c *gin.Context) {
    userID, _ := c.Get("userID")
    roles, _ := c.Get("roles")
    c.JSON(200, gin.H{"userID": userID, "roles": roles})
})
```

#### net/http

```go
import (
    "net/http"
    "github.com/hardikm9850/authkit/middleware"
)

http.Handle("/protected", middleware.JWTAuthHTTP(manager)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("userID")
    w.Write([]byte("User ID: " + userID.(string)))
})))
```


## Installation Guide

```bash
go get github.com/hardikm9850/authkit@v1.1.0
```

* v1.1.0 adds:

  * Password hashing & verification
  * Middleware for Gin and net/http

* v1.0.0 was JWT only.

## License

MIT License
