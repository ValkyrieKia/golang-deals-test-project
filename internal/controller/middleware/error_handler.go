package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/ValkyrieKia/golang-deals-test-project/internal/config"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/controller/common"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/util"
	"github.com/gin-gonic/gin"
)

func UseErrorHandler(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			// we would want to get only the first error occurence.
			handlerError := c.Errors[0]
			var cmErr *util.CommonError
			if errors.As(handlerError.Unwrap(), &cmErr) {
				errPayload := common.CreateCommonHTTPErrorResponse(cmErr)
				c.JSON(cmErr.HTTPStatus, errPayload)
				return
			}
			log.Printf("unexpected error: %s", handlerError.Error())
			c.JSON(http.StatusInternalServerError, common.CreateCommonHTTPErrorResponse(nil))
		}
	}
}
