package store_test

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type StoreTestSuite struct {
	suite.Suite
	db *sqlx.DB
}

func (s *StoreTestSuite) SetupTest() {
	conn := os.Getenv("TEST_DATABASE_URL")
	s.db = sqlx.MustConnect("postgres", conn)
	s.db.MustExec("truncate users cascade")
}

func TestStoreTestSuite(t *testing.T) {
	suite.Run(t, new(StoreTestSuite))
}
