package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/timohahaa/userbot/internal/service"
)

type userbotRoutes struct {
	telegramService service.TelegramService
}

func newUserbotRoutes(g *echo.Group, telegramService service.TelegramService) {
	r := &userbotRoutes{
		telegramService: telegramService,
	}

	g.GET("/channel/:channelId", r.GetChannel)
	g.GET("/channel", r.ListChannels)
	g.POST("/channel", r.SaveChannelByName)
	g.POST("/channel/:channelId/react", r.ReactNewPost)
}

type getChannelOutput struct {
	Id        int64  `json:"channel_id"`
	UserCount int    `json:"user_count"`
	Name      string `json:"channel_name"`
}

// GET /api/v1/channel/{channelId}
func (r *userbotRoutes) GetChannel(c echo.Context) error {
	cId := c.Param("channelId")
	channelId, err := strconv.ParseInt(cId, 10, 64)
	if err != nil {
		newErrorMessage(c, http.StatusBadRequest, ErrInvalidPathParameter.Error())
		return err
	}

	channel, err := r.telegramService.GetChannelByChannelId(c.Request().Context(), channelId)
	if errors.Is(err, service.ErrChannelNotFound) {
		newErrorMessage(c, http.StatusNotFound, ErrChannelNotFound.Error())
		return ErrChannelNotFound
	}
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, ErrInternalServer.Error())
		return err
	}

	output := getChannelOutput{}
	output.Id = channel.Id
	output.UserCount = channel.UserCount
	output.Name = channel.Name

	return c.JSON(http.StatusOK, output)
}

type listChannelsOutput struct {
	Data []getChannelOutput
}

// GET /api/v1/channel
func (r *userbotRoutes) ListChannels(c echo.Context) error {
	channels, err := r.telegramService.ListChannels(c.Request().Context())
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, ErrInternalServer.Error())
		return err
	}

	output := listChannelsOutput{}
	for _, channel := range channels {
		output.Data = append(output.Data, getChannelOutput{Id: channel.Id, UserCount: channel.UserCount, Name: channel.Name})
	}

	return c.JSON(http.StatusOK, output)
}

type saveChannelByNameInput struct {
	ChannelName string `json:"channel_name" validate:"required"`
}

// POST /api/v1/channel
func (r *userbotRoutes) SaveChannelByName(c echo.Context) error {
	var input saveChannelByNameInput
	if err := c.Bind(&input); err != nil {
		newErrorMessage(c, http.StatusBadRequest, ErrInvalidRequestBody.Error())
		return err
	}

	if validationErr := c.Validate(input); validationErr != nil {
		newErrorMessage(c, http.StatusBadRequest, validationErr.Error())
		return validationErr
	}

	err := r.telegramService.SaveChannelByName(c.Request().Context(), input.ChannelName)
	if errors.Is(err, service.ErrChannelNotFound) {
		newErrorMessage(c, http.StatusNotFound, ErrChannelNotFound.Error())
		return ErrChannelNotFound
	}
	if errors.Is(err, service.ErrFailedToSaveChannel) {
		newErrorMessage(c, http.StatusInternalServerError, ErrFailedToSaveChannel.Error())
		return ErrFailedToSaveChannel
	}
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, ErrInternalServer.Error())
		return err
	}

	return c.NoContent(http.StatusCreated)
}

// POST /api/v1/channel/{channelId}/react
func (r *userbotRoutes) ReactNewPost(c echo.Context) error {
	cId := c.Param("channelId")
	channelId, err := strconv.ParseInt(cId, 10, 64)
	if err != nil {
		newErrorMessage(c, http.StatusBadRequest, ErrInvalidPathParameter.Error())
		return err
	}

	err = r.telegramService.ReactNewPost(c.Request().Context(), channelId)
	if errors.Is(err, service.ErrChannelNotFound) {
		newErrorMessage(c, http.StatusNotFound, ErrChannelNotFound.Error())
		return ErrChannelNotFound
	}
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, ErrInternalServer.Error())
		return err
	}

	return c.NoContent(http.StatusOK)
}
