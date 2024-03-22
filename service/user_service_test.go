package service_test

import (
	"database/sql"

	"github.com/alexedwards/argon2id"
	"github.com/moroz/goma/service"
	"github.com/moroz/goma/store"
	"github.com/moroz/goma/types"
)

func (s *ServiceTestSuite) TestAuthenticateUserByEmailPassword() {
	password := "hunter2"
	store := store.NewUserStore(s.db)
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	s.NoError(err)
	user := &types.User{
		Email:        "authenticate@example.com",
		PasswordHash: &hash,
	}
	user, err = store.InsertUser(user)
	s.NoError(err)

	srv := service.NewUserService(s.db)
	actual, err := srv.AuthenticateUserByEmailPassword(user.Email, password)
	s.NoError(err)
	s.Equal(user.ID, actual.ID)

	actual, err = srv.AuthenticateUserByEmailPassword(user.Email, "invalid")
	s.Nil(actual)
	s.ErrorIs(err, types.ErrInvalidPassword)

	actual, err = srv.AuthenticateUserByEmailPassword("invalid@example.com", password)
	s.Nil(actual)
	s.ErrorIs(err, sql.ErrNoRows)
}
