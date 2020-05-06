package main

import (
	"log"

	"github.com/akarnisdrakari/backend_golang/bd"
	"github.com/akarnisdrakari/backend_golang/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
