package model

type Register struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	NickName string `json:"nickname"`
}

type User struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Role     string `json:"role"`
	Phone 	 string `json:"phone"`
	Label 	 string `json:"label"`
	Head 	 string `json:"head"`
	IsBan 	 string `json:"isban"`
	Token    string `json:"token,omitempty"`
}

type Login struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
