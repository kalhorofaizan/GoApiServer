package middleware

import (
	"log"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	status int
}

func (w *wrappedWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func LogApi(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrappedWriter := &wrappedWriter{ResponseWriter: w, status: http.StatusOK}
		handler.ServeHTTP(wrappedWriter, r)
		log.Println(wrappedWriter.status, r.Method, r.URL.Path, time.Since(start))
	})
}
