package model

type GetAllAccountsResponseBody struct {
	Accountname   string `json:"accountname"`
	Aliases       string `json:"aliases"`
	Accounttypeid string `json:"accounttypeid"`
	AccountOwner  string `json:"account_owner"`
	Masterstatus  string `json:"masterstatus"`
}

type GetAllAccountDetailsResponse struct {
	Status  int
	Payload []*GetAllAccountsResponseBody
}
