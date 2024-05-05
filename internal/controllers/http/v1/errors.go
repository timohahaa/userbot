package v1

import (
	"errors"
	"github.com/labstack/echo/v4"
)

var (
	ErrInternalServer       = errors.New("internal server error")
	ErrInvalidRequestBody   = errors.New("invalid request body")
	ErrInvalidPathParameter = errors.New("invalid path parameter")

	ErrChannelNotFound     = errors.New("channel not found")
	ErrFailedToSaveChannel = errors.New("failed to save channel")
)

func newErrorMessage(c echo.Context, statusCode int, message string) error {
	httpErr := echo.NewHTTPError(statusCode, message)

	return c.JSON(statusCode, httpErr)
}
