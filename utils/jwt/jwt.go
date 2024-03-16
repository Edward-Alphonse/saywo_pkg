package jwt

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
)

//var JwtSecret = "23347$040412"
//
//var jwtSecret = []byte(JwtSecret)

type TokenType string

const (
	TokenTypeAccessToken  TokenType = "access_token"
	TokenTypeRefreshToken TokenType = "refresh_token"
)

type Claims struct {
	UserId    string    `json:"user_id"`
	TokenType TokenType `json:"token_type"`
	jwt.StandardClaims
}

// 生成JWT
func GenerateToken(userId uint64, tokenType TokenType, expireAt int64) (string, error) {
	claims := Claims{
		fmt.Sprintf("%d", userId),
		tokenType,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    "gin-saywo",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secrect := getJwtSecret(userId, tokenType)
	token, err := tokenClaims.SignedString([]byte(secrect))
	return token, err
}

// ParseToken 解析JWT， 使用userId和tokenType生成secrect，当外部传入的userId或者tokenType变化时，则jwt解析失败
func ParseToken(token string, userId uint64, tokenType TokenType) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (any, error) {
		secrect := getJwtSecret(userId, tokenType)
		return []byte(secrect), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims == nil {
		return nil, errors.New("empty claim")
	}
	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	}
	return nil, errors.New("empty claim")
}

func getJwtSecret(userId uint64, tokenType TokenType) string {
	return fmt.Sprintf("say_wo:%d_%v", userId, tokenType)
}
