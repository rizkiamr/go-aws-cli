package model

type GetCallerIdentityResponse struct {
	UserId  string `json:"UserId"`
	Account string `json:"Account"`
	Arn     string `json:"Arn"`
}
