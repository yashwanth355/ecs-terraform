package entity

type LeadDetails struct {
	Accountname                string                `json:"accountname"`
	Aliases                    string                `json:"aliases"`
	Accounttypeid              string                `json:"accounttypeid"`
	Website                    string                `json:"website"`
	Approximativeannualrevenue string                `json:"approxannualrev"`
	Productsegmentid           string                `json:"productsegmentid"`
	ContactSalutationid        int                   `json:"contact_saluatationid"`
	Salutations                Salutations           `json:"salutation"`
	Contactfirstname           string                `json:"contact_firstname"`
	Contactlastname            string                `json:"contact_lastname"`
	ContactPosition            string                `json:"contact_position"`
	ContactEmail               string                `json:"email"`
	ContactPhone               string                `json:"phone"`
	ContactMobile              string                `json:"contact_mobile"`
	Manfacunit                 int                   `json:"manfacunit"`
	Instcoffee                 int                   `json:"instcoffee"`
	Price                      int                   `json:"sample_ready"`
	Leadscore                  int                   `json:"leadscore"`
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
	Contact_extid              string                `json:"contact_extid"`
	Contact_ext                PhoneCodes            `json:"contact_extcode"`
}

type Salutations struct {
	Salutationid string `json:"id"`
	Salutation   string `json:"salutation"`
}

type ProductSegments struct {
	Productsegmentid int    `json:"id"`
	Productsegment   string `json:"productsegment"`
}

type PhoneCodes struct {
	Id          int    `json:"id"`
	Countryname string `json:"countryname"`
	Dialcode    string `json:"dialcode"`
}
type AccountsInformation struct {
	Accounttypeid string `json:"id"`
	Accounttype   string `json:"accounttype"`
}

type CoffeeTypes struct {
	CoffeeType   string `json:"coffeetype"`
	CoffeeTypeId string `json:"id"`
}
