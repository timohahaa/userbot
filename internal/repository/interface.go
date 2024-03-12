package repository

import (
	"context"

	"github.com/timohahaa/userbot/internal/entity"
)

type AccountRepository interface {
	GetRandomAccount(ctx context.Context) (entity.Account, error)
}

type ChannelRepository interface {
	SaveChannel(ctx context.Context, channel entity.Channel) error
	GetChannelByChannelId(ctx context.Context, id int64) (entity.Channel, error)
}
