package auth

import (
	"context"
)

type StaticConfigProvider struct {
	secretKey string
	issuer    string
}

var _ ConfigProvider = StaticConfigProvider{}

func NewStaticProvider(secret, issuer string) StaticConfigProvider {
	return StaticConfigProvider{
		secretKey: secret,
		issuer:    issuer,
	}
}

func (s StaticConfigProvider) GetSecretKey(_ context.Context) (string, error) {
	return s.secretKey, nil
}

func (s StaticConfigProvider) GetIssuer(_ context.Context) (string, error) {
	return s.issuer, nil
}
