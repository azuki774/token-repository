package model

import "time"

type OAuth2Update struct {
	TokenName    string    `json:"token_name"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiredIn    int64     `json:"expired_in"`
	RefreshURL   string    `json:"refresh_url"`
	RequestAt    time.Time `json:"requested_at"`
}

type OAuth2Get struct {
	TokenName    string    `json:"token_name"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiredAt    time.Time `json:"expired_at"`
}

func (o *OAuth2Update) ConvDBOAuth2() OAuth2 {
	ret := OAuth2{
		TokenName:    o.TokenName,
		AccessToken:  o.AccessToken,
		RefreshToken: o.RefreshToken,
		RefreshURL:   o.RefreshURL,
		ExpiredAt:    o.RequestAt.Add(time.Second * time.Duration(o.ExpiredIn)),
	}
	return ret
}

func (o *OAuth2Update) ConvDBOAuth2History() OAuth2History {
	ret := OAuth2History{
		TokenName:   o.TokenName,
		ExpiredIn:   o.ExpiredIn,
		RefreshURL:  o.RefreshURL,
		RequestedAt: o.RequestAt,
	}
	return ret
}
