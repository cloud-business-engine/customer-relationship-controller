package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/cloud-business-engine/customer-relationship-controller/internal/app"
)

func SetupHandlers(engine *gin.Engine) {
	engine.GET("/time", GetTime)
}

func GetTime(ctx *gin.Context) {
	responsePayload := map[string]string{
		"request_id": app.GetRequestID(ctx.Request.Context()).String(),
		"time":       time.Now().Format(time.RFC3339),
		"host_ip":    ctx.Request.Host,
	}
	fmt.Println(responsePayload) //nolint:forbidigo
	ctx.JSON(http.StatusOK, responsePayload)
}
