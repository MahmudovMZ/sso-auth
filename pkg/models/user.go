package models

import (
	"time"

	"github.com/google/uuid"
)
// User represents an authenticated user in the system.
type User struct {
	ID           uuid.UUID
	Username     string
	Email        string
	PasswordHash []byte
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
