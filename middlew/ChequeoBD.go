package middlew

import (
	"net/http"

	"github.com/akarnisdrakari/backend_golang/bd"
)

/*ChequeoBD es una funcion enfocada a la conexion de los middlewrs con mongodb*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida de la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
