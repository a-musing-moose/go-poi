package main

import (
	"fmt"
	"net/http"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	w.length = len(b)
	return w.ResponseWriter.Write(b)
}

func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		method := r.Method
		url := r.URL
		writer := statusWriter{w, 0, 0}
		h.ServeHTTP(&writer, r)
		end := time.Now()
		latency := end.Sub(start)
		length := writer.length
		statusCode := writer.status
		fmt.Printf("[%s] %d %s %d Bytes %v\n", method, statusCode, url, length, latency)
	})
}
