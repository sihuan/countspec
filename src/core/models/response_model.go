package models

import "github.com/gin-gonic/gin"

type ResponseModel struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data"`
	Errors  []ErrorModel `json:"errors"`
	// Endpoint string       `json:"endpoint"`
	// Method   string       `json:"method"`
}

func (r *ResponseModel) NewResponse() gin.H {
	return gin.H{
		"status":  r.Status,
		"message": r.Message,
		"data":    r.Data,
		"errors":  r.Errors,
		// "endpoint": r.Endpoint,
		// "method":   r.Method,
	}
}
