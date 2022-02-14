package model

import "encoding/json"

type GetAlleadsResonseBody struct {
	AccountName      string      `json:"accountname"`
	Aliases          string      `json:"aliases"`
	Contactfirstname string      `json:"contactfirstname"`
	Contactlastname  string      `json:"contactlastname"`
	Phone            json.Number `json:"phone"`
	Email            string      `json:"email"`
	ApprovalStatus   int64       `json:"approvalstatus"`
}

type GetAlleadsResonse struct {
	Status  int
	Payload []*GetAlleadsResonseBody
}
