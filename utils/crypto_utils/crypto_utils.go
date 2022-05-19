package crypto_utils

import "golang.org/x/crypto/bcrypt"

func GetHash(input string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(input), 8)
	return string(hashed)
}
