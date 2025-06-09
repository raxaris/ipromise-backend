package dto

type PredictionResponse struct {
	PromiseID   string  `json:"promise_id"`
	SuccessRate float64 `json:"success_rate"`
	Advice      string  `json:"advice"`
}
