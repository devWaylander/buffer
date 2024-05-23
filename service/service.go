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
	// Мьютекс буфера
	bufferMu sync.RWMutex
	// Буфер запросов
	buffer map[uuid.UUID]map[string]string

	// Сигнал о том, что запрос записан в буфер
	signalStored chan struct{}
}

func New() *service {
	// Инициализация
	service := &service{
		buffer:       make(map[uuid.UUID]map[string]string),
		signalStored: make(chan struct{}),
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

// Сохранение запроса в буфер
func (s *service) storeToBuffer(index uuid.UUID, data map[string]string) {
	s.bufferMu.RLock()
	defer s.bufferMu.RUnlock()

	s.buffer[index] = data
	// Подача сигнала о том, что запрос записан
	s.signalStored <- struct{}{}
}

// Отправка из буфера
func (s *service) sendFromBuffer(index uuid.UUID) (*http.Response, error) {
	s.bufferMu.RLock()
	defer s.bufferMu.RUnlock()

	// Формирование form-data
	contentType, body, _ := s.createForm(s.buffer[index])

	resp, err := s.request(contentType, body)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Удаление отправленного запроса из буфера
func (s *service) deleteFromBuffer(index uuid.UUID) {
	s.bufferMu.Lock()
	defer s.bufferMu.Unlock()

	delete(s.buffer, index)
}

// Утилитарная функция конвертации map в form-data
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

// Функция отправки запроса
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

// Мок функция генерации запросов
func (s *service) MockSaveFact(ctx context.Context) {
	for i := 0; i < MockReqCount; i++ {
		resp, err := s.SaveFact(ctx, model.MockJson)
		if err != nil {
			println(err.Error())
		}

		println("IndicatorToMoFactID:", resp.Data.IndicatorToMoFactID)
	}
}

// Бизнес-функция для proxy API
func (s *service) SaveFact(ctx context.Context, data model.SaveFact) (model.Response, error) {
	// Конвертация json в map
	formMap := data.SaveFactToFormV1()
	// Индекс для буфера запросов
	uuid := uuid.New()

	// Сохранение в буфер
	go s.storeToBuffer(uuid, formMap)

	// Отправка запроса из буфера по сигналу
	var resp *http.Response
	signal := <-s.signalStored
	if signal == struct{}{} {
		var err error
		resp, err = s.sendFromBuffer(uuid)
		if err != nil {
			return model.Response{}, err
		}
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
}
