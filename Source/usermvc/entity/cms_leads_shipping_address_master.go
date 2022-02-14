package entity

type CmsLeadsShippingAddressMaster struct {
	City          string `json:"city"`
	Country       string `json:"country"`
	Leadid        string `json:"leadid"`
	Postalcode    string `json:"postalcode"`
	Shippingid    string `json:"shippingid"`
	Stateprovince string `json:"stateprovince"`
	Street        string `json:"street"`
}
