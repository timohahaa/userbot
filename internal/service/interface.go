package service

import (
	"context"

	"github.com/timohahaa/userbot/internal/entity"
)

type TelegramService interface {
	GetRandomAccount(ctx context.Context) (entity.Account, error)
	SaveChannelByName(ctx context.Context, channelName string) error
	ListChannels(ctx context.Context) ([]entity.Channel, error)
	ReactNewPost(ctx context.Context, channelId int64) error
	GetChannelByChannelId(ctx context.Context, id int64) (entity.Channel, error)
}

var _ TelegramService = &telegramService{}
