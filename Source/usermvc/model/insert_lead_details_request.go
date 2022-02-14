package model

import "time"

type InsertLeadDetailsRequest struct {
	Update                     bool      `json:"update"`
	LeadId                     string    `json:"leadid"`
	Accountname                string    `json:"accountname"`
	Aliases                    string    `json:"aliases"`
	Accounttypeid              string    `json:"accounttypeid"`
	Website                    string    `json:"website"`
	Approximativeannualrevenue string    `json:"approxannualrev"`
	Productsegmentid           string    `json:"productsegmentid"`
	ContactSalutationid        int64     `json:"contact_salutationid"`
	Contactfirstname           string    `json:"contact_firstname"`
	Contactlastname            string    `json:"contact_lastname"`
	ContactPosition            string    `json:"contact_position"`
	ContactEmail               string    `json:"contact_email"`
	ContactPhone               string    `json:"contact_phone"`
	ContactMobile              string    `json:"contact_mobile"`
	Manfacunit                 int       `json:"manfacunit"`
	Instcoffee                 int       `json:"instcoffee"`
	Price                      int64     `json:"price"`
	Coffeetypeid               string    `json:"coffeetypeid"`
	OtherInformation           string    `json:"otherinformation"`
	BillingStreetAddress       string    `json:"billing_street"`
	BillingCity                string    `json:"billing_citycode"`
	BillingState               string    `json:"billing_statecode"`
	BillingPostalCode          string    `json:"billing_postalcode"`
	BillingCountry             string    `json:"billing_countrycode"`
	ContactStreetAddress       string    `json:"contact_street"`
	ContactCity                string    `json:"contact_citycode"`
	ContactState               string    `json:"contact_statecode"`
	ContactPostalCode          string    `json:"contact_postalcode"`
	ContactCountrycode         string    `json:"contact_countrycode"`
	CreatedDate                time.Time `json:"createddate"`
	CreatedUserid              string    `json:"createduserid"`
	CreatorsEmail              string    `json:"emailid"`
	ModifiedDate               time.Time `json:"modifieddate"`
	ModifiedUserid             int64     `json:"modifieduserid"`
	ShippingContinentid        string    `json:"shipping_continentid"`
	ShippingCountryid          string    `json:"countryid"`
	Leadscore                  int64     `json:"leadscore"`
	Masterstatus               string    `json:"masterstatus"`
	Approvalstatus             int64     `json:"approvalstatus"`
	ShippingContinent          string    `json:"shipping_continent"`
	ShippingCountry            string    `json:"shipping_country"`
	Contact_ext                string    `json:"contact_ext"`
	Isactive                   int       `json:"isactive"`
}
