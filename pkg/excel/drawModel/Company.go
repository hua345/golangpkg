package drawModel

type Company struct {
	UserList    []User `json:"userList"`
	CompanyName string `json:"companyName"`
}
