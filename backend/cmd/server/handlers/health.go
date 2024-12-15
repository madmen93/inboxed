package handlers

import "net/http"

//HealthCheckHandler: Chequea que servidor se encuentre en buenas condiciones.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("El servidor está en óptimas condiciones."))
}
