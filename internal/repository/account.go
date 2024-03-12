package repository

import (
	"context"

	"github.com/timohahaa/userbot/internal/entity"

	"github.com/timohahaa/postgres"
)

type accountRepo struct {
	db *postgres.Postgres
}

func NewAccountRepo(pg *postgres.Postgres) *accountRepo {
	return &accountRepo{
		db: pg,
	}
}

func (r *accountRepo) GetRandomAccount(ctx context.Context) (entity.Account, error) {
	// not the fastest method, but the accounts table is very small ( << 1_000_000 rows, maybe will never reach 100_000)
	sql, args, _ := r.db.Builder.
		Select("user_id", "phone_number", "username", "session_string").
		From("accounts").
		OrderBy("RANDOM()").
		Limit(1).
		ToSql()

	var account entity.Account
	err := r.db.ConnPool.QueryRow(ctx, sql, args).Scan(&account)

	if err != nil {
		return entity.Account{}, err
	}

	return account, nil
}
