package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ValkyrieKia/golang-deals-test-project/internal/app/entity"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/config"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/controller/common"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/util"
	"github.com/gin-gonic/gin"
)

func UseAuthentication() gin.HandlerFunc {
	cfg := config.GetConfig().AuthConfig

	handleUnauthorized := func(c *gin.Context) {
		c.JSON(http.StatusUnauthorized, common.CreateCommonHTTPErrorResponse(
			util.NewCommonError(nil, util.ErrUnauthorized, "invalid authorization header")),
		)
		c.Abort()
	}

	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			handleUnauthorized(c)
			return
		}

		splitToken := strings.Split(authHeader, " ")
		if len(splitToken) < 2 {
			handleUnauthorized(c)
			return
		}
		bearerToken := strings.TrimSpace(splitToken[1])

		validation, err := util.ValidateJwt(bearerToken, cfg.JwtTokenSecret)
		if err != nil {
			handleUnauthorized(c)
			return
		}

		validation = map[string]interface{}(validation)
		validationString, _ := json.Marshal(validation)

		var authData *entity.AuthTokenData
		_ = json.Unmarshal(validationString, &authData)

		c.Set("auth", authData)
		c.Next()
	}
}
