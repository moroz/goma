package main

import (
	"fmt"
	"log"

	"github.com/alexedwards/argon2id"
)

func main() {
	password := "hunter2"
	params := &argon2id.Params{
		Memory:      19 * 1024,
		Iterations:  2,
		Parallelism: 1,
		SaltLength:  16,
		KeyLength:   16,
	}
	hash, err := argon2id.CreateHash(password, params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hashed password:", hash)

	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Match: %v\n", match)
}
