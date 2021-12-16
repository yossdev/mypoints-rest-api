package startserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yossdev/mypoints-rest-api/configs"
	"log"
	"os"
	"os/signal"
)

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := a.Listen(configs.Get().ServerPort); err != nil {
		log.Fatalf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

// StartServer func for starting server with manual shutdown.
func StartServer(a *fiber.App) {
	// Run server.
	if err := a.Listen(configs.Get().ServerPort); err != nil {
		log.Fatalf("Oops... Server is not running! Reason: %v", err)
	}
}
