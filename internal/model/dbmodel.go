package model

import "time"

type OAuth2 struct {
	TokenName    string `gorm:"primaryKey"`
	AccessToken  string
	RefreshToken string
	RefreshURL   string
	ExpiredAt    time.Time `gorm:"default:null"`
}

type OAuth2History struct {
	Id          int64 `gorm:"primaryKey"`
	TokenName   string
	ExpiredIn   int64
	RefreshURL  string
	RequestedAt time.Time
}

// 構造体とテーブル名を一致させる
func (l *OAuth2) TableName() string {
	return "oauth2"
}

func (l *OAuth2History) TableName() string {
	return "oauth2_history"
}

func (o *OAuth2) ConvDBOAuth2Get() (token OAuth2Get) {
	token.TokenName = o.TokenName
	token.AccessToken = o.AccessToken
	token.RefreshToken = o.RefreshToken
	token.ExpiredAt = o.ExpiredAt
	return token
}
