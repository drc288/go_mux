package database

import (
	"github.com/drc288/go_mux/models"
	"golang.org/x/crypto/bcrypt"
)

/**
 * LoginPass - this function validate the user and the password
 * @email: the user email to validate
 * @password: the passord of the user
 *
 * Return: (user, if exists true if not false)
 */
func LoginPass(email string, password string) (models.User, bool) {
	myUser, yes, _ := CheckEmail(email)
	if yes == false {
		return myUser, false
	}

	passwordByte := []byte(password)
	passwordDB := []byte(myUser.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordByte)
	if err != nil {
		return myUser, false
	}

	return myUser, true

}
