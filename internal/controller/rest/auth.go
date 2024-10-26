package rest

import (
	"net/http"

	"github.com/ValkyrieKia/golang-deals-test-project/internal/app/entity"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/app/usecase"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/config"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/controller/common"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/controller/dto"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/infrastructure/repository/user"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/infrastructure/repository/user_session"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/util"
	"github.com/gin-gonic/gin"
)

type authController struct {
	rg     *gin.RouterGroup
	config *config.Config

	authUc *usecase.AuthUsecase
}

func NewAuthHttpController(
	rg *gin.RouterGroup,
	providers *common.HttpServerProviders,
) *authController {
	// repository and dependency initializations
	requiredUserRepository := user.NewUserMysqlRepository(providers.MainDbConnection)
	requiredUserSessionRepository := user_session.NewUserSessionMysqlRepository(providers.MainDbConnection)

	ctrl := &authController{
		rg:     rg,
		config: providers.Config,
		authUc: usecase.NewAuthUsecase(
			providers.Config,
			requiredUserRepository,
			requiredUserSessionRepository,
		),
	}

	// route initializations

	rg.POST("/sign-in", ctrl.signIn)
	rg.POST("/refresh", ctrl.refresh)
	rg.POST("/sign-out", ctrl.signOut)

	return ctrl
}

func (ac *authController) signIn(c *gin.Context) {
	var body *dto.SignInRequestDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		_ = c.Error(util.NewCommonError(err, util.ErrBadRequest, err.Error()))
		return
	}

	deviceInfo := c.Request.Header.Get("User-Agent")

	authResult, err := ac.authUc.SignIn(&entity.AuthSignInData{
		Username:   body.Username,
		Password:   body.Password,
		ClientIP:   c.ClientIP(),
		DeviceInfo: deviceInfo,
	})
	if err != nil {
		_ = c.Error(err)
		return
	}

	common.CreateCommonResponse(c, http.StatusOK, &gin.H{
		"access_token":  authResult.AccessToken,
		"refresh_token": authResult.UserSession.RefreshToken,
	})
	return
}

func (ac *authController) refresh(c *gin.Context) {
	var body *dto.RefreshTokenRequestDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		_ = c.Error(util.NewCommonError(err, util.ErrBadRequest, err.Error()))
		return
	}

	newToken, err := ac.authUc.Refresh(body.RefreshToken)
	if err != nil {
		_ = c.Error(err)
		return
	}

	common.CreateCommonResponse(c, http.StatusOK, &gin.H{
		"access_token": newToken,
	})
	return
}

func (ac *authController) signOut(c *gin.Context) {
	var body *dto.SignOutRequestDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		_ = c.Error(util.NewCommonError(err, util.ErrBadRequest, err.Error()))
		return
	}

	err := ac.authUc.SignOut(body.SessionUid)
	if err != nil {
		_ = c.Error(err)
		return
	}

	common.CreateCommonResponse(c, http.StatusOK, &gin.H{
		"success": true,
	})
	return
}
