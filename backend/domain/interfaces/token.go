package interfaces

import "backend-starter-project/domain/entities"


type RefreshTokenRepository interface{
	CreateRefreshToken(refreshToken *entities.RefreshToken) (*entities.RefreshToken, error)
	FindRefreshTokenByUserId(userId string) (*entities.RefreshToken, error)
	DeleteRefreshTokenByUserId(userId string) error
}

type TokenService interface {
	VerifyAccessToken(token string) error
	GenerateAccessToken(*entities.User) (string, error)
	GenerateRefreshToken(*entities.User) (*entities.RefreshToken, error)
	VerifyRefreshToken(token string) error
	InvalidateAccessToken(token string) (string, error)
	InvalidateRefreshToken(token string) (string, error)
	GetClaimsFromToken(token string) map[string]string
	RefreshAccessToken(token *entities.RefreshToken) (string,error)
	CreateRefreshToken(refreshToken *entities.RefreshToken) (*entities.RefreshToken, error)
	DeleteRefreshTokenByUserId(userId string) error
	FindRefreshTokenByUserId(userId string) (*entities.RefreshToken, error)
}
