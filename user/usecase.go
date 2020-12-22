package user

import (
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/entities"
	"context"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, user entities.User) (entities.User, error)
	Update(ctx context.Context, userID string, updateRequest entities.UpdateUserRequest) (entities.User, error)
	Delete(ctx context.Context, userID string) (entities.User, error)
}
