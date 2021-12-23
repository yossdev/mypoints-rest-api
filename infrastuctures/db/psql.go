package db

import (
	"fmt"
	"github.com/yossdev/mypoints-rest-api/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PsqlDB interface {
	DB() *gorm.DB
}

type psqlDB struct {
	db *gorm.DB
}

func NewPsqlClient() PsqlDB {
	var db *gorm.DB
	var err error

	dbHost := configs.Get().PostgreSqlHost
	dbPort := configs.Get().PostgreSqlPort
	dbName := configs.Get().PostgreSqlName
	dbUser := configs.Get().PostgreSqlUsername
	dbPassword := configs.Get().PostgreSqlPassword

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//	defer db.Close()

	if err != nil {
		panic("failed to connect database")
	}

	return &psqlDB{
		db: db,
	}
}

func (p psqlDB) DB() *gorm.DB {
	return p.db
}
