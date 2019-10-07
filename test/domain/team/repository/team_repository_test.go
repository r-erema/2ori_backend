package test

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
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

	f, _ := os.Getwd()
	fmt.Println(f)
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	db, err = sql.Open("postgres", psqlInfo)
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

func TestFindByIds(t *testing.T) {

	prepareTestDatabase()

}
