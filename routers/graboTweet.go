package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/santicuevas/TwitterGO/bd"
	"github.com/santicuevas/TwitterGO/models"
)

// GraboTweet permite grabar el tweet en la BD
func GraboTweet(res http.ResponseWriter, req *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(req.Body).Decode(&message)

	registro := models.GraboTweet{
		UserId:  IdUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(res, "Ocurrio un error al intentar insertar un tweet: "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(res, "No se ha logrado insertar el tweet: "+err.Error(), 400)
		return
	}

	res.WriteHeader(http.StatusCreated)
}
