package types

import (
	"errors"
	"fmt"
	"strings"
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

var ErrNoPasswordSet = errors.New("no password set for the user")
var ErrInvalidPassword = errors.New("invalid password")

func (u *User) CheckPassword(password string) (bool, error) {
	if u.PasswordHash == nil || !strings.HasPrefix(*u.PasswordHash, "$argon2id$") {
		return false, ErrNoPasswordSet
	}

	match, err := argon2id.ComparePasswordAndHash(password, *u.PasswordHash)
	if err != nil {
		return false, fmt.Errorf("CheckPassword: %w", err)
	}
	if !match {
		return false, ErrInvalidPassword
	}
	return true, nil
}
