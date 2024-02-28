package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var ConnString = MustGetEnv("DATABASE_URL")

func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("FATAL: Environment variable %s is not set!", key)
	}
	return value
}

func main() {
	db := sqlx.MustConnect("postgres", ConnString)
	var version string
	err := db.QueryRow("select version()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(version)
}
