package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SuccessResponseWithoutData struct {
	Message string `json:"message"`
}

func SendStatusForbiddenResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusForbidden, ErrorResponse{
		Message: message,
	})
}

func SendStatusCreatedResponse(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusCreated, SuccessResponse{
		Message: message,
		Data:    data,
	})
}

func SendStatusOkResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, SuccessResponseWithoutData{
		Message: message,
	})
}

func SendStatusOkWithDataResponse(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, SuccessResponse{
		Message: message,
		Data:    data,
	})
}

func SendStatusInternalServerResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusInternalServerError, ErrorResponse{
		Message: message,
	})
}

func SendBadRequestResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusBadRequest, ErrorResponse{
		Message: message,
	})
}

func SendSuccessResponse(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, SuccessResponse{
		Message: message,
		Data:    data,
	})
}

func SendStatusConflictResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusConflict, ErrorResponse{
		Message: message,
	})
}

func SendStatusNotFoundResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusNotFound, ErrorResponse{
		Message: message,
	})
}

func SendStatusUnauthorizedResponse(c echo.Context, message string) error {
	return c.JSON(http.StatusUnauthorized, ErrorResponse{
		Message: message,
	})
}
