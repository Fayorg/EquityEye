package types

type ApiResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ApiErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type ApiDataResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
