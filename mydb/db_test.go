package mydb

import (
	"database/sql"
	"testing"

	"tis-gf-api/secrets"
)

func TestGetDb(t *testing.T) {
	t.Run("testing db connection", func(t *testing.T) {
		db, err := sql.Open("mssql", secrets.SQL_CONN_STR)
		if err != nil {
			defer db.Close()
			t.Errorf("Not able to connect to the dataase")
		}
		err = db.Ping()
		if err != nil {
			t.Errorf("Not Ping")
		}
		defer db.Close()
	})

}
