package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LandingHandler(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
}
