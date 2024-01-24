package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	BadRequest          func(ctx *gin.Context, err error)
	InternalServerError func(ctx *gin.Context, err error)
	Success             func(ctx *gin.Context, resp any)
)

func init() {
	BadRequest = func(ctx *gin.Context, err error) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	InternalServerError = func(ctx *gin.Context, err error) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	Success = func(ctx *gin.Context, resp any) {
		ctx.JSON(http.StatusOK, resp)
	}
}
