package service

import (
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

func (us *UserService) AuthenticateUserByEmailPassword(email, password string) (*types.User, error) {
	user, err := us.store.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if _, err := user.CheckPassword(password); err != nil {
		return nil, err
	}

	return user, nil
}
