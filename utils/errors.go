package utils

import (
	"net/http"

	"errors"
	"github.com/gin-gonic/gin"
)

var (
	// InvalidAuthorizationHeader is returned when received
	// auth token is invalid.
	ErrorInvalidAuthorizationHeader = errors.New("authorization error, please login again")
	ErrorUnauthorizedRequest        = errors.New("unauthorized request, please try again")
	NotFoundError                   = errors.New("not found")
	ErrorUsernameAlreadyExists      = errors.New("username already exists, please choose another username")
	RepositoryError                 = errors.New("repsitory error")
	ErrorParsingRequest             = errors.New("error parsing request arguments, please try again")
	ErrorCreatingUser               = errors.New("error creating user, please try again")
	ErrorUpdatingUser               = errors.New("error updating user, please try again")

	// appErrors maps errors to status codes.
	appErrors = map[error]int{
		ErrorInvalidAuthorizationHeader: http.StatusUnauthorized,
		ErrorParsingRequest:             http.StatusBadRequest,
	}
)

// HandleAPIError matches the top most error in the
// given error chain to an HTTP status code. If no
// match was found, it will set the response status
// to Internal Server Error 500.
// It will also send along a response JSON body
// containing the error chain and a message refering
// to the top most error in the chain.
func HandleAPIError(c *gin.Context, err error) {
	// unwrap the error to get the top error in the error chain.
	topErr := errors.Unwrap(err)

	// Get the matched status code for the given error
	var statusCode int
	statusCode, ok := appErrors[topErr]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	if errors.Is(err, RepositoryError) {
		statusCode = http.StatusInternalServerError
	}
	if errors.Is(err, NotFoundError) {
		statusCode = http.StatusNotFound
	}

	// Write to the HTTP response
	c.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
	return
}
