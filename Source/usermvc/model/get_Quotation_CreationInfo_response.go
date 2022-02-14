package model

type GetQuoationCreateInfoResp struct {
	Status  int
	Payload interface{}
}

type GetQuoteDetails struct {
	Accountid              string `json:"accountid"`
	Accountname            string `json:"accountname"`
	Accounttypename        string `json:"accounttypename"`
	Contactname            string `json:"contactname"`
	Contactid              string `json:"contactid"`
	Currencycode           string `json:"currencycode"`
	Currencyid             string `json:"currencyid"`
	Paymentterms           string `json:"payment_terms"`
	Remarksfrommarketing   string `json:"remarks_marketing"`
	Remarksfromgmc         string `json:"remarks_gmc"`
	Destination            string `json:"destination_port"`
	PortLoading            string `json:"port_loading"`
	Otherspecifications    string `json:"other_specification"`
	Billingaddress         string `json:"billing_address"`
	Destinationcountryid   string `json:"destination_countryid"`
	CreatedDate            string `json:"createddate"`
	Createdby              string `json:"createdby"`
	Incoterms              string `json:"incoterms"`
	Incotermsid            string `json:"incotermsid"`
	Finalclientaccountid   string `json:"finalclientaccountid"`
	Finalclientaccountname string `json:"finalclientaccountname"`
	Portloadingid          string `json:"portloadingid"`
	Portdestinationid      string `json:"destinationid"`
	Currencyname           string `json:"currencyname"`
	Fromdate               string `json:"fromdate"`
	Todate                 string `json:"todate"`
	Status                 string `json:"status"`
}

type InCotermsInfo struct {
	Incotermsid string `json:"incotermsid"`
	Incoterms   string `json:"incoterms"`
}

type Currencies struct {
	Currencyid   string `json:"currencyid"`
	Currencyname string `json:"currencyname"`
	Currencycode string `json:"currencycode"`
}

type Loadingports struct {
	Id              int    `json:"id"`
	Portlaodingname string `json:"portloading_name"`
}

type Destinationports struct {
	Destinationid int    `json:"id"`
	Destination   string `json:"destination_port"`
}

type AccountDetails struct {
	Accountid       int              `json:"account_id"`
	Accounttypeid   string           `json:"accounttype_id"`
	Accounttypename string           `json:"accounttype_name"`
	Address         string           `json:"billing_address"`
	Accountname     string           `json:"account_name"`
	Contacts        []ContactDetails `json:"contact_details"`
}

type ContactDetails struct {
	Contactid   int    `json:"contact_id"`
	Contactname string `json:"contact_name"`
}
