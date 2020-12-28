package usecase

import (
	"context"
	"testing"

	"bitbucket.org/MoMoLab-dev/fuse.link-backend/config"
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/entities"
	mockuserrepo "bitbucket.org/MoMoLab-dev/fuse.link-backend/user/repository/mocks"
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
}
