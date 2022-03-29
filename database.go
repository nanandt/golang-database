package golangdatabase

import (
	"database/sql"
	"time"
)

func GetConnecttion() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)                  // set minimal koneksi yg dibuat
	db.SetMaxOpenConns(100)                 // set maksimal koneksi yg dibuat
	db.SetConnMaxIdleTime(20 * time.Minute) // set berapa lama koneksi yg tdk digunakan akan dihapus
	db.SetConnMaxLifetime(60 * time.Minute) // berapa lama koneksi bisa digunakan

	return db

}
