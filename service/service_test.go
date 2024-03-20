package service_test

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	db *sqlx.DB
}

func (s *ServiceTestSuite) SetupTest() {
	conn := os.Getenv("TEST_DATABASE_URL")
	s.db = sqlx.MustConnect("postgres", conn)
	s.db.MustExec("truncate users cascade")
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
