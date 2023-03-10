package contract

import "go-mongo/pkg/validator"

type Response struct {
	OK         bool        `json:"ok"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	OK                bool                          `json:"ok"`
	StatusCode        int                           `json:"status_code"`
	Error             string                        `json:"error"`
	StructErrorFields []validator.StructErrorFields `json:"error_fields"`
}
