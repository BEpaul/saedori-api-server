package pkg

type ApiResponse struct {
	Message string `json:"message"`
}

func NewApiResponse(message string) *ApiResponse {
	return &ApiResponse{
		Message: message,
	}
}
