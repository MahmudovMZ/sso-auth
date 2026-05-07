package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct{}

func (b BcryptHasher) Hash(pass string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func (b BcryptHasher) Compare(pass string, hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, []byte(pass))
	if err != nil {
		return false, err
	}
	return true, nil
}
