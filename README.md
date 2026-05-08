# sso-auth

![CI](https://github.com/MahmudovMZ/sso-auth/actions/workflows/ci/yml/badge.svg)

> A lightweight, interface-driven Go SDK for JWT authentication and password hashing.

Built to be dropped into any Go project that needs user registration and login — without rewriting the same auth boilerplate every time.

---

## Features

- ✅ User registration with bcrypt password hashing
- ✅ User login with JWT token generation
- ✅ Interface-driven design — bring your own database
- ✅ Zero dependencies on any specific database or framework
- ✅ Clean, minimal public API

---

## Installation

```bash
go get github.com/MahmudovMZ/sso-auth@v0.1.0
```

---

## Quick Start

### 1. Implement the `UserProvider` interface

`sso-auth` doesn't care about your database. You just implement one interface:

```go
type UserProvider interface {
    SaveUser(ctx context.Context, user *models.User) error
    UserByEmail(ctx context.Context, email string) (*models.User, error)
}
```

Example with a hypothetical Postgres storage:

```go
type PostgresStorage struct {
    db *sql.DB
}

func (s *PostgresStorage) SaveUser(ctx context.Context, user *models.User) error {
    _, err := s.db.ExecContext(ctx,
        "INSERT INTO users (id, email, password_hash, created_at) VALUES ($1, $2, $3, $4)",
        user.ID, user.Email, user.PasswordHash, user.CreatedAt,
    )
    return err
}

func (s *PostgresStorage) UserByEmail(ctx context.Context, email string) (*models.User, error) {
    user := &models.User{}
    err := s.db.QueryRowContext(ctx,
        "SELECT id, email, password_hash FROM users WHERE email = $1", email,
    ).Scan(&user.ID, &user.Email, &user.PasswordHash)
    return user, err
}
```

### 2. Initialize the library

```go
import auth "github.com/MahmudovMZ/sso-auth"

storage := &PostgresStorage{db: db}

authService, err := auth.New(storage, "your-secret-key")
if err != nil {
    log.Fatal(err)
}
```

### 3. Register a user

```go
err := authService.Register(ctx, "user@example.com", "securepassword")
if err != nil {
    log.Println("registration failed:", err)
}
```

### 4. Login and get a JWT token

```go
token, err := authService.Login(ctx, "user@example.com", "securepassword")
if err != nil {
    log.Println("login failed:", err)
}

fmt.Println("token:", token)
```

---

## How It Works

```
Your App
   │
   ▼
auth.New(storage, secretKey)
   │
   ├── Register(email, password)
   │      └── bcrypt.Hash(password) → SaveUser(user)
   │
   └── Login(email, password)
          └── UserByEmail(email) → bcrypt.Compare(password) → jwt.NewToken()
```

All internals (bcrypt, JWT signing, token generation) are handled under the hood. You only deal with your database layer.

---

## User Model

```go
type User struct {
    ID           uuid.UUID
    Email        string
    PasswordHash []byte
    CreatedAt    time.Time
    UpdatedAt    time.Time
}
```

---

## Requirements

- Go 1.21+
- A JWT secret key (min. recommended: 32 characters)

---

## Running Tests

```bash
go test ./...
```

---

## License

MIT
