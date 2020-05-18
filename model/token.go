package model

/*type Token   struct {
	UserName    string `json:"username"`
	AccessToken string `json:"accesstoken"`
	ExpiresAt   int64  `json:"expires_at"`
	RefreshToken string `json:"refreshtoken"`
	RefreshAt   int64  `json:"refresh_at"`
	CreateAt   int64  `json:"create_at"`
}*/

type Token struct {
	UserName    string `json:"username"`
	AccessToken string `json:"accesstoken"`
	ExpiresAt   int64  `json:"expires_at"`
	CreateAt    int64  `json:"create_at"`
}
