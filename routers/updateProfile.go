package routers

import (
	"encoding/json"
	"net/http"

	"github.com/santicuevas/TwitterGO/bd"
	"github.com/santicuevas/TwitterGO/models"
)

// UpdateProfile modifica el perfil del usuario
func UpdateProfile(res http.ResponseWriter, req *http.Request) {

	var user models.User

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(res, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool

	status, err = bd.ModificoRegistro(user, IdUser)
	if err != nil {
		http.Error(res, "Ocurrio un error al intentar modificar el perfil, reintente nuevamente "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(res, "No se ha logrado modificar el perfil del usuario", 400)
		return
	}

	res.WriteHeader(http.StatusCreated)
}
