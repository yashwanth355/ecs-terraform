package entity

type AccountMaster struct {
	Accountname       string `json:"accountname"`
	Accounttypeid     int    `json:"accounttypeid"`
	Phone             int    `json:"phone"`
	Email             string `json:"email"`
	Createddate       string `json:"createddate"`
	Createduserid     int    `json:"createduserid"`
	Modifieddate      string `json:"modifieddate"`
	Modifieduserid    string `json:"modifieduserid"`
	Showid            string `json:"showid"`
	Fax               string `json:"fax"`
	Approxannualrev   int    `json:"approxannualrev"`
	Website           string `json:"website"`
	Productsegmentid  int    `json:"productsegmentid"`
	Recordtypeid      string `json:"recordtypeid"`
	Masterstatus      string `json:"masterstatus"`
	AccountOwner      string `json:"account_owner"`
	ShippingCountry   string `json:"shipping_country"`
	ShippingContinent string `json:"shipping_continent"`
	Comments          string `json:"comments"`
	Aliases           string `json:"aliases"`
	Otherinformation  string `json:"otherinformation"`
	Autogencode       string `json:"autogencode"`
	Extlegacyid       string `json:"extlegacyid"`
	Isactive          int    `json:"isactive"`
	Accountid         int    `json:"accountid"`
	Custid            string `json:"custid"`
	RefCustid         string `json:"ref_custid"`
}
