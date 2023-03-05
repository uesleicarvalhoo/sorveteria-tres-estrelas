package auth

import (
	"context"

	"github.com/kong/go-kong/kong"
)

type KongProvider struct {
	jwtKey   string
	username string
	cli      *kong.Client
}

var _ ConfigProvider = KongProvider{}

func NewKongProvider(cli *kong.Client, username, jwtKey string) KongProvider {
	return KongProvider{
		jwtKey:   jwtKey,
		username: username,
		cli:      cli,
	}
}

func (k KongProvider) getConsumer(ctx context.Context, username string) (*kong.Consumer, error) {
	return k.cli.Consumers.Get(ctx, &username)
}

func (k KongProvider) getAuthJwt(ctx context.Context) (*kong.JWTAuth, error) {
	consumer, err := k.getConsumer(ctx, k.username)
	if err != nil {
		return nil, err
	}

	auth, err := k.cli.JWTAuths.Get(ctx, consumer.ID, &k.jwtKey)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (k KongProvider) GetSecretKey(ctx context.Context) (string, error) {
	auth, err := k.getAuthJwt(ctx)
	if err != nil {
		return "", err
	}

	return *auth.Secret, nil
}

func (k KongProvider) GetIssuer(ctx context.Context) (string, error) {
	auth, err := k.getAuthJwt(ctx)
	if err != nil {
		return "", err
	}

	return *auth.Key, nil
}
