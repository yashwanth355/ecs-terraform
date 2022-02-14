package model
import (
	// "database/sql"
	// "reflect"
	"database/sql/driver"
    "errors"

)
type PurchaseOrderDetails struct {
	Status          string `json:"status"`
	CreatedUserID   string `json:"createduserid"`
	GCCreatedUserID string `json:"gccreateduserid"`
	GCCoffeeType    string `json:"coffee_type"'`
	Type            string `json:"type"`
	//Contract Information
	Contract string `json:"contract"`
	//PO Info Section::
	POTypeID        string `json:"po_type_id"`
	PoId            string `json:"poid"`
	PoIdsNo         int    `json:"poidsno"`
	PoNO            string `json:"po_no"`
	PoNOsno         int    `json:"po_nosno"`
	PoDate          string `json:"po_date"`
	POCategory      string `json:"po_category"`
	POSubCategory   string `json:"po_sub_category"`
	SupplierTypeID  string `json:"supplier_type_id"`
	SupplierCountry string `json:"supplier_country"`
	//---------Currency & Advance Information//------------------
	CurrencyID   string `json:"currency_id"`
	CurrencyName string `json:"currency_name"`
	CurrencyCode string `json:"currency_code"`

	//Supplier/Vendor Information
	SupplierName    string `json:"supplier_name"`
	SupplierID      string `json:"supplier_id"`
	SupplierType    string `json:"supplier_type"`
	SupplierEmail   string `json:"supplier_email"`
	SupplierAddress string `json:"supplier_address"`

	//Vendor      			string `json:"supplier_id"`
	// VendorType  			string `json:"vendor_type"`
	QuotNo    string `json:"quot_no"`
	QuotDate  string `json:"quot_date"`
	QuotPrice string `json:"quot_price"`

	LastPoIdsno int `json:"last_poidsno"`
	//currency & incoterms
	IncoTermsID string `json:"incotermsid"`
	IncoTerms   string `json:"incoterms"`
	Origin      string `json:"origin"`
	PortOfLoad  string `json:"ports"`
	// TransportMode		 	string `json:"mode_of_transport"`
	Insurance          string `json:"insurance"`
	PlaceOfDestination string `json:"place_of_destination"`
	Forwarding         string `json:"forwarding"`
	NoOfContainers     string `json:"no_of_containers"`
	ContainerType      string `json:"container_type"`
	PaymentTerms       string `json:"payment_terms"`
	Comments           string `json:"comments"`
	PaymentTermsDays   string `json:"payment_terms_days"` //int to string
	//Billing & Delivery Info
	POBillTypeID   string `json:"billing_at_id"`
	POBillTypeName string `json:"billing_at_name"`
	POBillAddress  string `json:"billing_at_address"`
	PODelTypeID    string `json:"delivery_at_id"`
	PODelTypeName  string `json:"delivery_at_name"`
	PODelAddress   string `json:"delivery_at_address"`

	//Green Coffee Info Section-Done--------------------------

	ItemID        string `json:"item_id"`
	ItemName      string `json:"item_name"`
	TotalQuantity string `json:"total_quantity"`
	Density       string `json:"density"`
	Moisture      string `json:"moisture"`
	Browns        string `json:"browns"`
	Blacks        string `json:"blacks"`
	BrokenBits    string `json:"brokenbits"`
	InsectedBeans string `json:"insectedbeans"`
	Bleached      string `json:"bleached"`
	Husk          string `json:"husk"`
	Sticks        string `json:"sticks"`
	Stones        string `json:"stones"`
	BeansRetained string `json:"beansretained"`

	//Price Information-Done------------------------------

	PurchaseType       string `json:"purchase_type"`
	TerminalMonth      string `json:"terminal_month"`
	BookedTerminalRate string `json:"booked_terminal_rate"`
	BookedDifferential string `json:"booked_differential"`
	FixedTerminalRate  string `json:"fixed_terminal_rate"`
	FixedDifferential  string `json:"fixed_differential"`
	PurchasePrice      string `json:"purchase_price"`
	MarketPrice        string `json:"market_price"`
	POMargin           string `json:"po_margin"`
	// FinalPrice			 string `json:"final_price"`

	Advance     string `json:"advance"`      //changed
	AdvanceType string `json:"advance_type"` //changed
	PoQty       string `json:"po_qty"`
	// Price 				 string `json:"price"`

	ApprovalStatus bool `json:"approval_status"`

	//GC Information-Dispatch Section

	DispatchType  string `json:"dispatch_type"`
	DispatchCount string `json:"dispatch_count"`

	LastDetIDSNo int    `json:"last_det_ids_no"`
	DetIDSNo     int    `json:"det_ids_no"`
	DetID        string `json:"det_id_no"`
	// DispatchID			string `json:"dispatch_id"`
	ItemDispatchDetails []ItemDispatch `json:"item_dispatch"`

	// Domestic Tax Info
	SGST string `json:"sgst"`
	CGST string `json:"cgst"`
	IGST NullString `json:"igst"`
	//domestic section
	PurchasePriceInr string `json:"purchasePriceInr"`
	MarketPriceInr   string `json:"marketPriceInr"`
	FinalPriceInr    string `json:"finalPriceInr"`
	DTerminalPrice   string `json:"terminalPrice"`
	TotalPrice       string `json:"totalPrice"`
	//Other Information
	TaxDuties        string `json:"taxes_duties"`
	ModeOfTransport  string `json:"mode_of_transport"`
	TransitInsurance string `json:"transit_insurance"`
	PackForward      string `json:"packing_forwarding"`
	//Other charges
	OtherCharges    string         `json:"otherCharges"`
	Rate            string         `json:"rate"`
	GrossPrice      string         `json:"grossPrice"`
	AuditLogDetails []AuditLogGCPO `json:"audit_log_gc_po"`
	//Consolidated Finance

	QCStatus      string `json:"qcStatus"`
	APStatus      string `json:"apStatus"`
	PayableAmount NullString `json:"payable_amount"`
	//new fields
	NoOfBags     string  `json:"no_of_bags"`
	NetWt        string  `json:"net_weight"`
	MTQuantity   float64 `json:"quantity_mt"`
	FixationDate string  `json:"fixation_date"`
	//Other Charges--Domestic
	DPackForward  string `json:"packing_forward_charges"`
	DInstallation string `json:"installation_charges"`
	DFreight      string `json:"freight_charges"`
	DHandling     string `json:"handling_charges"`
	DMisc         NullString `json:"misc_charges"`
	DHamali       NullString `json:"hamali_charges"`
	DMandiFee     string `json:"mandifee_charges"`
	DFullTax      string `json:"fulltax_charges"`
	DInsurance    string `json:"insurance_charges"`
}
type ItemDispatch struct {
	DispatchID        string `json:"dispatch_id"`
	DispatchQuantity  string `json:"dispatch_quantity"`
	DispatchDate      string `json:"dispatch_date"`
	DSNo              string `json:"number"`
	DDate             string `json:"date"`
	DeliveredQuantity NullString `json:"delivered_quantity"`
	BalanceQuantity   NullString `json:"balance_quantity"`
}
type AuditLogGCPO struct {
	CreatedDate    string `json:"createddate"`
	CreatedUserid  NullString `json:"createduserid"`
	ModifiedDate   NullString `json:"modifieddate"`
	ModifiedUserid NullString `json:"modifieduserid"`
	Description    NullString `json:"description"`
}
// type NullString sql.NullString
// func (ns *NullString) Scan(value interface{}) error {
// 	var s sql.NullString
// 	if err := s.Scan(value); err != nil {
// 		return err
// 	}

// 	// if nil then make Valid false
// 	if reflect.TypeOf(value) == nil {
// 		*ns = NullString{s.String, false}
// 	} else {
// 		*ns = NullString{s.String, true}
// 	}

// 	return nil
// }

type NullString string

func (s *NullString) Scan(value interface{}) error {
    if value == nil {
        *s = ""
        return nil
    }
    strVal, ok := value.(string)
    if !ok {
        return errors.New("Column is not a string")
    }
    *s = NullString(strVal)
    return nil
}
func (s NullString) Value() (driver.Value, error) {
    if len(s) == 0 { // if nil or empty string
        return nil, nil
    }
    return string(s), nil
}