package repository

import (
	"context"

	"github.com/timohahaa/userbot/internal/entity"
)

type TelegramRepository interface {
	GetRandomAccount(ctx context.Context) (entity.Account, error)
	SaveAccount(ctx context.Context, acc entity.Account) error
	SaveChannel(ctx context.Context, channel entity.Channel) error
	GetChannelByChannelId(ctx context.Context, id int64) (entity.Channel, error)
	ListChannels(ctx context.Context) ([]entity.Channel, error)
}

var _ TelegramRepository = &telegramRepo{}
