package contract

import "go-mongo/pkg/validator"

type Response struct {
	OK         bool        `json:"ok"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	OK          bool                      `json:"ok"`
	StatusCode  int                       `json:"status_code"`
	Error       string                    `json:"error"`
	ErrorFields []validator.ErrorResponse `json:"error_fields"`
}
