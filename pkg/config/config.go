// config.go
package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/sijms/go-ora/v2"
)

var DB *sql.DB

// InitializeDB menginisialisasi koneksi ke database Oracle
func InitializeDB() {
	// Gunakan format URL untuk koneksi go-ora
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_SID := os.Getenv("DB_SID")
	dsn := "oracle://" + DB_USER + ":" + DB_PASSWORD + "@" + DB_HOST + ":" + DB_PORT + "/" + DB_SID

	// dsn := "oracle://system:yourpassword@localhost:1521/XEPDB1"

	var err error
	DB, err = sql.Open("oracle", dsn)
	if err != nil {
		log.Fatalf("Gagal membuka koneksi: %v", err)
	}

	// Mengecek koneksi
	if err = DB.Ping(); err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	fmt.Println("Berhasil terhubung ke Oracle Database!")
}

// CloseDB menutup koneksi ke database
func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatalf("Gagal menutup koneksi database: %v", err)
	}
	fmt.Println("Koneksi ke Oracle Database ditutup.")
}
