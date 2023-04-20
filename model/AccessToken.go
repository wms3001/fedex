package model

type AccessToken struct {
	Access_Token string `json:"access_Token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	FedexError
}
