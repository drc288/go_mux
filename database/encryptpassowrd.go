package database

import "golang.org/x/crypto/bcrypt"

func EncryptPassowrd(pass string) (string, error) {
	// 2 ^^ 8 => 256 iterations to encrypt
	// 6 to normal user
	// 8 to superadmin
	cost := 6
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
