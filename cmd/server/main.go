package main

import (
	"log"
	"os"

	"example.com/internal-service/internal/di"
	"go.uber.org/zap"
)

func main() {
	app, err := di.SetupApplication()
	if err != nil {
		log.Fatalf("Failed to setup application: %v", err)
	}

	exiteCode := 0

	defer func() {
		os.Exit(exiteCode)
	}()

	app.GetLogger().Info("Running Application")

	if err := app.Run(); err != nil {
		app.GetLogger().Error("Application failed", zap.Error(err))
		exiteCode = 1
	}
}
