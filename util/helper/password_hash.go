package helper

import (
	"crypto/sha1"
	"fmt"
)

const salt = "hjqrhjqw124617aj564u564a654u65465aufhajs"

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
