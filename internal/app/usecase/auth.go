package usecase

import (
	"time"

	"github.com/ValkyrieKia/golang-deals-test-project/internal/app/entity"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/app/repository"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/config"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/util"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	config                *config.Config
	userRepository        repository.UserRepository
	userSessionRepository repository.UserSessionRepository
}

func NewAuthUsecase(
	config *config.Config,
	userRepository repository.UserRepository,
	userSessionRepository repository.UserSessionRepository,
) *AuthUsecase {
	return &AuthUsecase{
		config:                config,
		userRepository:        userRepository,
		userSessionRepository: userSessionRepository,
	}
}

func (au *AuthUsecase) SignIn(data *entity.AuthSignInData) (*entity.AuthSignInResult, error) {
	user, err := au.userRepository.GetByUsername(data.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, util.NewCommonError(nil, util.ErrUnauthorized, "user not found")
	}

	passError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if passError != nil {
		return nil, util.NewCommonError(passError, util.ErrUnauthorized, "invalid password")
	}

	// create user session
	sessionID := uuid.New()
	tokenID := uuid.New()
	accessToken, err := util.GenerateAccessTokenJwt(
		au.config,
		tokenID.String(),
		sessionID.String(),
		nil,
	)
	refreshToken, err := util.GenerateRefreshTokenJwt(
		au.config,
		sessionID.String(),
		nil,
	)

	now := time.Now()
	refreshDuration, parseErr := time.ParseDuration("")
	refreshExpiry := time.Now().Add(refreshDuration)
	if parseErr != nil {
		refreshExpiry = time.Now().Add(30 * 24 * time.Hour)
	}
	sess := &entity.UserSession{
		SessionUID:   sessionID.String(),
		User:         user,
		RefreshToken: refreshToken,
		IPAddress:    data.ClientIP,
		ExpiresAt:    refreshExpiry,
		CreatedAt:    now,
	}
	if data.DeviceInfo != "" {
		sess.DeviceInfo = &data.DeviceInfo
	}

	sess, err = au.userSessionRepository.Create(sess)
	if err != nil {
		return nil, util.NewCommonError(err, util.ErrInternal, err.Error())
	}

	return &entity.AuthSignInResult{
		AccessToken: accessToken,
		UserSession: sess,
	}, nil
}

func (au *AuthUsecase) Refresh(refreshToken string) (string, error) {
	validation, err := util.ValidateJwt(refreshToken, au.config.AuthConfig.JwtRefreshTokenSecret)
	if err != nil {
		return "", util.NewCommonError(err, util.ErrUnauthorized, "invalid refresh token")
	}

	sessId, _ := validation["id"].(string)
	session, sessionErr := au.userSessionRepository.GetByUid(sessId)
	if sessionErr != nil {
		return "", util.NewCommonError(sessionErr, util.ErrInternal, sessionErr.Error())
	}

	if session == nil {
		return "", util.NewCommonError(nil, util.ErrUnauthorized, "session not found")
	}

	genTokenId := util.GenerateRandomTokenString(16)
	newToken, err := util.GenerateAccessTokenJwt(au.config, genTokenId, sessId, nil)
	if err != nil {
		return "", util.NewCommonError(err, util.ErrInternal, "failed to generate access token")
	}

	//_, updateErr := au.userSessionRepository.Update(sessId, session)
	//if updateErr != nil {
	//	return "", util.NewCommonError(err, util.ErrInternal, "failed to update session")
	//}

	return newToken, nil
}

func (au *AuthUsecase) SignOut(sessionUid string) error {
	err := au.userSessionRepository.Destroy(sessionUid)
	if err != nil {
		return util.NewCommonError(err, util.ErrInternal, err.Error())
	}
	return nil
}
