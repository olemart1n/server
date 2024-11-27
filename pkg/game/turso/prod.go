package turso

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)


func ProdTursoDB() (*sql.DB, error) {
	url := os.Getenv("TURSO_DATABASE_URL")
	if url == "" {
		return nil, fmt.Errorf("TURSO_DATABASE_URL environment variable not set")
	}
	token := os.Getenv("TURSO_AUTH_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("TURSO_AUTH_TOKEN environment variable not set")
	}
	connectionUrl := url + "?authToken=" + token
	db, err := sql.Open("libsql", connectionUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to open db %s: %s", url, err)
	}

	// Ensure the database connection is valid
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %s", err)
	}

	return db, nil
}
