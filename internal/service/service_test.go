package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/MahmudovMZ/sso-auth/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockStorage struct{}

func (m *mockStorage) SaveUser(ctx context.Context, user *models.User) error {
	return nil
}

func (m *mockStorage) UserByEmail(ctx context.Context, email string) (*models.User, error) {
	return &models.User{
		Email:        email,
		PasswordHash: []byte("fake_hash"),
	}, nil
}

func (m *mockStorage) UserByID(ctx context.Context, id string) (*models.User, error) {
	return nil, nil
}

type mockHasher struct{}

func (m *mockHasher) Hash(password string) ([]byte, error) {
	return []byte("fake_hash"), nil
}

func (m *mockHasher) Compare(password string, hash []byte) (bool, error) {
	return true, nil
}

func TestAuthService_Register(t *testing.T) {
	//Arrange
	storage := &mockStorage{}
	hasher := &mockHasher{}
	tokens := &mockTockenManager{}
	s := NewAuthService(storage, hasher, tokens)

	//Act
	err := s.Register(context.Background(), "test@email.com", "password123")

	//Assert
	assert.NoError(t, err)
}

type mockErrorStorage struct {
	mockStorage
}

type mockTockenManager struct{}

func (m *mockTockenManager) NewToken(user models.User, duration time.Duration) (string, error) {
	return "fake_token", nil
}
func (m *mockErrorStorage) SaveUser(ctx context.Context, user *models.User) error {
	return errors.New("database connection failed")
}

func TestAuthService_Register2(t *testing.T) {
	//Arrange
	storage2 := &mockErrorStorage{}
	hasher2 := &mockHasher{}
	tokens := &mockTockenManager{}
	s := NewAuthService(storage2, hasher2, tokens)

	//Act
	err := s.Register(context.Background(), "test@email.com", "123")

	//Assert
	require.Error(t, err)
	assert.Equal(t, "database connection failed", err.Error())
}

func TestAuthServiceLogin(t *testing.T) {
	//Arrange
	storage3 := &mockStorage{}
	hasher3 := &mockHasher{}
	tokens := &mockTockenManager{}

	s := NewAuthService(storage3, hasher3, tokens)
	//Act
	token, err := s.Login(context.Background(), "test@email.com", "password123")

	assert.NoError(t, err)
	assert.Equal(t, "fake_token", token)
}
