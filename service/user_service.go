package service

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/moroz/goma/store"
	"github.com/moroz/goma/types"
)

type UserService struct {
	store store.UserStore
}

func NewUserService(db *sqlx.DB) UserService {
	return UserService{store: store.NewUserStore(db)}
}

var InvalidPasswordError = errors.New("invalid password")

func (us *UserService) AuthenticateUserByEmailPassword(email, password string) (*types.User, error) {
	user, err := us.store.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	match, err := user.CheckPassword(password)
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, InvalidPasswordError
	}
	return user, nil
}
