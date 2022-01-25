package connector

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/user/datasource/postgresql/"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("postgres", postgresql.ConnectionString)
	if err != nil {
		panic(err)
	}
	DB = db
}
