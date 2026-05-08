// Package models contains the data structures used by the sso-auth library.
package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents an authenticated user in the system.
// This struct is passed to the UserProvider interface methods
// and is populated during registration and login.
type User struct {
	// ID is the unique identifier for the user, generated automatically on registration.
	ID uuid.UUID
	// Username is an optional display name for the user.
	Username string
	// Email is the user's email address, used as the primary login identifier.
	Email string
	// PasswordHash is the bcrypt hash of the user's password.
	PasswordHash []byte
	// CreatedAt is the timestamp when the user was registered.
	CreatedAt time.Time
	// UpdatedAt is the timestamp of the last update to the user record.
	UpdatedAt time.Time
}
