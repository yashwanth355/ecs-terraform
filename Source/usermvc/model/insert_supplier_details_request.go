package model

type VendorDetails struct {
	Create        bool   `json:"create"`
	Update        bool   `json:"update"`
	View          bool   `json:"view"`
	CreatedUserID string `json:"createduserid"`
	// CreatedUserName string `json:"createdusername"`
	AutoGenID           string `json:"autogen_id"`
	VendorId            string `json:"vendor_id"`
	VendorIdSno         int    `json:"vendor_idsno"`
	VendorName          string `json:"vendor_name"`
	VendorType          int    `json:"vendor_type"`
	VendorCategory      int    `json:"vendor_cat_name"`
	Lastvendorid        int    `json:"lastvendorid"`
	VendorGroup         string `json:"vendor_group"`
	VendorTypeid        string `json:"vendor_type_id"`
	VendorCategoryid    string `json:"vendor_cat_id"`
	VendorGroupid       string `json:"vendor_group_id"`
	PanNo               string `json:"pan_no"`
	GSTIdentificationNo string `json:"gst_no"`
	MSMESSI             string `json:"msmessi"`
	BankName            string `json:"bank_name"`
	Branch              string `json:"branch"`
	AccountType         string `json:"account_type"`
	AccountNumber       string `json:"account_number"`
	IfscCode            string `json:"ifsc_code"`
	MicrCode            string `json:"micr_code"`
	ContactName         string `json:"contact_name"`
	Address1            string `json:"address1"`
	Address2            string `json:"address2"`
	City                string `city:"city"`
	Pincode             string `json:"pincode"`
	State               string `json:"state"`
	Country             string `json:"country"`
	Phone               string `json:"phone"`
	Mobile              string `json:"mobile"`
	Email               string `json:"email"`
	Website             string `json:"website"`
	//AuditLogDetails []AuditLogSupplier `json:"audit_log_vendor"`
	VendorCategoryName string `json:"vendorcatname"`
	VendorGroupName    string `json:"groupname"`
}
type VendorID struct {
	ID string `json:"vendorid"`
}
