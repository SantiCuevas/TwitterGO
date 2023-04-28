package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/santicuevas/TwitterGO/bd"
)

// ListaUsuarios leo la lista de los usuarios
func ListaUsuarios(res http.ResponseWriter, req *http.Request) {
	typeUser := req.URL.Query().Get("type")
	page := req.URL.Query().Get("page")
	search := req.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(res, "Debe enviar el parametro como entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := bd.LeoUsuarios(IdUser, pag, search, typeUser)
	if status == false {
		http.Error(res, "Error al leer usuarios", http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(result)

}
