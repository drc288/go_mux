package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/drc288/go_mux/models"
)

func GenerateJWT(t models.User) (string, error) {
	myPassword := []byte("GO_MUX_PROJECT_TEST")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastName":  t.LastName,
		"birthDate": t.BirthDate,
		"biography": t.Biography,
		"country":   t.Country,
		"website":   t.WebSite,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myPassword)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
