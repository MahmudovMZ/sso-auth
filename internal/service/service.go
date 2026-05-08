package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/MahmudovMZ/sso-auth/internal/hasher"
	"github.com/MahmudovMZ/sso-auth/pkg/models"
	"github.com/google/uuid"
)

// UserProvider is the interface that must be implemented by the caller.
// It abstracts the database layer so the library stays storage-agnostic.
type UserProvider interface {
	SaveUser(ctx context.Context, user *models.User) error
	// SaveUser persists a new user to the storage.
	UserByEmail(ctx context.Context, email string) (*models.User, error)
	// UserByEmail retrieves a user by their email address.
	UserByID(ctx context.Context, id string) (*models.User, error)
}
type TokenManager interface {
	NewToken(user models.User, duration time.Duration) (string, error) //New token generation
}

type AuthService struct { //Struct for the DataBase and the hash service
	storage      UserProvider
	hasher       hasher.PasswordHasher
	tokenManager TokenManager
}

func NewAuthService(storage UserProvider, hasher hasher.PasswordHasher, tokenManager TokenManager) *AuthService { //returning service's data such as storage(postgreSQL, MySQL)
	return &AuthService{storage: storage, hasher: hasher, tokenManager: tokenManager}
}

// Register creates a new user with the given email and password.
// The password is hashed using bcrypt before being stored.
func (s *AuthService) Register(ctx context.Context, email, password string) error { //filling the user's data
	pass, err := s.hasher.Hash(password)
	if err != nil {
		return err
	}

	user := models.User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: pass,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := s.storage.SaveUser(ctx, &user); err != nil {
		return err
	}
	return nil
}

// Login authenticates a user and returns a signed JWT token on success.
// Returns an error if the credentials are invalid.
func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.storage.UserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	ok, err := s.hasher.Compare(password, user.PasswordHash)
	if err != nil {
		return "", fmt.Errorf("failed to compare password: %w", err)
	}
	if !ok {
		return "", errors.New("invalid credentials")
	}
	token, err := s.tokenManager.NewToken(*user, time.Hour*24)
	if err != nil {
		return "", err
	}
	return token, nil
}
