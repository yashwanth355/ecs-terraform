package entity

type UserDetails struct {
	Active            bool   `json:"active"`
	Alias             string `json:"alias"`
	City              string `json:"city"`
	Company           string `json:"company"`
	Country           string `json:"country"`
	Delegatedapprover string `json:"delegatedapprover"`
	Department        string `json:"department"`
	Designation       string `json:"designation"`
	Division          string `json:"division"`
	Emailid           string `json:"emailid"`
	Empcode           string  `json:"empcode"`
	Employee          bool   `json:"employee"`
	Ext               string `json:"ext"`
	Firstname         string `json:"firstname"`
	Lastname          string `json:"lastname"`
	Manager           string `json:"manager"`
	Middlename        string `json:"middlename"`
	Mobile            string  `json:"mobile"`
	Password          string `json:"password"`
	Phone             string  `json:"phone"`
	Postalcode        string  `json:"postalcode"`
	Profile           string `json:"profile"`
	Role              string `json:"role"`
	State             string `json:"state"`
	Street            string `json:"street"`
	Title             string `json:"title"`
	Userid            string  `json:"userid"`
	Username          string `json:"username"`
}
