// package model

// type UserResquest struct {
// 	Email        string `json:"email"`
// 	FirstName    string `json:"first_name"`
// 	LastName     string `json:"last_name"`
// 	Entitlements string `json:"entitlements"`
// }

// type UserResponse struct {
// 	Status  int
// 	Message string
// }

package model

type UserResquest struct {
	Id  string `json:"id"`
	Id1 string `json:"id1"`
	Id2 int32  `json:"id3"`
	Id4 string `json:"id4"`
}

type UserResponse struct {
	Status  int
	Message string
}
