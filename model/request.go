package model

// тело запроса из прокси
type SaveFactBody struct {
	PeriodStart       string `json:"period_start"`
	PeriodEnd         string `json:"period_end"`
	PeriodKey         string `json:"period_key"`
	IndicatorToMoID   string `json:"indicator_to_mo_id"`
	IndicatorToFactID string `json:"indicator_to_mo_fact_id"`
	Value             string `json:"value"`
	FactTime          string `json:"fact_time"`
	IsPlan            string `json:"is_plan"`
	AuthUserID        string `json:"auth_user_id"`
	Comment           string `json:"comment"`
}

// тело запроса form/data в API бд
var MockForm = map[string]string{
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