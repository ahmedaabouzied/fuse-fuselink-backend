package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mockauth "bitbucket.org/MoMoLab-dev/fuse.link-backend/auth/mocks"
	"bitbucket.org/MoMoLab-dev/fuse.link-backend/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleare(t *testing.T) {
	t.Run("handle error parsing token", func(t *testing.T) {
		// setup a mock auth handler and server config
		mockAuthHandler := &mockauth.AuthHandler{}
		serverConfig := &config.Server{
			AuthHandler: mockAuthHandler,
		}

		gin.SetMode(gin.TestMode)

		// Return error on parsing token
		mockAuthToken := "MOCK_AUTH_TOKEN"
		authError := fmt.Errorf("%s", "AUTH_ERROR")
		mockAuthHandler.On("ParseAuthToken", mockAuthToken).Return(nil, authError)

		// create a gin router with middlware
		router := gin.Default()
		router.GET("/test", authMiddleware(serverConfig))
		request := httptest.NewRequest("GET", "/test", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, request)
		assert.Equal(t, http.StatusUnauthorized, rr.Code)
	})
	t.Run("trims Bearer prefix from auth header", func(t *testing.T) {
		// setup a mock auth handler and server config
		mockAuthHandler := &mockauth.AuthHandler{}
		serverConfig := &config.Server{
			AuthHandler: mockAuthHandler,
		}

		gin.SetMode(gin.TestMode)

		// Return error on parsing token
		mockAuthToken := "Bearer MOCK_AUTH_TOKEN"
		trimmedAuthToken := "MOCK_AUTH_TOKEN"
		mockAuthHandler.On("ParseAuthToken", trimmedAuthToken).Return(nil, nil)

		// create a gin router with middlware
		router := gin.Default()
		router.GET("/test", authMiddleware(serverConfig))
		request := httptest.NewRequest("GET", "/test", nil)
		request.Header.Add("Authorization", mockAuthToken)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, request)
		mockAuthHandler.AssertCalled(t, "ParseAuthToken", trimmedAuthToken)
	})
}
