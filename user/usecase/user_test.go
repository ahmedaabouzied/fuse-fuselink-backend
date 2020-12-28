package usecase

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"bitbucket.org/MoMoLab-dev/fuse.link-backend/config"
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/entities"
	mockuserrepo "bitbucket.org/MoMoLab-dev/fuse.link-backend/user/repository/mocks"
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	t.Run("Returns the created user", func(t *testing.T) {
		mockUserRepo := &mockuserrepo.UserRepository{}
		repos := config.Repositories{
			UserRepository: mockUserRepo,
		}
		userUsecase := NewUserUsecase(&repos)
		testUser := &entities.User{}
		mockUserRepo.On("Create", mock.Anything, testUser).Return(testUser, nil)
		_, err := userUsecase.CreateUser(context.Background(), testUser)
		assert.Nil(t, err)
	})
	t.Run("Returns wrapped error on repository error", func(t *testing.T) {
		mockUserRepo := &mockuserrepo.UserRepository{}
		repos := config.Repositories{
			UserRepository: mockUserRepo,
		}
		userUsecase := NewUserUsecase(&repos)
		testUser := &entities.User{}
		expectedErr := fmt.Errorf("create user error")
		mockUserRepo.On("Create", mock.Anything, testUser).Return(testUser, expectedErr)
		_, err := userUsecase.CreateUser(context.Background(), testUser)
		t.Log(err)
		t.Log(errors.Is(err, utils.RepositoryError))
		assert.True(t, errors.Is(err, utils.RepositoryError))
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("Handles missing user ID in context", func(t *testing.T) {
		mockUserRepo := &mockuserrepo.UserRepository{}
		repos := config.Repositories{
			UserRepository: mockUserRepo,
		}
		userUsecase := NewUserUsecase(&repos)
		emptyCtx := context.Background()
		res, err := userUsecase.Update(emptyCtx, "userID", &entities.UpdateUserRequest{})
		assert.Nil(t, res)
		assert.True(t, errors.Is(err, utils.ErrorUnauthorizedRequest))
	})
}
