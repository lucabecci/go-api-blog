package data

import (
	"database/sql"
	"os"

	//db server
	_ "github.com/lib/pq"
)

func getConnection() (*sql.DB, error) {
	uri := os.Getenv("DB_URI")
	return sql.Open("postgres", uri)
}
