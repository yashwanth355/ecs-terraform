package entity

type CmsLeadsBillingAddressMaster struct {
	Billingid     string `json:"billingid"`
	City          string `json:"city"`
	Country       string `json:"country"`
	Leadid        string `json:"leadid"`
	Postalcode    string `json:"postalcode"`
	Stateprovince string `json:"stateprovince"`
	Street        string `json:"street"`
}
