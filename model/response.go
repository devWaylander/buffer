package model

// Структуры ответа на запрос
type Data struct {
	IndicatorToMoFactID int `json:"indicator_to_mo_fact_id"`
}
type Response struct {
	Data Data `json:"DATA"`
}
