package entity

type CmsPhonecodesMaster struct {
	Countryname string `db:"Country_Name"`
	Dial        int64  `db:"Dial"`
	Id          int64  `db:"id"`
}
