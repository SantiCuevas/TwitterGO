package middlew

import (
	"net/http"

	"github.com/santicuevas/TwitterGO/routers"
)

// ValidoJWT nos permite validar el jwt que viene en la peticion
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("authorization"))
		if err != nil {
			http.Error(w, "Error en el token: "+err.Error(), http.StatusBadRequest)
		}
		next.ServeHTTP(w, r)
	}
}
