package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/auth/jwt"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/infrastructure/cache"
	"github.com/uesleicarvalhoo/sorveteria-tres-estrelas/user"
)

const (
	accessToken  = "access-token"
	refreshToken = "refresh-token"
)

const (
	AccessTokenDuration  = time.Minute * 15 // Pegar das variavies de ambiente
	RefreshTokenDuration = time.Hour * 1    // Pegar das variavies de ambiente
)

func getCacheTokenKey(prefix string, id uuid.UUID) string {
	return fmt.Sprintf("%s-%s", prefix, id.String())
}

type Service struct {
	jwt    jwt.UseCase
	cache  cache.Cache
	userUc user.UseCase
}

func NewService(userUc user.UseCase, jwtUc jwt.UseCase, cache cache.Cache) *Service {
	return &Service{
		userUc: userUc,
		cache:  cache,
		jwt:    jwtUc,
	}
}

func (s *Service) generateToken(ctx context.Context, u user.User, prefix string, exp time.Time) (string, error) {
	token, err := s.jwt.Generate(ctx, u, exp)
	if err != nil {
		return "", err
	}

	key := getCacheTokenKey(prefix, u.ID)

	var duration time.Duration
	if prefix == accessToken {
		duration = AccessTokenDuration
	} else {
		duration = RefreshTokenDuration
	}

	if err := s.cache.Set(ctx, key, token, duration); err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) createAccessToken(ctx context.Context, u user.User) (JwtToken, error) {
	now := time.Now()
	accessExp := now.Add(AccessTokenDuration)
	refreshExp := now.Add(RefreshTokenDuration)

	accessToken, err := s.generateToken(ctx, u, accessToken, accessExp)
	if err != nil {
		return JwtToken{}, err
	}

	refreshToken, err := s.generateToken(ctx, u, refreshToken, refreshExp)
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

func (s *Service) validateToken(ctx context.Context, prefix, token string) (user.User, error) {
	u, err := s.jwt.Validate(ctx, token)
	if err != nil {
		return user.User{}, err
	}

	cachedToken, err := s.cache.Get(ctx, getCacheTokenKey(prefix, u.ID))
	if err != nil {
		return user.User{}, err
	}

	if cachedToken != token {
		return user.User{}, ErrTokenNotFound
	}

	return u, nil
}

func (s *Service) Login(ctx context.Context, email, password string) (JwtToken, error) {
	u, err := s.userUc.GetByEmail(ctx, email)
	if err != nil {
		return JwtToken{}, err
	}

	if !u.CheckPassword(password) {
		return JwtToken{}, ErrNotAuthorized
	}

	return s.createAccessToken(ctx, u)
}

func (s *Service) RefreshToken(ctx context.Context, token string) (JwtToken, error) {
	id, err := s.validateToken(ctx, refreshToken, token)
	if err != nil {
		return JwtToken{}, err
	}

	return s.createAccessToken(ctx, id)
}

func (s *Service) Authorize(ctx context.Context, token string) (user.User, error) {
	u, err := s.validateToken(ctx, accessToken, token)
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}
