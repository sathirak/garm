package hashing

import (
	"encoding/base64"

	"github.com/hotelbear/garm/internal/errx"
)

// The full code for authentication using email and password

func GenerateHashSalt(password string) (string, string, errx.Errx) {

	argon := NewArgon2idHash(1, 16, 64*1024, 4, 32)
	hashSalt, err := argon.GenerateHash([]byte(password), nil)

	if err != nil {
		return "", "", errx.NewError(err, errx.ErrInternalServer)
	}

	hash := base64.StdEncoding.EncodeToString(hashSalt.Hash)
	salt := base64.StdEncoding.EncodeToString(hashSalt.Salt)
	return hash, salt, errx.Nil()
}

func ValidateCredentials(hash string, salt string, password string) errx.Errx {
	argon := NewArgon2idHash(1, 16, 64*1024, 4, 32)

	decodedHash, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return errx.NewError(err, errx.ErrInternalServer)
	}

	decodedSalt, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return errx.NewError(err, errx.ErrInternalServer)
	}

	err = argon.Compare(decodedHash, decodedSalt, []byte(password))
	if err != nil {
		return errx.NewError(err, errx.ErrInvalidCredentials)
	}

	return errx.Nil()
}
