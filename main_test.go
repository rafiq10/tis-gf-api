package main

import (
	"log"
	"os"
	"testing"
	"tis-gf-api/mydb"
)

func TestMain(m *testing.M) {
	// initTests()
	os.Exit(m.Run())
}

func initTests() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error while opening SQL file: '%s'; error message: %s", path, err.Error())
	}
	path = path + "/init_db.sql"
	err = mydb.ExecSqlFromFile(path)
	if err != nil {
		log.Fatalf("Error while executing SQL from file: '%s', error message: %s", path, err.Error())
	}

}

func TestGracefulShutdown(t *testing.T) {
	// Just to run init test setup fo all tests:
	// Create combo and config tables in testDB
}
