package auth
// Package auth provides a simple interface for user authentication.
// It handles password hashing and JWT token generation under the hood.
package auth

// New creates a new AuthService with the provided storage and JWT secret key.
// The secretKey must not be empty — it is used to sign JWT tokens.
//
// Example:
//
//	service, err := auth.New(storage, "my-secret-key")
//	if err != nil {
//	    log.Fatal(err)
//	}
import (
	"fmt"

	"github.com/MahmudovMZ/sso-auth/internal/hasher"
	"github.com/MahmudovMZ/sso-auth/internal/lib/jwt"
	"github.com/MahmudovMZ/sso-auth/internal/service"
)

func New(storage service.UserProvider, secretKey string) (*service.AuthService, error) {
	passHasher := hasher.BcryptHasher{}
	jwtManager, err := jwt.New(secretKey)
	if err != nil {
		fmt.Errorf("failed to create jwt manager")
	}
	return service.NewAuthService(storage, passHasher, jwtManager), nil
}
