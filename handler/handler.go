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
	SaveFact(ctx context.Context, data model.SaveFact) (model.Response, error)
}

func Configure(ctx context.Context, mux *http.ServeMux, service Service) {
	// Прокси ручка для сохранения факта в БД
	mux.HandleFunc("POST /api/v1/proxy/save_fact", func(w http.ResponseWriter, r *http.Request) {
		body := model.SaveFact{}
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

		resp, err := service.SaveFact(ctx, body)
		if err != nil {
			log.Println(err)
			http.Error(w, internalErrors.ErrFailedToSendReq, http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(resp)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, internalErrors.ErrMarshalResp, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(data); err != nil {
			log.Println(err)
		}
	})
}
