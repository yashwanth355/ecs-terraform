package model

type ListSupplier struct {
	Type string `json:"type"`
}
type ListSupplierDetails struct {
	Name          string `json:"vendor_name"`
	ContactPerson string `json:"contact_person"`
	Phone         string `json:"phone"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	Group         string `json:"group"`
	Suppliers     string `json:"suppliers"`
}
