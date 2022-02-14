package entity

import (
	"time"
)

type CmsLeadsMaster struct {
	Accountname         string    `gorm:"column:accountname"`
	Accountid           int64     `gorm:"column:accountid"`
	Aliases             string    `gorm:"column:aliases"`
	Accounttypeid       string    `gorm:"column:accounttypeid"`
	Website             string    `gorm:"column:website"`
	Leadid              string    `gorm:"column:leadid"`
	Approvalstatus      int64     `gorm:"column:approvalstatus"`
	Approverid          int64     `gorm:"column:approverid"`
	Approxannualrev     string    `gorm:"column:approxannualrev"`
	Coffeetypeid        string    `gorm:"column:coffeetypeid"`
	Comments            string    `gorm:"column:comments"`
	ContactExt          string    `gorm:"column:contact_ext_id"`
	ContactMobile       string    `gorm:"column:contact_mobile"`
	ContactPosition     string    `gorm:"column:contact_position"`
	ContactSalutationid int64     `gorm:"column:contact_salutationid"`
	Contactfirstname    string    `gorm:"column:contactfirstname"`
	Contactlastname     string    `gorm:"column:contactlastname"`
	Countryid           int64     `gorm:"column:countryid"`
	Createddate         time.Time `gorm:"column:createddate"`
	Createduserid       string    `gorm:"column:createduserid"`
	Email               string    `gorm:"column:email"`
	Instcoffee          int       `gorm:"column:instcoffee"`
	Isactive            int       `gorm:"column:isactive"`
	Manfacunit          int       `gorm:"column:manfacunit"`
	Leadscore           int64     `gorm:"column:leadscore"`
	Masterstatus        string    `gorm:"column:masterstatus"`
	Modifieddate        time.Time `gorm:"column:modifieddate"`
	Modifieduserid      int64     `gorm:"column:modifieduserid"`
	Otherinformation    string    `gorm:"column:otherinformation"`
	Phone               string    `gorm:"column:phone"`
	Price               int64     `gorm:"column:price"`
	Productsegmentid    string    `gorm:"column:productsegmentid"`
	Recordtypeid        int64     `gorm:"column:recordtypeid"`
	ShippingContinent   string    `gorm:"column:shipping_continent"`
	ShippingContinentid int64     `gorm:"column:shipping_continentid"`
	ShippingCountry     string    `gorm:"column:shipping_country"`
}

func (cms CmsLeadsMaster) TableName() string {
	return "cms_leads_master"
}
