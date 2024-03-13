package main

import (
	"fmt"
	"log"

	"github.com/alexedwards/argon2id"
)

func main() {
	password := "hunter2"
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hashed password:", hash)

	passwordsToCheck := []string{"hunter2", "Hunter2", "hunter3"}
	for _, password := range passwordsToCheck {
		match, err := argon2id.ComparePasswordAndHash(password, hash)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\"%s\":\t %v\n", password, match)
	}
}
