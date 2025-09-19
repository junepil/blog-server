package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		start := time.Now()

		defer func() {
			log.Printf(
				"\"%s %s %s\" %d %s",
				r.Method,
				r.RequestURI,
				r.Proto,
				ww.Status(),
				time.Since(start),
			)
		}()

		next.ServeHTTP(ww, r)
	})
}
