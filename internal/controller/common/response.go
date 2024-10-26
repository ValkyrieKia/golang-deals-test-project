package common

import (
	"github.com/ValkyrieKia/golang-deals-test-project/internal/util"
	"github.com/gin-gonic/gin"
)

type CommonResponse struct {
	Response interface{} `json:"response"`
}

func CreateCommonResponse(c *gin.Context, httpStatus int, data interface{}) {
	if data == nil {
		c.Status(httpStatus)
		return
	}
	c.JSON(httpStatus, CommonResponse{Response: data})
}

type CommonHTTPErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func CreateCommonHTTPErrorResponse(error *util.CommonError) *CommonHTTPErrorResponse {
	if error == nil {
		return &CommonHTTPErrorResponse{
			Error: struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}{
				Code:    "err/unknown",
				Message: "internal server error",
			},
		}
	}
	return &CommonHTTPErrorResponse{
		Error: struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}{
			Code:    error.Code,
			Message: error.Message,
		},
	}
}
