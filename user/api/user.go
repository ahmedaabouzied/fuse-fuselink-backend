package api

import (
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/config"
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/entities"
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/user"
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/utils"
	"context"
	"github.com/gin-gonic/gin"
	errors "github.com/pkg/errors"
	"net/http"
)

type UserAPIHandler struct {
	userUsecase user.UserUsecase
}

func NewUserAPI(server *config.Server) *UserAPIHandler {
	return &UserAPIHandler{
		userUsecase: server.UserUsecase,
	}
}

// CreateUser creates a new user record with a username and
// empty social profile records.
func (h *UserAPIHandler) CreateUser(c *gin.Context) {
	ctx := context.Background()
	cognitoID, ok := c.Get(entities.UserIDContextKey)
	if !ok {
		utils.HandleAPIError(c, errors.Wrap(errors.New("token does not encode a cognito user ID"), utils.ErrorInvalidAuthorizationHeader.Error()))
	}
	var userReq entities.CreateUserRequest
	err := c.BindJSON(&userReq)
	if err != nil {
		utils.HandleAPIError(c, errors.Wrap(err, utils.ErrorParsingRequest.Error()))
	}
	user := &entities.User{
		Username:      userReq.Username,
		CognitoUserID: cognitoID.(string),
	}
	createdUser, err := h.userUsecase.CreateUser(ctx, user)
	if err != nil {
		err = errors.Wrap(err, utils.ErrorCreatingUser.Error())
	}
	c.JSON(http.StatusCreated, gin.H{
		"user": createdUser,
	})
}

func (h *UserAPIHandler) UpdateUser(c *gin.Context) {
	ctx := context.Background()
	cognitoID, ok := c.Get(entities.UserIDContextKey)
	if !ok {
		utils.HandleAPIError(c, errors.Wrap(errors.New("token does not encode a cognito user ID"), utils.ErrorInvalidAuthorizationHeader.Error()))
		return
	}
	ctx = context.WithValue(ctx, entities.UserIDContextKey, cognitoID)
	userID := c.Param("id")
	var updateUserReq entities.UpdateUserRequest
	err := c.BindJSON(&updateUserReq)
	if err != nil {
		utils.HandleAPIError(c, errors.Wrap(err, utils.ErrorParsingRequest.Error()))
		return
	}
	updatedUser, err := h.userUsecase.Update(ctx, userID, &updateUserReq)
	if err != nil {
		utils.HandleAPIError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": updatedUser,
	})
	return
}

func (h *UserAPIHandler) DeleteUser(c *gin.Context) {
	ctx := context.Background()
	cognitoID, ok := c.Get(entities.UserIDContextKey)
	if !ok {
		utils.HandleAPIError(c, errors.Wrap(errors.New("token does not encode a cognito user ID"), utils.ErrorInvalidAuthorizationHeader.Error()))
		return
	}
	ctx = context.WithValue(ctx, entities.UserIDContextKey, cognitoID)
	userID := c.Param("id")
	deletedUser, err := h.userUsecase.Delete(ctx, userID)
	if err != nil {
		utils.HandleAPIError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": deletedUser,
	})
	return
}

func (h *UserAPIHandler) GetUserByUsername(c *gin.Context) {
	ctx := context.Background()
	username := c.Param("username")
	user, err := h.userUsecase.GetByUsername(ctx, username)
	if err != nil {
		utils.HandleAPIError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}

func (h *UserAPIHandler) GetUser(c *gin.Context) {
	ctx := context.Background()
	cognitoID, ok := c.Get(entities.UserIDContextKey)
	if !ok {
		utils.HandleAPIError(c, errors.Wrap(errors.New("token does not encode a cognito user ID"), utils.ErrorInvalidAuthorizationHeader.Error()))
	}
	user, err := h.userUsecase.GetByCognitoID(ctx, cognitoID.(string))
	if err != nil {
		utils.HandleAPIError(c, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}
