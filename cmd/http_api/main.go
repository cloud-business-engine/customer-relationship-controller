package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"

	"github.com/cloud-business-engine/customer-relationship-controller/internal"
	"github.com/cloud-business-engine/customer-relationship-controller/internal/presentation/api"
)

const (
	service = "cloud-business-engine.customer-relationship-controller"
)

func main() {
	logger := internal.NewLogger(service, os.Stdout)
	defer panicHandler(logger)

	engine := api.MakeGinEngine(logger)
	if err := engine.Run(":8080"); err != nil {
		err = fmt.Errorf("error during engine work: %w", err)
		panic(err)
	}
}

func panicHandler(logger *zerolog.Logger) {
	if err := recover(); err != nil {
		message := fmt.Sprintf("error: %s", err)
		logger.Fatal().Str("logger", "root").Msg(message)
	}
}
