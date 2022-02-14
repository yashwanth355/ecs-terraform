package model

type GetLeadInfoResponse struct {
	Status  int
	Payload interface{}
}
type LeadInfo struct {
	Leadid           string `json:"leadid"`
	Accountname      string `json:"accountname"`
	Aliases          string `json:"aliases"`
	Contactfirstname string `json:"contactfirstname"`
	Contactlastname  string `json:"contactlastname"`
	Contact_Mobile   string `json:"contact_mobile"`
	Email            string `json:"email"`
	Leadscore        int    `json:"leadscore"`
	Masterstatus     string `json:"masterstatus"`
}
