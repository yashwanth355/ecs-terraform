package entity

type CmsDestinationMaster struct {
	Createdate    string `json:"createdate"`
	Createdby     string `json:"createdby"`
	Destination   string `json:"destination"`
	Destinationid int64  `json:"destinationid"`
	Isactive      string `json:"isactive"`
}
