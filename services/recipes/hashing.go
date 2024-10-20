package recipes

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"

	"github.com/sathirak/garm/internal/errx"
	"golang.org/x/crypto/argon2"
)

// The full code for authentication using email and password

type HashSalt struct {
	Hash []byte
	Salt []byte
}

type Argon2idHash struct {
	// time represents the number of
	// passed over the specified memory.
	time uint32
	// cpu memory to be used.
	memory uint32
	// threads for parallelism aspect
	// of the algorithm.
	threads uint8
	// keyLen of the generate hash key.
	keyLen uint32
	// saltLen the length of the salt used.
	saltLen uint32
}

// NewArgon2idHash constructor function for
// Argon2idHash.
func NewArgon2idHash(time, saltLen uint32, memory uint32, threads uint8, keyLen uint32) *Argon2idHash {
	return &Argon2idHash{
		time:    time,
		saltLen: saltLen,
		memory:  memory,
		threads: threads,
		keyLen:  keyLen,
	}
}

func randomSecret(length uint32) ([]byte, error) {
	secret := make([]byte, length)

	if _, err := rand.Read(secret); err != nil {
		return nil, err
	}

	return secret, nil
}

// GenerateHash using the password and provided salt.
// If not salt value provided fallback to random value
// generated of a given length.
func (a *Argon2idHash) GenerateHash(password, salt []byte) (*HashSalt, error) {
	var err error
	// If salt is not provided generate a salt of
	// the configured salt length.
	if len(salt) == 0 {
		salt, err = randomSecret(a.saltLen)
	}

	if err != nil {
		return nil, err
	}
	// Generate hash
	hash := argon2.IDKey(password, salt, a.time, a.memory, a.threads, a.keyLen)
	// Return the generated hash and salt used for storage.
	return &HashSalt{Hash: hash, Salt: salt}, nil
}

// Compare generated hash with store hash.
func (a *Argon2idHash) Compare(hash, salt, password []byte) error {
	// Generate hash for comparison.

	hashSalt, err := a.GenerateHash(password, salt)
	if err != nil {
		return err
	}
	// Compare the generated hash with the stored hash.
	// If they don't match return error.
	if !bytes.Equal(hash, hashSalt.Hash) {
		return errors.New("hash doesnt match")
	}
	return nil
}

func CreateEmailPassword(userID string, password string) (string, string, error) {

	argon := NewArgon2idHash(1, 16, 64*1024, 4, 32)
	hashSalt, err := argon.GenerateHash([]byte(password), nil)
	if err != nil {
		return "", "", err
	}
	hash := base64.StdEncoding.EncodeToString(hashSalt.Hash)
	salt := base64.StdEncoding.EncodeToString(hashSalt.Salt)
	return hash, salt, nil
}

func ValidateEmailPassword(hash string, salt string, password string) errx.Errx {
	argon := NewArgon2idHash(1, 16, 64*1024, 4, 32)

	decodedHash, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return errx.NewError(err, errx.ErrInternalServerErr)
	}

	decodedSalt, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return errx.NewError(err, errx.ErrInternalServerErr)
	}

	err = argon.Compare(decodedHash, decodedSalt, []byte(password))
	if err != nil {
		return errx.NewError(err, errx.ErrInvalidCredentials)
	}

	return errx.Nil()
}
