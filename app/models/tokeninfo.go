package models

type TokenInfo struct {
	AccessToken string `json:"access_token" binding:"required"`
	ClientId    string `json:"client_id"`
	Uuid        string `json:"uuid"`
	Scope       string `json:"scope"`
	ExpiresIn   int64  `json:"expires_in"`
}
