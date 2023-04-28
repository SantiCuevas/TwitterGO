package routers

import (
	"net/http"

	"github.com/santicuevas/TwitterGO/bd"
	"github.com/santicuevas/TwitterGO/models"
)

// BajaRelacion realiza el borrado de la relacion entre usuarios
func BajaRelacion(res http.ResponseWriter, req *http.Request) {
	Id := req.URL.Query().Get("id")
	if len(Id) < 1 {
		http.Error(res, "Debe enviar el parametro Id por query", 400)
		return
	}

	var t models.Relation
	t.UserId = IdUser
	t.UserRelationId = Id

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(res, "Error al borrar relacion: "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(res, "No se ha logrado borrar la relacion "+err.Error(), 400)
		return
	}
	res.WriteHeader(http.StatusCreated)
}
