package auth

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
