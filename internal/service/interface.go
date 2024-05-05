package service

import (
	"context"

	"github.com/timohahaa/userbot/internal/entity"
)

type TelegramService interface {
	GetRandomAccount(ctx context.Context) (entity.Account, error)
	SaveChannelByName(ctx context.Context, channelName string) error
	ReactNewPost(ctx context.Context, channelId int64) error
}

var _ TelegramService = &telegramService{}
