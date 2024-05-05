package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/timohahaa/postgres"
	"github.com/timohahaa/userbot/internal/entity"
)

type telegramRepo struct {
	db *postgres.Postgres
}

func NewTelegramRepo(pg *postgres.Postgres) *telegramRepo {
	return &telegramRepo{
		db: pg,
	}
}

func (r *telegramRepo) GetRandomAccount(ctx context.Context) (entity.Account, error) {
	// not the fastest method, but the accounts table is very small ( << 1_000_000 rows, maybe will never reach 100_000)
	sql, args, _ := r.db.Builder.
		Select("user_id", "phone_number", "username", "session_string").
		From("accounts").
		OrderBy("RANDOM()").
		Limit(1).
		ToSql()

	var account entity.Account
	err := r.db.ConnPool.QueryRow(ctx, sql, args).
		Scan(
			&account.UserId,
			&account.PhoneNumber,
			&account.Username,
			&account.SessionString,
		)

	if err != nil {
		return entity.Account{}, err
	}

	return account, nil
}

func (r *telegramRepo) SaveAccount(ctx context.Context, acc entity.Account) error {
	sql, args, _ := r.db.Builder.
		Insert("accounts").
		Columns("user_id", "phone_number", "username", "session_string").
		Values(acc.UserId, acc.PhoneNumber, acc.Username, acc.SessionString).
		ToSql()

	_, err := r.db.ConnPool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *telegramRepo) SaveChannel(ctx context.Context, channel entity.Channel) error {
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

func (r *telegramRepo) GetChannelByChannelId(ctx context.Context, id int64) (entity.Channel, error) {
	sql, args, _ := r.db.Builder.
		Select("channel_id", "access_hash", "user_count", "name").
		From("channels").
		Where("channel_id = ?", id).
		ToSql()

	var channel entity.Channel
	err := r.db.ConnPool.QueryRow(ctx, sql, args).
		Scan(
			&channel.Id,
			&channel.AccessHash,
			&channel.UserCount,
			&channel.Name,
		)

	if errors.Is(err, pgx.ErrNoRows) {
		return entity.Channel{}, ErrChannelNotFound
	}
	if err != nil {
		return entity.Channel{}, err
	}

	return channel, nil
}

func (r *telegramRepo) ListChannels(ctx context.Context) ([]entity.Channel, error) {
	sql, args, _ := r.db.Builder.
		Select("channel_id", "access_hash", "user_count", "name").
		From("channels").
		ToSql()

	rows, err := r.db.ConnPool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var channels []entity.Channel
	for rows.Next() {
		var channel entity.Channel
		err := rows.Scan(
			&channel.Id,
			&channel.AccessHash,
			&channel.UserCount,
			&channel.Name,
		)

		if err != nil {
			channels = append(channels, channel)
		}
	}

	return channels, nil
}
