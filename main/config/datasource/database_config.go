package datasource

import (
	"database/sql"
	"os"
	_ "github.com/lib/pq"
	"fmt"
)

// sql connection handler
var DB *sql.DB

// init sql handler
func init() {
	dbname := os.Getenv("DATASOURCE_DATABASE")
	user := os.Getenv("DATASOURCE_USER")
	password := os.Getenv("DATASOURCE_PASSWORD")

	dataSourceName := fmt.Sprintf("dbname=%s user=%s password=%s sslmode=disable", dbname, user, password)
	database, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		panic(err)
	}
	defer database.Close()

	DB = database
}
