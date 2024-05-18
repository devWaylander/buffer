package main

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Структуры ответа на запрос
type data struct {
	IndicatorToMoFactID int `json:"indicator_to_mo_fact_id"`
}
type response struct {
	Data data `json:"DATA"`
}

// Константные значения
const (
	// Размер буфера
	bufferSize = 10
	// URL до api endpoint
	URL = "http://development.kpi-drive.ru/_api/facts/save_fact"
)

var (
	// тело запроса form/data
	form = map[string]string{
		"period_start":            "2024-05-01",
		"period_end":              "2024-05-31",
		"period_key":              "month",
		"indicator_to_mo_id":      "227373",
		"indicator_to_mo_fact_id": "0",
		"value":                   "1",
		"fact_time":               "2024-05-31",
		"is_plan":                 "0",
		"auth_user_id":            "40",
		"comment":                 "buffer Chubakov",
	}
)

func createForm(form map[string]string) (string, io.Reader, error) {
	body := new(bytes.Buffer)
	mp := multipart.NewWriter(body)
	defer mp.Close()

	// записываем поля для FormData
	for key, val := range form {
		mp.WriteField(key, val)
	}

	return mp.FormDataContentType(), body, nil
}

func request(bearerToken string) (*http.Response, error) {
	// Формируем запрос
	contentType, body, err := createForm(form)
	if err != nil {
		println("failed to create form-data")
	}

	req, err := http.NewRequest("POST", URL, body)
	if err != nil {
		println("failed to create request on %s %v", URL, err)
		return nil, err
	}

	// Добавляем заголовки с авторизацией и типом контента
	req.Header.Add("Authorization", "Bearer "+bearerToken)
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

func main() {
	// Считываем env файл проекта
	err := godotenv.Load()
	if err != nil {
		println("failed to load .env file, using system env variables")
	}

	// Сохраняем bearer токен для доступа к API
	bearerToken := os.Getenv("KPI_API_BEARER")

	// Запускаем запросы
	for i := 0; i < bufferSize; i++ {
		resp, err := request(bearerToken)
		if err != nil {
			continue
		}

		// Логируем код ответа
		println("Code:", resp.Status)

		// Декодируем полученный body response в случае 200
		if resp.StatusCode == 200 {
			response := response{}
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				println("failed to decode response")
			}

			println("IndicatorToMoFactID:", response.Data.IndicatorToMoFactID)
		}

		resp.Body.Close()
	}
}
