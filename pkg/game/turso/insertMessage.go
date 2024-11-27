package turso

import (
	"database/sql"
	"fmt"
)


func InsertMessage(db *sql.DB, username string, message string, id string ) error {
	if db == nil {
		return fmt.Errorf("database connection is nil")
	}
	err := db.Ping()
	if err != nil {
	return	fmt.Errorf("failed to ping database: %v", err)
	}
	stmt, err := db.Prepare("INSERT INTO messages (senderUsername, message, senderId ) VALUES (?, ?, ?)")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, message, id)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %v", err)
	}
	return nil
}