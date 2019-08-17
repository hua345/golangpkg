package drawModel

type DrawResult struct {
	UserList    []User `json:"userList"`
	CompanyName string `json:"selectedCompanyName"`
}

func (drawResult *DrawResult) GetTypeList() map[string][]User {
	userTypeMap := make(map[string][]User)
	for _, item := range drawResult.UserList {
		userList, ok := userTypeMap[item.Type]
		if ok {
			userTypeMap[item.Type] = append(userList, item)
		} else {
			userTypeMap[item.Type] = append([]User{}, item)
		}
	}
	return userTypeMap
}
