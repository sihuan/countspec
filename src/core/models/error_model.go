package models

type ErrorModel struct {
	ErrorCode int    `json:"code"`
	Message   string `json:"message"`
	// Hints     string `json:"hints"`
	// Info      string `json:"info"`
}
