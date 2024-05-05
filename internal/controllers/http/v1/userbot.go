package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/timohahaa/userbot/internal/service"
)

type userbotRoutes struct {
	telegramService service.TelegramService
}

func newUserbotRoutes(g *echo.Group, telegramService service.TelegramService) {
	_ = &userbotRoutes{
		telegramService: telegramService,
	}
}




func SaveChannelByName(c echo.Context) error {

	return nil
}
