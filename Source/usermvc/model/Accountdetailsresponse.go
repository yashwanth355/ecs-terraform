package model

type AccountDetailsResponse struct {
	StatusCode int         `json:"status_code"`
	Payload    interface{} `json:"payload"`
}
