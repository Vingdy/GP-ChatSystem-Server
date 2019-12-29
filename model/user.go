package model

type Register struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	NickName string `json:"nickname"`
}

type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Role     string `json:"role"`
}

type Login struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
