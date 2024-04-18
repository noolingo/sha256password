package sha256password

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"strings"
)

func generateSalt() (string, error) {
	buf := make([]byte, 8)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", buf), nil
}

func EncryptPassword(password string) (string, error) {
	salt, err := generateSalt()
	if err != nil {
		return "", err
	}
	saltedPassword := salt + password
	sumPassword := sha256.Sum256([]byte(saltedPassword))
	return fmt.Sprintf("%v:%v", salt, sumPassword), nil
}

func CompWithEncrypted(password, encpassword string) bool {
	encPasPart := strings.Split(encpassword, ":")
	if len(encPasPart) != 2 {
		return false
	}
	saltedPassword := encPasPart[0] + password
	sumPassword := sha256.Sum256([]byte(saltedPassword))
	return fmt.Sprintf("%v", sumPassword) == encPasPart[1]
}
