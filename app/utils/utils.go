package utils

import (
	"fmt"
	"reflect"
	"strings"
	"time"

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

func UpdatedFieldsMap(data interface{}) map[string]interface{} {
	updateFields := make(map[string]interface{})
	now := time.Now().UTC()
	v := reflect.ValueOf(data)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.IsNil() {
			var tag string
			tag = t.Field(i).Tag.Get("db")
			if tag == "" {
				tag = t.Field(i).Tag.Get("json")
			}
			if tag != "" {
				updateFields[tag] = field.Interface()
			}
		}
	}

	updateFields["updated_at"] = now
	return updateFields
}
