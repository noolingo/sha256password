package sha256password

import "crypto/sha256"

func PassHash(password string) string {
	hashpass := sha256.Sum256([]byte(password))
	return string(hashpass[:])
}

func CheckPass(verpassword, passwordbd string) bool {
	tmp := sha256.Sum224([]byte(verpassword))
	verpassword = string(tmp[:])
	if verpassword == passwordbd {
		return true
	}
	return false
}
