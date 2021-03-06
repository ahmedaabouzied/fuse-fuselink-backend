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
	t.Run("Handle repository errog with getting user", func(t *testing.T) {
		mockUserRepo := &mockuserrepo.UserRepository{}
		repos := config.Repositories{
			UserRepository: mockUserRepo,
		}
		userUsecase := NewUserUsecase(&repos)
		testCtx := context.WithValue(context.Background(), entities.UserIDContextKey, "USER_ID")
		mockUserRepo.On("GetByID", mock.Anything, "USER_ID").Return(nil, errors.New("REPOSITORY_ERROR"))
		res, err := userUsecase.Update(testCtx, "USER_ID", &entities.UpdateUserRequest{})
		assert.Nil(t, res)
		assert.True(t, errors.Is(err, utils.RepositoryError))
	})
	t.Run("Returns error when user ID does not match current user", func(t *testing.T) {
		mockUserRepo := &mockuserrepo.UserRepository{}
		repos := config.Repositories{
			UserRepository: mockUserRepo,
		}
		userUsecase := NewUserUsecase(&repos)
		testUser := &entities.User{
			CognitoUserID: "ANOTHER_USER_ID",
		}
		testCtx := context.WithValue(context.Background(), entities.UserIDContextKey, "USER_ID")
		mockUserRepo.On("GetByID", mock.Anything, "USER_ID").Return(testUser, nil)
		res, err := userUsecase.Update(testCtx, "USER_ID", &entities.UpdateUserRequest{})
		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.True(t, errors.Is(err, utils.ErrorUnauthorizedRequest))
	})
	t.Run("Set the updated fields from the updates object", func(t *testing.T) {
		mockUserRepo := &mockuserrepo.UserRepository{}
		repos := config.Repositories{
			UserRepository: mockUserRepo,
		}
		userUsecase := NewUserUsecase(&repos)
		testUser := &entities.User{
			CognitoUserID: "USER_ID",
		}
		testCtx := context.WithValue(context.Background(), entities.UserIDContextKey, "USER_ID")
		userUpdateReq := &entities.UpdateUserRequest{
			Username: "updated",
			SocialAccounts: []entities.SocialAccount{{
				Platform: "facebook",
			}},
		}
		updatedTestUser := &entities.User{
			CognitoUserID:  testUser.CognitoUserID,
			Username:       userUpdateReq.Username,
			SocialAccounts: userUpdateReq.SocialAccounts,
		}
		mockUserRepo.On("GetByID", mock.Anything, "USER_ID").Return(testUser, nil)
		mockUserRepo.On("Update", mock.Anything, mock.Anything).Return(nil, nil)
		_, err := userUsecase.Update(testCtx, "USER_ID", userUpdateReq)
		assert.Nil(t, err)
		mockUserRepo.AssertCalled(t, "Update", mock.Anything, updatedTestUser)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("Handles missing user ID in context", func(t *testing.T) {
		mockUserRepo := &mockuserrepo.UserRepository{}
		repos := config.Repositories{
			UserRepository: mockUserRepo,
		}
		userUsecase := NewUserUsecase(&repos)
		emptyCtx := context.Background()
		res, err := userUsecase.Delete(emptyCtx, "userID")
		assert.Nil(t, res)
		assert.True(t, errors.Is(err, utils.ErrorUnauthorizedRequest))
	})
	t.Run("Handle repository errog with getting user", func(t *testing.T) {
		mockUserRepo := &mockuserrepo.UserRepository{}
		repos := config.Repositories{
			UserRepository: mockUserRepo,
		}
		userUsecase := NewUserUsecase(&repos)
		testCtx := context.WithValue(context.Background(), entities.UserIDContextKey, "USER_ID")
		mockUserRepo.On("GetByID", mock.Anything, "USER_ID").Return(nil, errors.New("REPOSITORY_ERROR"))
		res, err := userUsecase.Delete(testCtx, "USER_ID")
		assert.Nil(t, res)
		assert.True(t, errors.Is(err, utils.RepositoryError))
	})
	t.Run("Returns error when user ID does not match current user", func(t *testing.T) {
		mockUserRepo := &mockuserrepo.UserRepository{}
		repos := config.Repositories{
			UserRepository: mockUserRepo,
		}
		userUsecase := NewUserUsecase(&repos)
		testUser := &entities.User{
			CognitoUserID: "ANOTHER_USER_ID",
		}
		testCtx := context.WithValue(context.Background(), entities.UserIDContextKey, "USER_ID")
		mockUserRepo.On("GetByID", mock.Anything, "USER_ID").Return(testUser, nil)
		res, err := userUsecase.Delete(testCtx, "USER_ID")
		assert.NotNil(t, err)
		assert.Nil(t, res)
		assert.True(t, errors.Is(err, utils.ErrorUnauthorizedRequest))
	})
	t.Run("Calls Delete user repo method with the correct user ID", func(t *testing.T) {
		mockUserRepo := &mockuserrepo.UserRepository{}
		repos := config.Repositories{
			UserRepository: mockUserRepo,
		}
		userUsecase := NewUserUsecase(&repos)
		testUser := &entities.User{
			CognitoUserID: "USER_ID",
		}
		testCtx := context.WithValue(context.Background(), entities.UserIDContextKey, "USER_ID")
		mockUserRepo.On("GetByID", mock.Anything, "USER_ID").Return(testUser, nil)
		mockUserRepo.On("Delete", mock.Anything, mock.Anything).Return(nil, nil)
		_, err := userUsecase.Delete(testCtx, "USER_ID")
		assert.Nil(t, err)
		mockUserRepo.AssertCalled(t, "Delete", mock.Anything, "USER_ID")
	})
}

func TestGetByUsername(t *testing.T) {
	t.Run("Calls GetByUsername with the correct username", func(t *testing.T) {
		mockUserRepo := &mockuserrepo.UserRepository{}
		repos := config.Repositories{
			UserRepository: mockUserRepo,
		}
		userUsecase := NewUserUsecase(&repos)
		testUser := &entities.User{
			CognitoUserID: "USER_ID",
		}
		testCtx := context.WithValue(context.Background(), entities.UserIDContextKey, "USER_ID")
		mockUserRepo.On("GetByUsername", mock.Anything, "username").Return(testUser, nil)
		_, err := userUsecase.GetByUsername(testCtx, "username")
		assert.Nil(t, err)
		mockUserRepo.AssertCalled(t, "GetByUsername", mock.Anything, "username")
	})
}

func TestGetByCognitoID(t *testing.T) {
	t.Run("Calls GetByCognitoID with the correct ID", func(t *testing.T) {
		mockUserRepo := &mockuserrepo.UserRepository{}
		repos := config.Repositories{
			UserRepository: mockUserRepo,
		}
		userUsecase := NewUserUsecase(&repos)
		testUser := &entities.User{
			CognitoUserID: "USER_ID",
		}
		testCtx := context.WithValue(context.Background(), entities.UserIDContextKey, "USER_ID")
		mockUserRepo.On("GetByCognitoID", mock.Anything, "username").Return(testUser, nil)
		_, err := userUsecase.GetByCognitoID(testCtx, "username")
		assert.Nil(t, err)
		mockUserRepo.AssertCalled(t, "GetByCognitoID", mock.Anything, "username")
	})
}
