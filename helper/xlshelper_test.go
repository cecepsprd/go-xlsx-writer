package helper_test

import (
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestInsertUsergroup(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	defer mockDB.Close()
	sqlxDB = sqlx.NewDb(mockDB, "sqlmock")
}
