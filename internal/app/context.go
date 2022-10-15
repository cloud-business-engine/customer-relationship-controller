package app

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type (
	requestIdKey struct{}
)

func GetLogger(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}

func GetRequestID(ctx context.Context) uuid.UUID {
	return ctx.Value(requestIdKey{}).(uuid.UUID) //nolint:forcetypeassert
}

func ContextWithLogger(ctx context.Context, logger *zerolog.Logger) context.Context {
	return logger.WithContext(ctx)
}

func ContextWithRequestID(ctx context.Context, requestID uuid.UUID) context.Context {
	return context.WithValue(ctx, requestIdKey{}, requestID)
}
