package store_test

import (
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/moroz/goma/store"
	"github.com/moroz/goma/types"
)

func countUsers(db *sqlx.DB) (int, error) {
	var result int
	err := db.QueryRow(`select count(*) from users`).Scan(&result)
	return result, err
}

func (s *StoreTestSuite) TestInsertUser() {
	passwordHash := "test"
	user := types.User{
		Email:        "insert@example.com",
		PasswordHash: &passwordHash,
	}
	us := store.NewUserStore(s.db)

	before, err := countUsers(s.db)
	s.NoError(err)
	actual, err := us.InsertUser(&user)
	s.NoError(err)
	after, err := countUsers(s.db)
	s.NoError(err)

	s.Greater(actual.ID, 0)
	s.Equal(user.Email, actual.Email)
	s.Equal(user.PasswordHash, actual.PasswordHash)
	s.Equal(before+1, after)
}

func (s *StoreTestSuite) TestGetUserByEmail() {
	passwordHash := "test"
	newUser := types.User{
		Email:        "by-email@example.com",
		PasswordHash: &passwordHash,
	}
	us := store.NewUserStore(s.db)
	user, err := us.InsertUser(&newUser)
	s.NoError(err)

	examples := []string{user.Email, strings.ToUpper(user.Email), "By-Email@Example.Com"}
	for _, email := range examples {
		actual, err := us.GetUserByEmail(email)
		s.NoError(err)
		s.Equal(user.ID, actual.ID)
	}

	actual, err := us.GetUserByEmail("non-existent@example.com")
	s.ErrorIs(err, sql.ErrNoRows)
	s.Nil(actual)
}
