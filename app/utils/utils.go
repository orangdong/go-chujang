package utils

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
