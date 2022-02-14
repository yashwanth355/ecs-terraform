package model

type UserDetails struct {
	Userid                   int    `json:"userid"`
	Firstname                string `json:"firstname"`
	Middlename               string `json:"middlename"`
	Lastname                 string `json:"lastname"`
	Emailid                  string `json:"emailid"`
	Alias                    string `json:"alias"`
	Username                 string `json:"username"`
	Empcode                  int    `json:"empcode"`
	Designation              string `json:"designation"`
	Company                  string `json:"company"`
	Department               string `json:"department"`
	Role                     string `json:"role"`
	Division                 string `json:"division"`
	Employee                 bool   `json:"employee"`
	Profile                  string `json:"profile"`
	Delegatedapprover        string `json:"delegatedapprover"`
	Manager                  string `json:"manager"`
	Receiveapprovalreqemails string `json:"receiveapprovalreqemails"`
}

type UserDetailsResponse struct {
	Status  int
	Payload []*UserDetails
}

type MyJsonName struct {
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
	Empcode           int64  `json:"empcode"`
	Employee          bool   `json:"employee"`
	Ext               string `json:"ext"`
	Firstname         string `json:"firstname"`
	Lastname          string `json:"lastname"`
	Manager           string `json:"manager"`
	Middlename        string `json:"middlename"`
	Mobile            int64  `json:"mobile"`
	Password          string `json:"password"`
	Phone             int64  `json:"phone"`
	Postalcode        int64  `json:"postalcode"`
	Profile           string `json:"profile"`
	Role              string `json:"role"`
	State             string `json:"state"`
	Street            string `json:"street"`
	Title             string `json:"title"`
	Userid            int64  `json:"userid"`
	Username          string `json:"username"`
}
