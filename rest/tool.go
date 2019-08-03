package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func json200(context *gin.Context, v interface{}) {
	context.JSON(http.StatusOK, v)
}
func json500(context *gin.Context, v interface{}) {
	context.JSON(http.StatusInternalServerError, v)
}
