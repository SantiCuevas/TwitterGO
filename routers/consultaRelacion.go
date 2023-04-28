package routers

import (
	"encoding/json"
	"net/http"

	"github.com/santicuevas/TwitterGO/bd"
	"github.com/santicuevas/TwitterGO/models"
)

// ConsultaRelacion chequea que exista relacion entre 2 usuarios
func ConsultaRelacion(res http.ResponseWriter, req *http.Request) {
	Id := req.URL.Query().Get("id")
	if len(Id) < 1 {
		http.Error(res, "Debe enviar el parametro Id por query", 400)
		return
	}

	var t models.Relation
	t.UserId = IdUser
	t.UserRelationId = Id

	var resp models.ConsultaRelacion
	status, err := bd.ConsultoRelacion(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(resp)

}
