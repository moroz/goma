package store_test

import (
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
	s.Equal(before+1, after)
}
