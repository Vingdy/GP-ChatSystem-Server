package model

type Register struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	NickName string `json:"nickname"`
}

type User struct {
	Id        string `json:"id"`
	UserName  string `json:"username"`
	NickName  string `json:"nickname"`
	Role      string `json:"role"`
	Phone     string `json:"phone"`
	Label     string `json:"label"`
	FontType  string `json:"fonttype"`
	FontColor string `json:"fontcolor"`
	IsBan     string `json:"isban"`
	Token     string `json:"token,omitempty"`
}

type Login struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
