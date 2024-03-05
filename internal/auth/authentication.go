package auth

import (
	"log"

	"github.com/alexedwards/argon2id"
)

func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}

	return hash, err
}

func ComparePasswordAndHash(password, encodedHash string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(password, encodedHash)
	if err != nil {
		log.Fatal(err)
	}

	return match, err
}
