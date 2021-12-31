package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(viper.GetString("SERVER_READ_TIMEOUT"))
	writeTimeoutSecondsCount, _ := strconv.Atoi(viper.GetString("SERVER_WRITE_TIMEOUT"))
	idleTimeoutSecondsCount, _ := strconv.Atoi(viper.GetString("SERVER_IDLE_TIMEOUT"))

	// Return Fiber configuration.
	return fiber.Config{
		AppName:      "MyPoints App v0.0.1",
		ServerHeader: "MyPoints Server",
		ReadTimeout:  time.Second * time.Duration(readTimeoutSecondsCount),
		WriteTimeout: time.Second * time.Duration(writeTimeoutSecondsCount),
		IdleTimeout:  time.Second * time.Duration(idleTimeoutSecondsCount),
	}
}
