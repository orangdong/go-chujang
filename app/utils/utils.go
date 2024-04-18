package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type SuccessResp struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{
		Status:  "error",
		Message: message,
	}
}

func NewSuccessResponse(message string, data interface{}) *SuccessResp {
	return &SuccessResp{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

func Validate(data interface{}, validate *validator.Validate) string {
	var validationErrors strings.Builder

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			elemError := fmt.Sprintf("error field: %s failed on tag: %s with value: %s;", err.Field(), err.Tag(), err.Value().(string))
			validationErrors.WriteString(elemError)
		}
	}

	return validationErrors.String()
}
