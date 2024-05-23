package service

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
)

type service struct {
}

func New() *service {
	return &service{}
}

// Константные значения
const (
	// Размер буфера
	BufferSize = 10
	// URL до api endpoint
	URL = "https://development.kpi-drive.ru/_api/facts/save_fact"
)

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

// func (s *service) request(bearerToken string) (*http.Response, error) {
// 	// Формируем запрос
// 	contentType, body, err := s.createForm(form)
// 	if err != nil {
// 		println("failed to create form-data")
// 	}

// 	req, err := http.NewRequest("POST", URL, body)
// 	if err != nil {
// 		println("failed to create request on %s %v", URL, err)
// 		return nil, err
// 	}

// 	// Добавляем заголовки с авторизацией и типом контента
// 	req.Header.Add("Authorization", "Bearer "+bearerToken)
// 	req.Header.Add("Content-Type", contentType)
// 	req.Header.Add("Accept", "*/*")
// 	req.Header.Add("User-Agent", "buffer client")

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		println("failed to get response from %s %v", URL, err)
// 		return nil, err
// 	}

// 	return resp, nil
// }

func (s *service) SaveFact(ctx context.Context) {
	// Запускаем запросы
	// for i := 0; i < BufferSize; i++ {
	// 	resp, err := request(BearerToken)
	// 	if err != nil {
	// 		continue
	// 	}

	// 	// Логируем код ответа
	// 	println("Code:", resp.Status)

	// 	// Декодируем полученный body response в случае 200
	// 	if resp.StatusCode == 200 {
	// 		response := model.Response{}
	// 		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
	// 			println("failed to decode response")
	// 		}

	// 		println("IndicatorToMoFactID:", response.Data.IndicatorToMoFactID)
	// 	}

	// 	resp.Body.Close()
	// }
}
