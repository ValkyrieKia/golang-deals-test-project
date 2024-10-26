package entity

type AuthSignInData struct {
	Username   string
	Password   string
	ClientIP   string
	DeviceInfo string
}

type AuthSignInResult struct {
	UserSession *UserSession
	AccessToken string
}

type AuthTokenData struct {
	ExpiresAt  int64  `json:"exp"`
	IssuedAt   int64  `json:"iat"`
	SessionUid string `json:"id"`
	TokenUid   string `json:"jti"`
	Type       string `json:"type"`
}
