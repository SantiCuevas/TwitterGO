package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/santicuevas/TwitterGO/bd"
)

func LeoTweets(res http.ResponseWriter, req *http.Request) {
	Id := req.URL.Query().Get("id")
	if len(Id) < 1 {
		http.Error(res, "Debe enviar el parametro id por query", http.StatusBadRequest)
		return
	}
	pagina := req.URL.Query().Get("pagina")
	if len(pagina) < 1 {
		http.Error(res, "Debe enviar el parametro pagina por query", http.StatusBadRequest)
		return
	}

	pag, err := strconv.Atoi(pagina)
	if err != nil {
		http.Error(res, "Pagin debe ser un valor numerico mayor a 0", http.StatusBadRequest)
		return
	}

	page64 := int64(pag)
	respuesta, correcto := bd.LeoTweets(Id, page64)
	if correcto == false {
		http.Error(res, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	res.Header().Set("content-type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(respuesta)
}
