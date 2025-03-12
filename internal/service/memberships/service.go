package memberships

import (
	"context"
	"time"

	"github.com/tamaaa13/fastcampus/internal/configs"
	"github.com/tamaaa13/fastcampus/internal/model/memberships"
)

type membershipRepo interface {
	GetUSer(ctx context.Context, email, username string, userID int64) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model memberships.UserModel) error
	InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error
	GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*memberships.RefreshTokenModel, error)
}

type service struct {
	cfg            *configs.Config
	membershipRepo membershipRepo
}

func NewService(cfg *configs.Config, membershipRepo membershipRepo) *service {
	return &service{
		cfg:            cfg,
		membershipRepo: membershipRepo,
	}
}
