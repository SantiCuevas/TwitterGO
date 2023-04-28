package routers

import (
	"net/http"

	"github.com/santicuevas/TwitterGO/bd"
	"github.com/santicuevas/TwitterGO/models"
)

// AltaRelacion realiza el registro de la relacion entre usuarios
func AltaRelacion(res http.ResponseWriter, req *http.Request) {
	Id := req.URL.Query().Get("id")
	if len(Id) < 1 {
		http.Error(res, "Debe enviar el parametro Id por query", 400)
		return
	}

	var t models.Relation
	t.UserId = IdUser
	t.UserRelationId = Id

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(res, "Error al insertar relacion: "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(res, "No se ha logrado insertar la relacion "+err.Error(), 400)
		return
	}
	res.WriteHeader(http.StatusCreated)
}
