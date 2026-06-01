package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "Hotel Booking API is running",
		})
	}).Methods(http.MethodGet)

	return r
}
