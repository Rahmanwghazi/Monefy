package presenter

import (
	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `jaon:"message"`
	}
	Data interface{} `json:"data"`
}

func NewSuccessResponse(echoContext echo.Context, status int, data interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = "success"
	response.Data = data
	return echoContext.JSON(status, response)
}

func NewErrorResponse(echoContext echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = err.Error()
	response.Data = nil
	return echoContext.JSON(status, response)
}
