package services

import (
	"context"

	"github.com/timohahaa/userbot/internal/entity"
)

type AccountService interface {
	GetRandomAccount(ctx context.Context) (entity.Account, error)
}
