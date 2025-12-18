# AuthKit

`authkit` is a **reusable Go library** for handling authentication via **JWT tokens**. It provides token generation and verification, enabling secure, stateless authentication across multiple Go applications.

## Features

* JWT token generation with custom claims
* JWT token verification with signature, expiry, issuer, and audience validation

## Installation

```bash
go get github.com/hardikm9850/authkit@v1.0.0
```

## Usage

```go
import "github.com/hardikm9850/authkit/jwt"

// Create JWT config
cfg := jwt.Config{
    Algorithm: jwt.HS256,
    Secret: "your-secret-key",
    Issuer: "your-app",
    Audience: "your-users",
    AccessTokenTTL: time.Hour * 24,
}

// Initialize manager
manager, _ := jwt.NewManager(cfg)

// Generate token
token, _ := manager.Generate(jwt.Claims{UserID: "123", Roles: []string{"user"}})

// Verify token
claims, _ := manager.Verify(token)
```
