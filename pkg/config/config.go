package config

import (
	"database/sql"
	"log"

	_ "github.com/godror/godror"
)

var DB *sql.DB

func InitDatabase() {
	var err error
	DB, err = sql.Open("godror", "system/yourpassword@localhost:1521/XEPDB1")
	if err != nil {
		log.Fatalf("Gagal menghubungkan ke database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Database tidak dapat diakses: %v", err)
	}
	log.Println("Berhasil terhubung ke database!")
}
