package types

import (
	"errors"
	"time"

	"github.com/alexedwards/argon2id"
)

type User struct {
	ID           int       `db:"id"`
	Email        string    `db:"email"`
	PasswordHash *string   `db:"password_hash"`
	InsertedAt   time.Time `db:"inserted_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func (u *User) CheckPassword(password string) (match bool, err error) {
	if u.PasswordHash == nil || u.PasswordHash == "" {
		return false, errors.New("no password set for the user")
	}
	match, err = argon2id.ComparePasswordAndHash()
}
