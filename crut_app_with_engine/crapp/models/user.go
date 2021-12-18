package models

type User struct {
	
	UserId uint16	`json:"userId"`
	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	UserRole string `json:""userRole`
}

type PostUser struct {

	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	UserRole string `json:""userRole`
}

