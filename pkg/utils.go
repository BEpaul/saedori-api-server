package pkg

type ApiResponse struct {
	Message string      `json:"message"`
	ErrCode interface{} `json:"error_code"`
}

func NewApiResponse(message string, errCode interface{}) *ApiResponse {
	return &ApiResponse{
		Message: message,
		ErrCode: errCode,
	}
}
