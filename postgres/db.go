package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func DB() *sql.DB {
	var psqlUrl = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"postgres",
		"Dost0n1k",
		"events",
	)

	psqlConn, err := sql.Open("postgres", psqlUrl)
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	return psqlConn
}
