package model

type Token   struct {
	username    string `json:"username"`
	AccessToken string `json:"accesstoken"`
	ExpiresAt   int64  `json:"expires_at"`
	Timestamp   int64  `json:"timestamp"`
}