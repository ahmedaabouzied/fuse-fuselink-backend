package user

import (
	"context"

	"bitbucket.org/MoMoLab-dev/fuse.link-backend/entities"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) (*entities.User, error)
	Delete(ctx context.Context, userID string) error
	GetByID(ctx context.Context, ID string) (*entities.User, error)
	GetByUsername(ctx context.Context, username string) (*entities.User, error)
}
