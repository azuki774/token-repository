package repository

import (
	"errors"
	"token-repository/internal/model"
	"token-repository/internal/usecase"

	"gorm.io/gorm"
)

type OAuth2Repo struct {
	Conn *gorm.DB
}

func (o *OAuth2Repo) OAuth2Get(tokenName string) (oa model.OAuth2, err error) {
	err = o.Conn.Where("token_name = ?", tokenName).Take(&oa).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return oa, usecase.ErrRecordNotFound
		}
		// Internal error
		return oa, err
	}
	return oa, nil
}

func (o *OAuth2Repo) OAuth2Update(in model.OAuth2Update) (err error) {
	err = o.Conn.Transaction(func(tx *gorm.DB) error {
		// Save to OAuth2
		oa2 := in.ConvDBOAuth2()
		if err := tx.Save(&oa2).Error; err != nil {
			return err
		}

		// Save to OAuth2_history
		oa2his := in.ConvDBOAuth2History()
		if err := tx.Save(&oa2his).Error; err != nil {
			return err
		}

		// commit
		return nil
	})

	return err
}
