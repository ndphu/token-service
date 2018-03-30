package domain

type Token struct {
	Exp        int64    `json:"exp"`
	Iss        string   `json:"iss"`
	Sub        string   `json:"sub"`
	Id         string   `json:"id"`
	ExternalId string   `json:"externalId"`
	FirstName  string   `json:"firstName"`
	LastName   string   `json:"lastName"`
	Roles      []string `json:"roles"`
	AppId      string   `json:"appId"`
}
