package utils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleAPIError(t *testing.T) {

	t.Run("Handle 500 errors", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		HandleAPIError(c, RepositoryError)
		assert.Equal(t, rr.Result().StatusCode, http.StatusInternalServerError)
	})

	t.Run("Handle 404 errors", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		HandleAPIError(c, NotFoundError)
		assert.Equal(t, rr.Result().StatusCode, http.StatusNotFound)
	})

	t.Run("Set status code to 500 if error type is not recognized", func(t *testing.T) {
		rr := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rr)
		HandleAPIError(c, fmt.Errorf("unrecognized error"))
		assert.Equal(t, rr.Result().StatusCode, http.StatusInternalServerError)
	})
}
