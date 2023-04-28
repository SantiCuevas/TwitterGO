package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/santicuevas/TwitterGO/bd"
)

func LeoTweetsFollowers(res http.ResponseWriter, req *http.Request) {
	if len(req.URL.Query().Get("pagina")) < 1 {
		http.Error(res, "Debe enviar el parametro pagina", 400)
		return
	}
	pagina, err := strconv.Atoi(req.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(res, "Debe enviar el parametro pagina como entero mayor a 0", 400)
		return
	}

	respuesta, correcto := bd.LeoTweetsFollowers(IdUser, pagina)
	if correcto == false {
		http.Error(res, "Error al leer los tweets", 400)
		return
	}
	res.Header().Set("content-type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(respuesta)
}
