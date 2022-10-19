package usecase

import (
	"errors"
	"token-repository/internal/model"

	"go.uber.org/zap"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrInvalidArgs    = errors.New("invalid args")
)

type OAuth2Repo interface {
	OAuth2Get(tokenName string) (oa model.OAuth2, err error)
	OAuth2Update(in model.OAuth2Update) (err error)
}
type TokenRepoService struct {
	Logger     *zap.Logger
	Repository OAuth2Repo
}

func (t *TokenRepoService) GetToken(tokenName string) (token model.OAuth2Get, err error) {
	tLogger := t.Logger.With(zap.String("token_name", tokenName))
	oa, err := t.Repository.OAuth2Get(tokenName)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			tLogger.Warn("token not found", zap.Error(err))
			return token, err
		}
		tLogger.Warn("token repository error", zap.Error(err))
		return token, err
	}
	token = oa.ConvDBOAuth2Get()
	return token, nil
}

func (t *TokenRepoService) UpdateToken(tokenUp model.OAuth2Update) (err error) {
	if tokenUp.TokenName == "" {
		return ErrInvalidArgs
	}

	tLogger := t.Logger.With(zap.String("token_name", tokenUp.TokenName))
	err = t.Repository.OAuth2Update(tokenUp)
	if err != nil {
		tLogger.Error("failed to update token information", zap.Error(err))
		return err
	}

	tLogger.Info("update token information")
	return nil
}
