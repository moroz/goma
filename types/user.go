package types

import "time"

type User struct {
	ID           int       `db:"id"`
	Email        string    `db:"email"`
	PasswordHash *string   `db:"password_hash"`
	InsertedAt   time.Time `db:"inserted_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
