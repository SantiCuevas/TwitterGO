package main

import (
	"log"

	"github.com/santicuevas/TwitterGO/bd"
	"github.com/santicuevas/TwitterGO/handlers"
)

func main() {
	if bd.ChequeoConexion() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
