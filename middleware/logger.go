package middleware

import (
	"log"
	"net/http"
	"time"
)

// Обёртка логирования прокси
type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s --- %s --- %v", r.Method, r.URL.Path, time.Since(start))
}

func NewRequestLogger(wrappedHandler http.Handler) *Logger {
	return &Logger{wrappedHandler}
}
