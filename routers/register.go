package routers

import (
	"encoding/json"
	"net/http"

	"github.com/santicuevas/TwitterGO/bd"
	"github.com/santicuevas/TwitterGO/models"
)

// Registro es la funcion para crear en la BD el registro del usuario
func Register(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario con este email", 400)
		return
	}

	_, status, _ := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar el registro del usuario "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
