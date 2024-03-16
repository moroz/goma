package store_test

import (
	"github.com/moroz/goma/store"
	"github.com/moroz/goma/types"
)

func (s *StoreTestSuite) countUsers() int {
	var result int
	row := s.db.QueryRow(`select count(*) from users`)
	s.NoError(row.Err())
	err := row.Scan(&result)
	s.NoError(err)
	return result
}

func (s *StoreTestSuite) TestInsertUser() {
	password := "test"
	user := types.User{
		Email:        "insert@example.com",
		PasswordHash: &password,
	}
	us := store.NewUserStore(s.db)

	before := s.countUsers()
	actual, err := us.InsertUser(&user)
	after := s.countUsers()

	s.NoError(err)
	s.Greater(actual.ID, 0)
	s.Equal(before+1, after)
}
