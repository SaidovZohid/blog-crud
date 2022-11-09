package storage

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"os"
	_ "github.com/lib/pq"
)

var (
	dbManager *DBManager
	user = "postgres"
	password = "1234"
	host = "localhost"
	port = 5432
	dbname = "blog"
)

func TestMain(m *testing.M) {
	connstr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatalf("Failed to open connection: %v", err)
	}
	dbManager = NewDBManager(db)
	os.Exit(m.Run())
}