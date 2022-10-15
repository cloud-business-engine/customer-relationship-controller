package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/cloud-business-engine/customer-relationship-controller/internal/app"
)

const (
	RequestIDHeader   = "X-Request-Id"
	ContentTypeHeader = "Content-Type"
)

func SetupMiddlewares(engine *gin.Engine, logger *zerolog.Logger) {
	engine.Use(
		LoggerMiddleware(logger),
		RequestIDMiddleware(),
		AccessLogMiddleware(logger),
		PanicRecoveryMiddleware(logger),
	)
}

func LoggerMiddleware(logger *zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loggerContext := app.ContextWithLogger(ctx.Request.Context(), logger)
		ctx.Request = ctx.Request.WithContext(loggerContext)
	}
}

func RequestIDMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID, err := resolveRequestID(ctx.Request)
		if err != nil {
			ctx.String(http.StatusBadRequest, "X-Request-Id must be in uuid4 format")
			ctx.Abort()
			return
		}

		ctx.Writer.Header().Set(RequestIDHeader, requestID.String())
		ctx.Request = ctx.Request.WithContext(
			app.ContextWithRequestID(ctx.Request.Context(), requestID),
		)
	}
}

func AccessLogMiddleware(logger *zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()

		requestTime := time.Since(start)
		logger.Info().
			Str("logger", "access").
			Str("transport.protocol", ctx.Request.Proto).
			Str("request.id", app.GetRequestID(ctx.Request.Context()).String()).
			Str("request.user_agent", ctx.Request.UserAgent()).
			Str("request.method", ctx.Request.Method).
			Str("request.url", ctx.Request.URL.String()).
			Str("request.content_type", ctx.Request.Header.Get(ContentTypeHeader)).
			Int64("request.content_length.bytes", ctx.Request.ContentLength).
			Int("response.status", ctx.Writer.Status()).
			Str("response.content_type", ctx.Writer.Header().Get(ContentTypeHeader)).
			Int("response.content_length.bytes", ctx.Writer.Size()).
			Dur("response.time.ms", requestTime).
			Send()
	}
}

func PanicRecoveryMiddleware(logger *zerolog.Logger) gin.HandlerFunc {
	handler := func(ctx *gin.Context, err any) {
		message := fmt.Sprintf("panic recovered: %s", err)
		logger.
			Error().
			Str("logger", "recovery").
			Str("request.id", app.GetRequestID(ctx.Request.Context()).String()).
			Msg(message)
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
	return gin.CustomRecoveryWithWriter(nil, handler)
}

func resolveRequestID(req *http.Request) (uuid.UUID, error) {
	requestIDString := req.Header.Get(RequestIDHeader)
	if requestIDString == "" {
		return uuid.New(), nil
	}
	return uuid.Parse(requestIDString) //nolint:wrapcheck // нет необходимости
}
