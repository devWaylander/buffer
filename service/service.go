package service

import (
	"buffer/config"
	"buffer/model"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"sync"

	internalErrors "buffer/errors"

	"github.com/google/uuid"
)

type service struct {
	bufferMu sync.RWMutex
	buffer   map[uuid.UUID]map[string]string
}

func New() *service {
	service := &service{
		buffer: make(map[uuid.UUID]map[string]string),
	}

	return service
}

// Константные значения
const (
	// Количество мок запросов
	MockReqCount = 10
	// URL до api endpoint
	URL = "https://development.kpi-drive.ru/_api/facts/save_fact"
)

// func (s *service) runBuffer() {
// 	for {
// 		select {
// 		case bufferForm := <-s.bufferChan:
// 			s.bufferMu.RLock()
// 			s.buffer[bufferForm.index] = bufferForm.form
// 			s.bufferMu.RUnlock()

// 		case index := <-s.sendFromBufferSignal:
// 			s.bufferMu.RLock()

// 			contentType, body, _ := s.createForm(s.buffer[index])

// 			resp, err := s.request(contentType, body)
// 			if err != nil {
// 				s.responseChan <- resp
// 			}
// 			if resp.StatusCode == 200 {
// 				s.deleteFromBuffer(index)
// 			}

// 			s.responseChan <- resp

// 			s.bufferMu.RUnlock()

// 		case index := <-s.deleteFromBufferSignal:
// 			s.bufferMu.Lock()
// 			delete(s.buffer, index)
// 			s.bufferMu.Unlock()
// 		}
// 	}
// }

func (s *service) storeToBuffer(index uuid.UUID, data map[string]string) {
	s.bufferMu.RLock()
	defer s.bufferMu.RUnlock()

	s.buffer[index] = data
}

func (s *service) sendFromBuffer(index uuid.UUID) (*http.Response, error) {
	s.bufferMu.RLock()
	defer s.bufferMu.RUnlock()

	contentType, body, _ := s.createForm(s.buffer[index])

	resp, err := s.request(contentType, body)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *service) deleteFromBuffer(index uuid.UUID) {
	s.bufferMu.Lock()
	defer s.bufferMu.Unlock()

	delete(s.buffer, index)
}

func (s *service) createForm(form map[string]string) (string, io.Reader, error) {
	body := new(bytes.Buffer)
	mp := multipart.NewWriter(body)
	defer mp.Close()

	// записываем поля для FormData
	for key, val := range form {
		mp.WriteField(key, val)
	}

	return mp.FormDataContentType(), body, nil
}

func (s *service) request(contentType string, body io.Reader) (*http.Response, error) {
	// Формируем запрос
	req, err := http.NewRequest("POST", URL, body)
	if err != nil {
		println("failed to create request on %s %v", URL, err)
		return nil, err
	}

	// Добавляем заголовки с авторизацией и типом контента
	req.Header.Add("Authorization", "Bearer "+config.GlobalConfig.BearerToken)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "buffer client")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		println("failed to get response from %s %v", URL, err)
		return nil, err
	}

	return resp, nil
}

// набивать буфер из ручки и мока
// удалять из буфера только в случае успешного 200 от kpi

func (s *service) SaveFact(ctx context.Context, data model.SaveFact) (model.Response, error) {
	formMap := data.SaveFactToFormV1()
	uuid := uuid.New()

	s.storeToBuffer(uuid, formMap)

	resp, err := s.sendFromBuffer(uuid)
	if err != nil {
		return model.Response{}, err
	}

	// Декодируем полученный body response в случае 200
	response := model.Response{}
	if resp.StatusCode == 200 {
		s.deleteFromBuffer(uuid)

		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			println("failed to decode response")
		}
	} else {
		resp.Body.Close()
		return model.Response{}, errors.New(internalErrors.ErrFailedToSendReq)
	}

	resp.Body.Close()

	return response, nil

	// Запускаем запросы
	// for i := 0; i < BufferSize; i++ {
	// 	resp, err := request(BearerToken)
	// 	if err != nil {
	// 		continue
	// 	}

	// 	// Логируем код ответа
	// 	println("Code:", resp.Status)

	// // Декодируем полученный body response в случае 200
	// if resp.StatusCode == 200 {
	// 	response := model.Response{}
	// 	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
	// 		println("failed to decode response")
	// 	}

	// 	println("IndicatorToMoFactID:", response.Data.IndicatorToMoFactID)
	// }

	// 	resp.Body.Close()
	// }
}
