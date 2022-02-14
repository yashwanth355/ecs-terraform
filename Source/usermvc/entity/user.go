package entity

// type User struct {
// 	Email        string `json:"email"`
// 	FirstName    string `json:"firstName"`
// 	LastName     string `json:"lastName"`
// 	Entitlements string `json:"entitlements,omitempty"`
// }

type User struct {
	Id  string `gorm:"column:id"`
	Id1 string `gorm:"column:id1"`
	Id2 int32  `gorm:"column:id3"`
	Id3 string `gorm:"column:id4"`
}