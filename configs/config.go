package configs

import (
	"fmt"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var (
	config appConfigStruct
	doOnce sync.Once
)

type appConfigStruct struct {
	// Server settings
	ServerPort string
	// PostgreSql
	PostgreSqlHost     string
	PostgreSqlPort     string
	PostgreSqlName     string
	PostgreSqlUsername string
	PostgreSqlPassword string
	// MongoDB
	MongodbAddress    string
	MongodbName       string
	MongodbCollection string
	MongodbUsername   string
	MongodbPassword   string
	// jwt
	JwtSecretKey      string
	JwtTokenExpired   time.Duration // in second
	JwtRefreshExpired time.Duration // in second
}

func init() {
	doOnce.Do(func() {
		viper.SetConfigFile(".env") // dev env TODO change for production
		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}

		config = load()
	})
}

func load() appConfigStruct {
	jwtTokenExp := viper.GetString("JWT_TOKEN_EXPIRED")
	jwtRefreshExp := viper.GetString("JWT_REFRESH_EXPIRED")

	jwtTokenDuration, _ := time.ParseDuration(jwtTokenExp)
	jwtRefreshDuration, _ := time.ParseDuration(jwtRefreshExp)

	return appConfigStruct{
		// Server settings
		ServerPort: viper.GetString("SERVER_PORT"),
		// PostgreSql
		PostgreSqlHost:     viper.GetString("POSTGRESQL_HOST"),
		PostgreSqlPort:     viper.GetString("POSTGRESQL_PORT"),
		PostgreSqlName:     viper.GetString("POSTGRESQL_NAME"),
		PostgreSqlUsername: viper.GetString("POSTGRESQL_USERNAME"),
		PostgreSqlPassword: viper.GetString("POSTGRESQL_PASSWORD"),
		// MongoDB
		MongodbAddress:    viper.GetString("MONGODB_ADDRESS"),
		MongodbName:       viper.GetString("MONGODB_NAME"),
		MongodbCollection: viper.GetString("MONGODB_COLLECTION"),
		MongodbUsername:   viper.GetString("MONGODB_USERNAME"),
		MongodbPassword:   viper.GetString("MONGODB_PASSWORD"),
		// Jwt
		JwtSecretKey:      viper.GetString("JWT_SECRET_KEY"),
		JwtTokenExpired:   jwtTokenDuration,   // in second
		JwtRefreshExpired: jwtRefreshDuration, // in second
	}
}

func Get() *appConfigStruct {
	return &config
}
