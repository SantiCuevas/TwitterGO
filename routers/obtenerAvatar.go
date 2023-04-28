package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/santicuevas/TwitterGO/bd"
)

func ObtenerAvatar(res http.ResponseWriter, req *http.Request) {
	Id := req.URL.Query().Get("id")
	if len(Id) < 1 {
		http.Error(res, "Debe enviar el parametro Id por query", 400)
		return
	}

	perfil, err := bd.BuscoPerfil(Id)
	if err != nil {
		http.Error(res, "Error al buscar el perfil en la BD", 400)
		return
	}

	OpenFile, err := os.Open("uploads/avatars/" + perfil.Avatar)
	if len(Id) < 1 {
		http.Error(res, "Imagen no encontrada", 400)
		return
	}

	_, err = io.Copy(res, OpenFile)
	if err != nil {
		http.Error(res, "Error al copiar la imagen", 400)
	}

}
