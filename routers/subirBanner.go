package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/santicuevas/TwitterGO/bd"
	"github.com/santicuevas/TwitterGO/models"
)

func SubirAvatar(res http.ResponseWriter, req *http.Request) {
	file, handler, err := req.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IdUser + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(res, "Error al subir la imagen: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(res, "Error al subir la imagen: "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = IdUser + "." + extension
	status, err = bd.ModificoRegistro(user, IdUser)
	if err != nil || status == false {
		http.Error(res, "Error al subir el avatar a la BD: "+err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
}
