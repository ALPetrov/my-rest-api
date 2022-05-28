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
			databaseURL = "root:mysql57@/mytestdb"
		}
		os.Exit(m.Run())
	}