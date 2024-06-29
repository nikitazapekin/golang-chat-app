package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
)



/*
func CreateTable() {
	fmt.Println("START")
	if DB == nil {
		log.Fatal("Database connection is not established. Call Connect function first.")
	}
	query := `
	CREATE TABLE IF NOT EXISTS user_data (
		username VARCHAR(255),
		country VARCHAR(255),
		tel VARCHAR(255),
		role VARCHAR(255),
		registration_data VARCHAR(255),
		avatar VARCHAR(255),
		last_time_at_network VARCHAR(255),
		chats JSONB,
        user_id VARCHAR(255),
        describtion VARCHAR(255),
		token VARCHAR(255),
		access_token VARCHAR(255)
	);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Println("Table user_data created successfully.")
//	AddUser("Nikita", "hh", "kk")
}
*/


func CreateTable() {
    fmt.Println("START")
    if DB == nil {
        log.Fatal("Database connection is not established. Call Connect function first.")
    }
    query := `
    CREATE TABLE IF NOT EXISTS user_data (
        username VARCHAR(255),
        country VARCHAR(255),
        tel VARCHAR(255),
        role VARCHAR(255), -- Changed to role column
        registration_data VARCHAR(255),
        avatar JSONB,
        last_time_at_network VARCHAR(255),
        chats JSONB,
        user_id VARCHAR(255),
        describtion VARCHAR(255),
        token VARCHAR(255),
        access_token VARCHAR(255)
    );
    `
    _, err := DB.Exec(query)
    if err != nil {
        log.Fatalf("Failed to create table: %v", err)
    }
    fmt.Println("Table user_data created successfully.")
}

 

/*
func CreateUserTable(username string) {
    fmt.Println("START")
    if DB == nil {
        log.Fatal("Database connection is not established. Call Connect function first.")
    }
    query := `
    CREATE TABLE IF NOT EXISTS user_data (
        username VARCHAR(255),
        country VARCHAR(255),
        tel VARCHAR(255),
        role VARCHAR(255), 
        registration_data VARCHAR(255),
        avatar JSONB,
        last_time_at_network VARCHAR(255),
        chats JSONB,
        user_id VARCHAR(255),
        describtion VARCHAR(255),
        token VARCHAR(255),
        access_token VARCHAR(255)
    );
    `
    _, err := DB.Exec(query)
    if err != nil {
        log.Fatalf("Failed to create table: %v", err)
    }
    fmt.Println("Table user_data created successfully.")
}


func CreateUserTable(username string) {
    fmt.Println("START")
    if DB == nil {
        log.Fatal("Database connection is not established. Call Connect function first.")
    }
    tableName := fmt.Sprintf("user_data_%s", username)
    query := fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %s (
        chats JSONB, 
        user_id VARCHAR(255)
    );
    `, tableName)
    _, err := DB.Exec(query)
    if err != nil {
        log.Fatalf("Failed to create table: %v", err)
    }
    fmt.Printf("Table %s created successfully.\n", tableName)
}

    */