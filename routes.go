package main

import (
	"fmt"
	errors "github.com/pkg/errors"
	"strings"

	"bitbucket.org/MoMoLab-dev/fuse.link-backend/config"
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/entities"
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/utils"
	"github.com/gin-gonic/gin"
)

func startServer(serverConfig *config.Server) {
	router := gin.Default()
	router.Run(fmt.Sprintf(":%s", serverConfig.Port))
}

func authMiddleware(serverConfig *config.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		parsedToken, err := serverConfig.AuthHandler.ParseAuthToken(token)
		if err != nil {
			err = errors.Wrap(err, utils.ErrorInvalidAuthorizationHeader.Error())
		}
		c.Set(entities.UserIDContextKey, parsedToken["cognito:username"].(string))
	}
}
