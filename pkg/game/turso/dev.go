package turso

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func DevTursoDB () (*sql.DB, error) {
	dbName := "file:./game.db"
	db, err := sql.Open("libsql", dbName)
	if err != nil {
		
	fmt.Fprintf(os.Stderr, "failed to open db %s", err)
	os.Exit(1)
	}
	
	return db, err
}