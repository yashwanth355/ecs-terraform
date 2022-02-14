package model

type CompanyNames struct {
	Compname string `json:"compname"`
}

type GetAllCompanyNamesResponse struct {
	Status  int
	Payload []*CompanyNames
}
