package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const ConnString = "postgres://postgres:postgres@localhost/goma_dev?sslmode=disable"

func main() {
	db := sqlx.MustConnect("postgres", ConnString)
	var version string
	err := db.QueryRow("select version()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(version)
}
