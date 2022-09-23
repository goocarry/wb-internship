package store_test

import (
	"os"
	"testing"
)

var (
	dsn string
)

func TestMain(m *testing.M) {
	dsn = "postgres://postgres:postgres@localhost:5432/wbl0_test?sslmode=disable"

	os.Exit(m.Run())
}
