package repository

import (
	"context"

	"github.com/timohahaa/userbot/internal/entity"

	"github.com/timohahaa/postgres"
)

type channelRepo struct {
	db *postgres.Postgres
}

func NewChannelRepo(pg *postgres.Postgres) *channelRepo {
	return &channelRepo{
		db: pg,
	}
}

func (r *channelRepo) SaveChannel(ctx context.Context, channel entity.Channel) error {
	sql, args, _ := r.db.Builder.
		Insert("channels").
		Columns("channel_id", "access_hash", "user_count", "name").
		Values(channel.Id, channel.AccessHash, channel.UserCount, channel.Name).
		ToSql()

	_, err := r.db.ConnPool.Exec(ctx, sql, args)
	if err != nil {
		return err
	}

	return nil
}

func (r *channelRepo) GetChannelByChannelId(ctx context.Context, id int64) (entity.Channel, error) {
	sql, args, _ := r.db.Builder.
		Select("channel_id", "access_hash", "user_count", "name").
		From("channels").
		Where("channel_id = ?", id).
		ToSql()

	var channel entity.Channel
	err := r.db.ConnPool.QueryRow(ctx, sql, args).Scan(&channel)

	if err != nil {
		return entity.Channel{}, err
	}

	return channel, nil
}
