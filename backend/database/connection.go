package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Функция подключения к базе данных
func Connection() *sql.DB {
	dsnLink := "postgresql://database_ptds_user:GM7f9yn1NrxBgSpVfa9re8IC91iW688q@dpg-d879pptckfvc73a1qa6g-a.virginia-postgres.render.com/database_ptds"
	db, err := sql.Open("postgres", dsnLink)
	if err != nil {
		log.Println("Не удалось подключиться к базе:", err)
		return nil
	}
	return db
}
