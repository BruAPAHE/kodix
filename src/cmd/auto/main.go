package main

import (
	"kodix/src/internal/app"
	"kodix/src/internal/logger"
)

func main() {
	application := app.New()

	if err := application.Run(); err != nil {
		logger.Logger.Fatalw("Application didn't start",
			"error", err,
		)
	}
}
