package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/santicuevas/TwitterGO/bd"
	"github.com/santicuevas/TwitterGO/models"
)

// Email del usuario
var Email string

// Id del usuario
var IdUser string

// ProcesoToken proceso token para extraer sus valores
func ProcesoToken(t string) (*models.Claim, bool, string, error) {
	clave := []byte("aprendiendoGOLANG")
	claims := &models.Claim{}

	splitToken := strings.Split(t, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	t = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(t, claims, func(token *jwt.Token) (interface{}, error) {
		return clave, nil
	})
	if err == nil {
		_, encontrado, ID := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IdUser = ID
		}
		return claims, encontrado, IdUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
