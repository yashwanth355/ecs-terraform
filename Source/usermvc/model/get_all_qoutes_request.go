package model

type GetAllQoutesRequestBody struct {
	Type      string `json:"type"`
	Createdby string `json:"created_userid"`
}
