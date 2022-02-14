package model

type Countries struct {
	CountryId string `json:"countryid"`
	CountryName   string `json:"country"`
}
type States struct {
	StateName string `json:"state"`
}
type Cities struct {
	CityName string `json:"city"`
}
// type Countries struct {
// 	Countryname string `json:"countryname"`
// }

type Continents struct {
	ContinentName string `json:"continent_name"`
	ContinentCode string `json:"continent_code"`
}
type PhoneCodes struct {
	Id          int    `json:"id"`
	Countryname string `json:"countryname"`
	Dialcode    int64  `json:"dialcode"`
}
// type CityResponse struct {
// 	status  int
// 	payload []*GetCity
// }
// type GetState struct {
// 	State string `json:"state"`
// }
// type StateResponse struct {
// 	status  int
// 	payload []*GetState
// }
