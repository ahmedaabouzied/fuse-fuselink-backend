package usecase

import (
	"context"

	"bitbucket.org/MoMoLab-dev/fuse.link-backend/config"
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/entities"
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/user"
	"github.com/pkg/errors"
)

type UserUsecase struct {
	userRepo user.UserRepository
}

func NewUserUsecase(repositories *config.Repositories) user.UserUsecase {
	return &UserUsecase{
		userRepo: repositories.UserRepository,
	}
}

func (u *UserUsecase) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	createdUser, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "repository error while creating user")
	}
	return createdUser, nil
}

func (u *UserUsecase) Update(ctx context.Context, userID string, updateRequest *entities.UpdateUserRequest) (*entities.User, error) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	userToUpdate, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "repository error while getting user")
	}
	userToUpdate.Username = updateRequest.Username
	userToUpdate.SocialAccounts = updateRequest.SocialAccounts
	updatedUser, err := u.userRepo.Update(ctx, userToUpdate)
	if err != nil {
		return nil, errors.Wrap(err, "repository error while updating user")
	}
	return updatedUser, nil
}

func (u *UserUsecase) Delete(ctx context.Context, userID string) (*entities.User, error) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	userToDelete, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "repository error while deleting user")
	}
	err = u.userRepo.Delete(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "repository error while deleting user")
	}
	return userToDelete, nil
}

func (u *UserUsecase) GetByUsername(ctx context.Context, username string) (*entities.User, error) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	user, err := u.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, errors.Wrap(err, "repository error while getting user")
	}
	return user, nil
}

func (u *UserUsecase) GetByCognitoID(ctx context.Context, cognitoID string) (*entities.User, error) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()
	user, err := u.userRepo.GetByCognitoID(ctx, cognitoID)
	if err != nil {
		return nil, errors.Wrap(err, "repository error while getting user")
	}
	return user, nil
}
