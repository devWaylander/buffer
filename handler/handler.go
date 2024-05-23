package handler

import (
	"context"
	"net/http"
)

type Service interface {
	SaveFact(ctx context.Context)
}

func Configure(ctx context.Context, mux *http.ServeMux, service Service) {
	// Прокси ручка для сохранения факта в БД
	mux.HandleFunc("POST /api/v1/proxy/save_fact", func(w http.ResponseWriter, r *http.Request) {
		println("Hi!")
	})
}
