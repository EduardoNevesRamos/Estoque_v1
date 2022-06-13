package util

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func MockGin() (*httptest.ResponseRecorder, *gin.Context) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	return w, ctx
}
