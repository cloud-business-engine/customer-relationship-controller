package internal

import (
	"io"

	"github.com/rs/zerolog"
)

func NewLogger(service string, output io.Writer) *zerolog.Logger {
	logger := zerolog.New(output).
		With().
		Timestamp().
		Str("service", service).
		Logger()
	return &logger
}
