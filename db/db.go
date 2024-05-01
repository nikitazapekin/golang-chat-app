package db

import (
	"database/sql"
	"fmt"
	"log"
		_ "github.com/lib/pq"
)
var (
	DB *sql.DB
)

func Connect() {
	fmt.Println("DB WORK")
	dbConnStr := "user=Nikita password=Backend dbname=chat sslmode=disable"
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	DB = db
	CreateTable()
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
}

func PingDB() error {
	err := DB.Ping()
	if err != nil {
		return err
	}
	return nil
}



func AddUser(username, country, tel string) error {
	if DB == nil {
		return fmt.Errorf("database connection is not established. Call Connect function first")
	}

	query := `
	INSERT INTO user_data (username, country, tel)
	VALUES ($1, $2, $3)
	`

	_, err := DB.Exec(query, username, country, tel)
	if err != nil {
		return fmt.Errorf("failed to add user: %v", err)
	}

	return nil
}
 