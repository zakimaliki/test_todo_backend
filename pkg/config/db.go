package config

import (
	"os"
	"strconv"

	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() {
	DB_PORT, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dialector := oracle.New(oracle.Config{
		DSN: oracle.BuildUrl(
			os.Getenv("DB_HOST"),
			DB_PORT,
			os.Getenv("DB_SID"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			map[string]string{
				"CONNECTION TIMEOUT": "90",
			},
		),
	})

	var err error
	DB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	if sqlDB, err := DB.DB(); err == nil {
		oracle.AddSessionParams(sqlDB, map[string]string{
			"TIME_ZONE": "+08:00",
		})
	}
}
