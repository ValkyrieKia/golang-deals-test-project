package dto

type SignInRequestDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshTokenRequestDTO struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type SignOutRequestDTO struct {
	SessionUid string `json:"session_uid" binding:"required"`
}
