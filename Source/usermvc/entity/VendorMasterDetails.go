package entity

type VendorDetails struct {
	Create        bool   `gorm:"column:create"`
	Update        bool   `gorm:"column:update"`
	View          bool   `gorm:"column:view"`
	CreatedUserID string `gorm:"column:createduserid"`
	// CreatedUserName string `json:"createdusername"`
	AutoGenID           string `gorm:"autogen_id"`
	VendorId            string `gorm:"column:vendor_id"`
	VendorIdSno         int    `gorm:"column:vendor_idsno"`
	VendorName          string `gorm:"column:vendor_name"`
	VendorType          int    `gorm:"column:vendor_type"`
	VendorCategory      int    `gorm:"column:vendor_cat_name"`
	Lastvendorid        int    `gorm:"column:lastvendorid"`
	VendorGroup         string `gorm:"column:vendor_group"`
	VendorTypeid        string `gorm:"column:vendor_type_id"`
	VendorCategoryid    string `gorm:"column:vendor_cat_id"`
	VendorGroupid       string `gorm:"column:vendor_group_id"`
	PanNo               string `gorm:"column:pan_no"`
	GSTIdentificationNo string `gorm:"column:gst_no"`
	MSMESSI             string `gorm:"column:msmessi"`
	BankName            string `gorm:"column:bank_name"`
	Branch              string `gorm:"column:branch"`
	AccountType         string `gorm:"column:account_type"`
	AccountNumber       string `gorm:"column:account_number"`
	IfscCode            string `gorm:"column:ifsc_code"`
	MicrCode            string `gorm:"column:micr_code"`
	ContactName         string `gorm:"column:contact_name"`
	Address1            string `gorm:"column:address1"`
	Address2            string `gorm:"column:address2"`
	City                string `gorm:"column:city"`
	Pincode             string `gorm:"column:pincode"`
	State               string `gorm:"column:state"`
	Country             string `gorm:"column:country"`
	Phone               string `gorm:"column:phone"`
	Mobile              string `gorm:"column:mobile"`
	Email               string `gorm:"column:email"`
	Website             string `gorm:"column:website"`
	//AuditLogDetails []AuditLogSupplier `json:"audit_log_vendor"`
	VendorCategoryName string `gorm:"column:vendorcatname"`
	VendorGroupName    string `gorm:"column:groupname"`
}
