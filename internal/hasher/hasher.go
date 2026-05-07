package hasher

type PasswordHasher interface {
	Hash(password string) ([]byte, error)
	Compare(password string, hash []byte) (bool, error)
}
