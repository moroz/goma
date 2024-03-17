package store

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/moroz/goma/types"
)

type userStore struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) userStore {
	return userStore{db: db}
}

func (us *userStore) InsertUser(user *types.User) (*types.User, error) {
	var result types.User
	err := us.db.Get(&result, `insert into users (email, password_hash) values ($1, $2) returning *`, user.Email, user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("InsertUser: %w", err)
	}
	return &result, nil
}

func (us *userStore) GetUserByEmail(email string) (*types.User, error) {
	var result types.User
	err := us.db.Get(&result, `select * from users where email = $1`, email)
	if err != nil {
		return nil, fmt.Errorf("GetUserByEmail: %w", err)
	}
	return &result, nil
}
