package main

import (
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Константные значения
const (
	// Размер буфера
	bufferSize = 10
	// URL до api endpoint
	URL = "https://development.kpi-drive.ru/_api/facts/save_fact"
)

var (
	// тело запроса form/data
	form = url.Values{
		"period_start":            {"2024-05-01"},
		"period_end":              {"2024-05-31"},
		"period_key":              {"month"},
		"indicator_to_mo_id":      {"227373"},
		"indicator_to_mo_fact_id": {"0"},
		"value":                   {"1"},
		"fact_time":               {"2024-05-31"},
		"is_plan":                 {"0"},
		"auth_user_id":            {"40"},
		"comment":                 {"buffer Chubakov"},
	}
)

func request(bearerToken string) (*http.Response, error) {
	// Формируем запрос
	data := form.Encode()
	req, err := http.NewRequest("POST", URL, strings.NewReader(data))
	if err != nil {
		println("failed to create request on %s %v", URL, err)
		return nil, err
	}

	// Добавляем заголовки с авторизацией и типом контента
	req.Header.Add("Authorization", "Bearer "+bearerToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))
	req.Header.Add("Accept", "*/*")

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

		println(resp.StatusCode)
		resp.Body.Close()
	}
}
