package routes

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/drc288/go_mux/database"
	"github.com/drc288/go_mux/models"
)

var (
	Email     string
	IDUsuario string
)

// Process the token to extract the values
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myPassword := []byte("GO_MUX_PROJECT_TEST")

	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")[1]
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("the format in the token are invalid")
	}

	tk = strings.TrimSpace(splitToken)

	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return myPassword, nil
	})

	if err == nil {
		_, ok, ID := database.CheckEmail(claims.Email)
		if ok {
			Email = claims.Email
			IDUsuario = ID
		}
		return claims, ok, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Invalid Token")
	}

	return claims, false, string(""), err
}
