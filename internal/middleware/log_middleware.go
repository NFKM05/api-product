package middleware

import (
	"log"
	"net/http"
	"time"
)

// Função que registra log toda vez que uma rota é chamada via requisição
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s (%v)", r.Method, r.URL.Path, time.Since(start))
	})
}
