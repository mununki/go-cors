package cors

import (
	"net/http"
)

// Set func sets CORS headers and returns http.Handler
func Set(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if (*r).Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
