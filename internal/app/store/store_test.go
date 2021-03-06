package store_test

import (
	"os"
	"testing"
)
var(
	databaseURL string
)
// 
func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
		if databaseURL == "" {
			databaseURL = "host=localhost dbname=mydb sslmode=disable user=postgres password=pass"
		}
		os.Exit(m.Run())
	}