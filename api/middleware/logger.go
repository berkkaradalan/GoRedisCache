package middleware

import (
	// "log"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.Method, " ", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}