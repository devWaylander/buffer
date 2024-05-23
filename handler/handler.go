package handler

import (
	"buffer/model"
	"context"
	"encoding/json"
	"log"
	"net/http"

	internalErrors "buffer/errors"
)

type Service interface {
	SaveFact(ctx context.Context)
}

func Configure(ctx context.Context, mux *http.ServeMux, service Service) {
	// Прокси ручка для сохранения факта в БД
	mux.HandleFunc("POST /api/v1/proxy/save_fact", func(w http.ResponseWriter, r *http.Request) {
		body := model.SaveFactBody{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, internalErrors.ErrDecodeReq, http.StatusBadRequest)
			return
		}

		if body.AuthUserID == "" {
			http.Error(w, internalErrors.ErrAuthUserIDIsMissing, http.StatusBadRequest)
			return
		}
		if body.Comment == "" {
			http.Error(w, internalErrors.ErrCommentIsMissing, http.StatusBadRequest)
			return
		}
		if body.FactTime == "" {
			http.Error(w, internalErrors.ErrFactTimeIsMissing, http.StatusBadRequest)
			return
		}
		if body.IndicatorToFactID == "" {
			http.Error(w, internalErrors.ErrIndicatorToFactIDIsMissing, http.StatusBadRequest)
			return
		}
		if body.IndicatorToMoID == "" {
			http.Error(w, internalErrors.ErrIndicatorToMoIDIsMissing, http.StatusBadRequest)
			return
		}
		if body.IsPlan == "" {
			http.Error(w, internalErrors.ErrIsPlanIsMissing, http.StatusBadRequest)
			return
		}
		if body.PeriodEnd == "" {
			http.Error(w, internalErrors.ErrPeriodEndIsMissing, http.StatusBadRequest)
			return
		}
		if body.PeriodKey == "" {
			http.Error(w, internalErrors.ErrPeriodKeyIsMissing, http.StatusBadRequest)
			return
		}
		if body.PeriodStart == "" {
			http.Error(w, internalErrors.ErrPeriodStartIsMissing, http.StatusBadRequest)
			return
		}
		if body.Value == "" {
			http.Error(w, internalErrors.ErrValueIsMissing, http.StatusBadRequest)
			return
		}
	})
}
