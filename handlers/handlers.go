package handlers
import (
	"log"
	"net/http"
	"os"

	"github.com/akarnisdrakari/backend_golang/middlew"
	"github.com/akarnisdrakari/backend_golang/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el handler y pongo escuchar al servidor */
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8787"
	}
	handler := cors.AllowAll().Handler(routers)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
