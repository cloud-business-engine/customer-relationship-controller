package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func MakeGinEngine(logger *zerolog.Logger) *gin.Engine {
	engine := gin.New()
	SetupMiddlewares(engine, logger)
	SetupHandlers(engine)
	return engine
}
