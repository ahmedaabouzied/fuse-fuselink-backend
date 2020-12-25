package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	errors "github.com/pkg/errors"
)

var (
	// InvalidAuthorizationHeader is returned when received
	// auth token is invalid.
	ErrorInvalidAuthorizationHeader = errors.New("authorization error, please login again")
	ErrorParsingRequest             = errors.New("error parsing request arguments, please try again")
	ErrorCreatingUser               = errors.New("error creating user, please try again")

	// appErrors maps errors to status codes.
	appErrors = map[error]int{
		ErrorInvalidAuthorizationHeader: http.StatusUnauthorized,
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

	// Write to the HTTP response
	c.JSON(statusCode, gin.H{
		"error":   err,
		"message": topErr,
	})
}
