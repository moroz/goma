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

func (u *User) CheckPassword(password string) (bool, error) {
	if u.PasswordHash == nil {
		return false, errors.New("no password set for the user")
	}
	return argon2id.ComparePasswordAndHash(password, *u.PasswordHash)
}
