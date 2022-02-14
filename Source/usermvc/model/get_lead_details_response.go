package model

type GetLeadDetailsResponse struct {
	Status  int
	Payload interface{}
}

type LeadDetails struct {
	Accountname                string                `json:"accountname"`
	Aliases                    string                `json:"aliases"`
	Accounttypeid              string                `json:"accounttypeid"`
	Website                    string                `json:"website"`
	Approximativeannualrevenue string                `json:"approxannualrev"`
	Productsegmentid           string                `json:"productsegmentid"`
	ContactSalutationid        int64                 `json:"contact_saluatationid"`
	Salutations                Salutations           `json:"salutation"`
	Contactfirstname           string                `json:"contactfirstname"`
	Contactlastname            string                `json:"contactlastname"`
	ContactPosition            string                `json:"contact_position"`
	ContactEmail               string                `json:"email"`
	ContactPhone               string                `json:"phone"`
	ContactMobile              string                `json:"contact_mobile"`
	Manfacunit                 int                   `json:"manfacunit"`
	Instcoffee                 int                   `json:"instcoffee"`
	Price                      int64                 `json:"price"`
	Leadscore                  int64                 `json:"leadscore"`
	Coffeetypeid               string                `json:"coffeetypeid"`
	Productsegment             []ProductSegments     `json:"Productsegment"`
	CoffeeTypes                []CoffeeTypes         `json:"coffeetypes"`
	AccountTypes               []AccountsInformation `json:"accounttypes"`
	OtherInformation           string                `json:"otherinformation"`
	BillingStreetAddress       string                `json:"billing_street"`
	BillingCity                string                `json:"billing_citycode"`
	BillingState               string                `json:"billing_statecode"`
	BillingPostalCode          string                `json:"billing_postalcode"`
	BillingCountry             string                `json:"billing_countrycode"`
	ContactStreetAddress       string                `json:"contact_street"`
	ContactCity                string                `json:"contact_citycode"`
	ContactState               string                `json:"contact_statecode"`
	ContactPostalCode          string                `json:"contact_postalcode"`
	ContactCountry             string                `json:"contact_countrycode"`
	ShippingContinent          string                `json:"shipping_continent"`
	ShippingCountry            string                `json:"shipping_country"`
	Status                     string                `json:"status"`
	Contact_extid              int64                 `json:"contact_extid"`
	Contact_ext                PhoneCodes            `json:"contact_extcode"`
}
