// Package auth provides a lightweight SDK for user authentication.
// It handles password hashing and JWT token generation under the hood.
// The caller only needs to implement the UserProvider interface.
package auth

import (
	"fmt"

	"github.com/MahmudovMZ/sso-auth/internal/hasher"
	"github.com/MahmudovMZ/sso-auth/internal/lib/jwt"
	"github.com/MahmudovMZ/sso-auth/internal/service"
)

// New creates a new authentication service.
// The secretKey is used to sign JWT tokens and must not be empty.
//
// Example:
//
// service, err := auth.New(storage, "my-secret-key")
//
//	if err != nil {
//	    log.Fatal(err)
//	}
func New(storage service.UserProvider, secretKey string) (*service.AuthService, error) {
	passHasher := hasher.BcryptHasher{}

	jwtManager, err := jwt.New(secretKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create jwt manager: %w", err)
	}

	return service.NewAuthService(storage, passHasher, jwtManager), nil
}
