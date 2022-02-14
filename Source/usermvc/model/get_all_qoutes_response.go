package model

type QuoteDetails struct {
	Quoteid     int    `json:"quotenumber"`
	Accountname string `json:"accountname"`
	Status      string `json:"status"`
	Createdby   string `json:"createdby"`
	Createddate string `json:"createddate"`
}

type GetAllQoutesResponse struct {
	Status  int
	Payload []*QuoteDetails
}
