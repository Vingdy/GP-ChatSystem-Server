package model

type UserSendMessage struct {
	GroupID int//0是指全部
	UserName string
	Content string
}

type ServerSendMessage struct {
	GroupID int//0是指全部
	UserName string
	Content string
}
