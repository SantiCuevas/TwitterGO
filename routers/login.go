package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/santicuevas/TwitterGO/bd"
	"github.com/santicuevas/TwitterGO/jwt"
	"github.com/santicuevas/TwitterGO/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña inválidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario y/o contraseña inválidos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generar el token "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	/* Guardar cookie en el usuario */
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "Token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
