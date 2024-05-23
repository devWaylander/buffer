package model

// модель данных для прокси запроса
type SaveFact struct {
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

// конвертация в map
func (SV *SaveFact) SaveFactToFormV1() map[string]string {
	return map[string]string{
		"period_start":            SV.PeriodStart,
		"period_end":              SV.PeriodEnd,
		"period_key":              SV.PeriodKey,
		"indicator_to_mo_id":      SV.IndicatorToMoID,
		"indicator_to_mo_fact_id": SV.IndicatorToFactID,
		"value":                   SV.Value,
		"fact_time":               SV.FactTime,
		"is_plan":                 SV.IsPlan,
		"auth_user_id":            SV.AuthUserID,
		"comment":                 SV.Comment,
	}
}

// Мок факт
var MockJson = SaveFact{
	PeriodStart:       "2024-05-01",
	PeriodEnd:         "2024-05-31",
	PeriodKey:         "month",
	IndicatorToMoID:   "227373",
	IndicatorToFactID: "0",
	Value:             "1",
	FactTime:          "2024-05-31",
	IsPlan:            "0",
	AuthUserID:        "40",
	Comment:           "buffer Chubakov",
}
