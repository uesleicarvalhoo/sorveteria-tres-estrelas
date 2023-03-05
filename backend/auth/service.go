package auth

import (
	"context"
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/backend/user"
)

const (
	AccessTokenDuration = time.Minute * 15 // Pegar das variavies de ambiente
)

type Service struct {
	provider ConfigProvider
	userUc   user.UseCase
}

func NewService(userUc user.UseCase, provider ConfigProvider) *Service {
	return &Service{
		userUc:   userUc,
		provider: provider,
	}
}

func (s *Service) generateToken(ctx context.Context, u user.User, exp time.Time) (string, error) {
	issuer, err := s.provider.GetIssuer(ctx)
	if err != nil {
		return "", err
	}

	secret, err := s.provider.GetSecretKey(ctx)
	if err != nil {
		return "", err
	}

	return GenerateJwtToken(ctx, u, exp, issuer, secret)
}

func (s *Service) createAccessToken(ctx context.Context, u user.User) (JwtToken, error) {
	exp := time.Now().Add(AccessTokenDuration)

	token, err := s.generateToken(ctx, u, exp)
	if err != nil {
		return JwtToken{}, err
	}

	return JwtToken{
		GrantType: "bearer",
		Token:     token,
		ExpiresAt: exp.Unix(),
	}, nil
}

func (s *Service) validateToken(ctx context.Context, token string) (user.User, error) {
	secret, err := s.provider.GetSecretKey(ctx)
	if err != nil {
		return user.User{}, err
	}

	return ValidateJwtToken(ctx, token, secret)
}

func (s *Service) Login(ctx context.Context, payload LoginPayload) (JwtToken, error) {
	u, err := s.userUc.GetByEmail(ctx, payload.Email)
	if err != nil {
		return JwtToken{}, err
	}

	if !u.CheckPassword(payload.Password) {
		return JwtToken{}, ErrNotAuthorized
	}

	return s.createAccessToken(ctx, u)
}

func (s *Service) RefreshToken(ctx context.Context, payload RefreshTokenPayload) (JwtToken, error) {
	id, err := s.validateToken(ctx, payload.RefreshToken)
	if err != nil {
		return JwtToken{}, err
	}

	return s.createAccessToken(ctx, id)
}

func (s *Service) Authorize(ctx context.Context, token string) (user.User, error) {
	u, err := s.validateToken(ctx, token)
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}
