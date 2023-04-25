package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/santicuevas/TwitterGO/models"
)

// GeneroJWT genera el encriptado con JWT
func GeneroJWT(t models.User) (string, error) {
	clave := []byte("aprendiendoGOLANG")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"surName":   t.SurName,
		"birthDate": t.BirthDate,
		"biography": t.Biography,
		"location":  t.Location,
		"webSite":   t.WebSite,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(clave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
