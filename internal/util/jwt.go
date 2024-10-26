package util

import (
	"errors"
	"log"
	"time"

	"github.com/ValkyrieKia/golang-deals-test-project/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

const (
	JwtAccessToken  = "access_token"
	JwtRefreshToken = "refresh_token"
)

func GenerateAccessTokenJwt(
	cfg *config.Config,
	tokenID string,
	sessionID string,
	additionalClaims *jwt.MapClaims,
) (string, error) {
	accessDuration, parseErr := time.ParseDuration(cfg.AuthConfig.JwtTokenExpiry)
	accessExpiry := time.Now().Add(accessDuration)
	if parseErr != nil {
		log.Printf("WARNING: Failed to parse token expiry: %s. Using default values.", parseErr.Error())
		accessExpiry = time.Now().Add(60 * time.Minute)
	}

	claims := jwt.MapClaims{
		"jti":  tokenID,
		"exp":  jwt.NewNumericDate(accessExpiry),
		"iat":  jwt.NewNumericDate(time.Now()),
		"id":   sessionID,
		"type": JwtAccessToken,
	}
	if additionalClaims != nil {
		for k, v := range *additionalClaims {
			claims[k] = v
		}
	}

	return generateJwt(cfg.AuthConfig.JwtTokenSecret, &claims)
}

func GenerateRefreshTokenJwt(cfg *config.Config, sessionID string, additionalClaims *jwt.MapClaims) (string, error) {
	duration, parseErr := time.ParseDuration(cfg.AuthConfig.JwtRefreshTokenExpiry)
	expiry := time.Now().Add(duration)
	if parseErr != nil {
		log.Printf("WARNING: Failed to parse refresh token expiry: %s. Using default values.", parseErr.Error())
		expiry = time.Now().Add(24 * 30 * time.Hour)
	}

	claims := jwt.MapClaims{
		"exp":  jwt.NewNumericDate(expiry),
		"iat":  jwt.NewNumericDate(time.Now()),
		"id":   sessionID,
		"type": JwtRefreshToken,
	}
	if additionalClaims != nil {
		for k, v := range *additionalClaims {
			claims[k] = v
		}
	}

	return generateJwt(cfg.AuthConfig.JwtRefreshTokenSecret, &claims)
}

func ValidateJwt(tokenString string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//return nil, NewCommonError(nil, ErrInternal, fmt.Sprintf("unexpected signing method: %v\n", token.Header["alg"]))
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	var claims jwt.MapClaims
	var tokenError error

	if token != nil {
		claims = token.Claims.(jwt.MapClaims)
		if !token.Valid {
			tokenError = err

		}
	}
	if err != nil && tokenError == nil {
		tokenError = err
	}

	return claims, tokenError
}

func generateJwt(secret string, claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
