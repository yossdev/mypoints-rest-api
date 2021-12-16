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

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
