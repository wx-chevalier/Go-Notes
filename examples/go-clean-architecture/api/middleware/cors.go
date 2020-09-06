package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func CorsEveryWhere(mux http.Handler) http.Handler {
	return cors.Default().Handler(mux)
}
