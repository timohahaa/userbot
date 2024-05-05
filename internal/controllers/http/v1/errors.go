package v1

import "github.com/labstack/echo/v4"

func newErrorMessage(c echo.Context, statusCode int, message string) error {
	httpErr := echo.NewHTTPError(statusCode, message)

	return c.JSON(statusCode, httpErr)
}
