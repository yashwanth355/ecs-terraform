package model

type LeadRequest struct {
	Accountid           int64  `json:"accountid"`
	Accountname         string `json:"accountname"`
	Accounttypeid       string `json:"accounttypeid"`
	Aliases             string `json:"aliases"`
	Approvalstatus      int64  `json:"approvalstatus"`
	Approverid          int64  `json:"approverid"`
	Approxannualrev     string `json:"approxannualrev"`
	Coffeetypeid        string `json:"coffeetypeid"`
	Comments            string `json:"comments"`
	ContactExt          string `json:"contact_ext"`
	ContactMobile       string `json:"contact_mobile"`
	ContactPosition     string `json:"contact_position"`
	ContactSalutationid int64  `json:"contact_salutationid"`
	Contactfirstname    string `json:"contactfirstname"`
	Contactlastname     string `json:"contactlastname"`
	Countryid           int64  `json:"countryid"`
	Createddate         string `json:"createddate"`
	Createduserid       int64  `json:"createduserid"`
	Email               string `json:"email"`
	Instcoffee          bool   `json:"instcoffee"`
	Isactive            bool   `json:"isactive"`
	Leadid              int    `json:"leadid"`
	Leadscore           int64  `json:"leadscore"`
	Manfacunit          bool   `json:"manfacunit"`
	Masterstatus        string `json:"masterstatus"`
	Modifieddate        string `json:"modifieddate"`
	Modifieduserid      int64  `json:"modifieduserid"`
	Otherinformation    string `json:"otherinformation"`
	Phone               string `json:"phone"`
	Price               int64  `json:"price"`
	Productsegmentid    string `json:"productsegmentid"`
	Recordtypeid        int64  `json:"recordtypeid"`
	ShippingContinent   string `json:"shipping_continent"`
	ShippingContinentid int64  `json:"shipping_continentid"`
	ShippingCountry     string `json:"shipping_country"`
	Website             string `json:"website"`
}

type LeadResponse struct {
	Status  int
	Message string
}
