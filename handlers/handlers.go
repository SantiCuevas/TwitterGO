package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/santicuevas/TwitterGO/middlew"
	"github.com/santicuevas/TwitterGO/routers"
)

// Manejadores seteo mi puerto, el handler y pongo a escuchar al servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.ChequeoBD(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
