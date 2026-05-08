// Package service contains the core authentication business logic.
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

// UserProvider is the storage interface that must be implemented by the caller.
// It abstracts the database layer so the library works with any storage backend.
type UserProvider interface {
	// SaveUser persists a new user to the storage.
	SaveUser(ctx context.Context, user *models.User) error
	// UserByEmail retrieves a user by their email address.
	UserByEmail(ctx context.Context, email string) (*models.User, error)
}

// TokenManager is the interface for JWT token generation.
type TokenManager interface {
	// NewToken generates a signed JWT token for the given user and duration.
	NewToken(user models.User, duration time.Duration) (string, error)
}

// AuthService handles user registration and login logic.
type AuthService struct {
	storage      UserProvider
	hasher       hasher.PasswordHasher
	tokenManager TokenManager
}

// NewAuthService creates a new AuthService with the given dependencies.
func NewAuthService(storage UserProvider, hasher hasher.PasswordHasher, tokenManager TokenManager) *AuthService {
	return &AuthService{storage: storage, hasher: hasher, tokenManager: tokenManager}
}

// Register creates a new user with the given email and password.
// The password is hashed with bcrypt before being stored.
func (s *AuthService) Register(ctx context.Context, email, password string) error {
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
// Returns an error if the user is not found or the password is incorrect.
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
