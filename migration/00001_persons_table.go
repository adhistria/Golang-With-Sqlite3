package migration

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00001, Down00001)
}

func Up00001(tx *sql.Tx) error {
	query := "CREATE TABLE USERS(" +
		"id int NOT NULL AUTO_INCREMENT," +
		"username varchar," +
		"name varchar," +
		"PRIMARY KEY ID)"
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func Down00001(tx *sql.Tx) error {
	query := "DROP TABLE USERS"
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
//  goose sqlite3 ./simple_rest_api.db create users_table sql
