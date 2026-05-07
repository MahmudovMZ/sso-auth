package jwt

import (
	"errors"
	"time"

	"github.com/MahmudovMZ/sso-auth/pkg/models"
	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secret string
}

func New(secret string) (*JWTManager, error) {
	if secret == "" {
		return nil, errors.New("jwt secret can not be empty")
	}
	return &JWTManager{secret: secret}, nil
}

func (j *JWTManager) NewToken(user models.User, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID.String(),
		"exp": time.Now().Add(duration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
