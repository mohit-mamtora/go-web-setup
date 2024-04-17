package repository

import (
	"crypto/sha1"
	"encoding/hex"
)

func createHash(value string) (string, error) {
	sha := sha1.New()
	_, err := sha.Write([]byte(value))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sha.Sum(nil)), nil
}
