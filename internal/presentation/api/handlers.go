package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupHandlers(engine *gin.Engine) {
	engine.GET("/time", GetTime)
}

func GetTime(ctx *gin.Context) {
	responsePayload := map[string]string{
		"time": time.Now().Format(time.RFC3339),
	}
	ctx.JSON(http.StatusOK, responsePayload)
}
