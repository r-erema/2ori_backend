package test

import (
	"database/sql"
	_ "github.com/lib/pq"
	"gopkg.in/testfixtures.v2"
	_ "gopkg.in/testfixtures.v2"
	"log"
	"os"
	"testing"
)

var (
	db       *sql.DB
	fixtures *testfixtures.Context
)

func TestMain(m *testing.M) {
	var err error
	db, err = sql.Open("postgres", "dbname_test")
	if err != nil {
		log.Fatal(err)
	}

	fixtures, err = testfixtures.NewFolder(db, &testfixtures.PostgreSQL{}, "./fixture")
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())

}

func prepareTestDatabase() {
	if err := fixtures.Load(); err != nil {
		log.Fatal(err)
	}
}
