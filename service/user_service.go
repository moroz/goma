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

func AuthenticateUserByEmailPassword(email, password string) (*types.User, error) {
	var user types.User
}
