// database connection
package utils

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Substitute your DB details
const (
	host     = "localhost"
	port     = 5455
	User     = "golang"
	password = "brewess"
	dbname   = "postgres"
)

func GetConnection() *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, User, password, dbname)

	db, err := sqlx.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	log.Println("DB Connection established...")
	return db
}
