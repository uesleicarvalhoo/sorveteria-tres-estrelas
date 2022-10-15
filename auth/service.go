package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/entity"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/pkg/password"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/usecase/user"
)

const (
	accessToken  = "access-token"
	refreshToken = "refresh-token"
)

const (
	AccessTokenDuration  = time.Minute * 15 // Pegar das variavies de ambiente
	RefreshTokenDuration = time.Hour * 1    // Pegar das variavies de ambiente
)

func getCacheTokenKey(prefix string, id entity.ID) string {
	return fmt.Sprintf("%s-%s", prefix, id.String())
}

func GetDefaultUserPermissions() []entity.Permission {
	return []entity.Permission{entity.ReadWritePopsicle, entity.ReadWriteSalesRole}
}

type Service struct {
	secret    string
	cache     Cache
	userUc    user.UseCase
	passwdSvc password.Service
}

func NewService(secret string, userUc user.UseCase, cache Cache, passwdSvc password.Service) *Service {
	return &Service{
		secret:    secret,
		userUc:    userUc,
		cache:     cache,
		passwdSvc: passwdSvc,
	}
}

func (s *Service) generateToken(ctx context.Context, id entity.ID, prefix string, exp time.Time) (string, error) {
	token, err := GenerateJwtToken(s.secret, id, exp)
	if err != nil {
		return "", err
	}

	key := getCacheTokenKey(prefix, id)

	if err := s.cache.Set(ctx, key, token); err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) createAccessToken(ctx context.Context, id entity.ID) (JwtToken, error) {
	now := time.Now()
	accessExp := now.Add(AccessTokenDuration)
	refreshExp := now.Add(RefreshTokenDuration)

	accessToken, err := s.generateToken(ctx, id, accessToken, accessExp)
	if err != nil {
		return JwtToken{}, err
	}

	refreshToken, err := s.generateToken(ctx, id, refreshToken, refreshExp)
	if err != nil {
		return JwtToken{}, err
	}

	return JwtToken{
		GrantType:    "bearer",
		AcessToken:   accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    accessExp.Unix(),
	}, nil
}

func (s *Service) validateToken(ctx context.Context, prefix, token string) (entity.ID, error) {
	id, err := ValidateJwtToken(token, s.secret)
	if err != nil {
		return entity.ID{}, err
	}

	cachedToken, err := s.cache.Get(ctx, getCacheTokenKey(prefix, id))
	if err != nil {
		return entity.ID{}, err
	}

	if cachedToken != token {
		return entity.ID{}, ErrTokenNotFound
	}

	return id, nil
}

func (s *Service) Login(ctx context.Context, email, password string) (JwtToken, error) {
	found, err := s.userUc.GetByEmail(ctx, email)
	if err != nil {
		return JwtToken{}, err
	}

	if !s.passwdSvc.CheckHash(password, found.PasswordHash) {
		return JwtToken{}, ErrNotAuthorized
	}

	return s.createAccessToken(ctx, found.ID)
}

func (s *Service) RefreshToken(ctx context.Context, token string) (JwtToken, error) {
	id, err := s.validateToken(ctx, refreshToken, token)
	if err != nil {
		return JwtToken{}, err
	}

	return s.createAccessToken(ctx, id)
}

func (s *Service) Authorize(ctx context.Context, token string) (entity.ID, error) {
	return s.validateToken(ctx, accessToken, token)
}
