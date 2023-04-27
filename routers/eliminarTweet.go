package routers

import (
	"net/http"

	"github.com/santicuevas/TwitterGO/bd"
)

func EliminarTweet(res http.ResponseWriter, req *http.Request) {
	Id := req.URL.Query().Get("id")
	if len(Id) < 1 {
		http.Error(res, "Debe enviar un id por query", http.StatusBadRequest)
		return
	}
	err := bd.BorroTweet(Id, IdUser)
	if err != nil {
		http.Error(res, "Ha ocurrido un error al borrar el tweet: "+err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("content-type", "application/json")
	res.WriteHeader(http.StatusCreated)
}
