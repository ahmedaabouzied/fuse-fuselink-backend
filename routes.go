package main

import (
	"fmt"
	errors "github.com/pkg/errors"
	"strings"

	"bitbucket.org/MoMoLab-dev/fuse.link-backend/config"
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/entities"
	userapi "bitbucket.org/MoMoLab-dev/fuse.link-backend/user/api"
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/utils"
	"github.com/gin-gonic/gin"
)

func startServer(serverConfig *config.Server) {
	router := gin.Default()
	userAPIHandler := userapi.NewUserAPI(serverConfig)
	v1 := router.Group("/v1/api")
	{
		v1.GET("/user", authMiddleware(serverConfig), userAPIHandler.GetUser)
		v1.GET("/user/:username", userAPIHandler.GetUserByUsername)
		v1.POST("/user", authMiddleware(serverConfig), userAPIHandler.CreateUser)
		v1.DELETE("/user/:id", authMiddleware(serverConfig), userAPIHandler.DeleteUser)
		v1.PUT("/user/:id", authMiddleware(serverConfig), userAPIHandler.UpdateUser)
	}
	router.Run(fmt.Sprintf(":%s", serverConfig.Port))
}

func authMiddleware(serverConfig *config.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.HandleAPIError(c, fmt.Errorf("unauthorized: %w", utils.ErrorInvalidAuthorizationHeader))
			c.Abort()
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		parsedToken, err := serverConfig.AuthHandler.ParseAuthToken(token)
		if err != nil {
			err = errors.Wrap(err, utils.ErrorInvalidAuthorizationHeader.Error())
			utils.HandleAPIError(c, fmt.Errorf("%s: %w", err.Error(), utils.ErrorInvalidAuthorizationHeader))
			c.Abort()
			return
		}
		c.Set(entities.UserIDContextKey, parsedToken["cognito:username"].(string))
	}
}
