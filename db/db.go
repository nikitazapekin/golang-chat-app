package db

import (
	"database/sql"
//	"encoding/json"
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
 


func FindUserByUsername(username string) (string, string, string, error) {
	if DB == nil {
		return "", "", "", fmt.Errorf("database connection is not established. Call Connect function first")
	}

	var (
		foundUsername string
		country       string
		tel           string
	)

	query := `
	SELECT username, country, tel
	FROM user_data
	WHERE username = $1
	`

	row := DB.QueryRow(query, username)
	err := row.Scan(&foundUsername, &country, &tel)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", "", fmt.Errorf("user with username '%s' not found", username)
		//	return "", "", "", "user with username  not found"
		}
		return "", "", "", fmt.Errorf("failed to find user: %v", err)
	}

	return foundUsername, country, tel, nil
}









/*
func FindUserDataByUsername(username string) (string, string, string, error) {
	if DB == nil {
		return "", "", "", fmt.Errorf("database connection is not established. Call Connect function first")
	}

	var (
		foundUsername string
		country       string
		tel           string
		token string
		avatar string
		describtion string
		chats json
	)

	query := `
	SELECT username, country, tel, token, chats, avatar, describtion
	FROM user_data
	WHERE username = $1
	`

	row := DB.QueryRow(query, username)
	err := row.Scan(&foundUsername, &country, &tel, &token &chats &avatar &describtion)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", "", fmt.Errorf("user with username '%s' not found", username)
		//	return "", "", "", "user with username  not found"
		}
		return "", "", "", fmt.Errorf("failed to find user: %v", err)
	}

	return foundUsername, country, tel, token, chats, avatar, describtion, nil
}
*/



/*
func FindUserDataByUsername(username string) (string, string, string, string, []byte, string, string, error) {
    if DB == nil {
        return "", "", "", "", nil, "", "", fmt.Errorf("database connection is not established. Call Connect function first")
    }

    var (
        foundUsername string
        country       string
        tel           string
        token         string
		avatar        sql.NullString
     //   avatar        string
        describtion   string
        chats         []byte  
    )

    query := `
    SELECT username, country, tel, token, chats, avatar, describtion
    FROM user_data
    WHERE username = $1
    `

    row := DB.QueryRow(query, username)
    err := row.Scan(&foundUsername, &country, &tel, &token, &chats, &avatar, &describtion)
    if err != nil {
        if err == sql.ErrNoRows {
            return "", "", "", "", nil, "", "", fmt.Errorf("user with username '%s' not found", username)
        }
        return "", "", "", "", nil, "", "", fmt.Errorf("failed to find user: %v", err)
    }

    return foundUsername, country, tel, token, chats, avatar, describtion, nil
}

*/


func FindUserDataByUsername(username string) (string, string, string, string, []byte, string, string, error) {
    if DB == nil {
        return "", "", "", "", nil, "", "", fmt.Errorf("database connection is not established. Call Connect function first")
    }

    var (
        foundUsername string
        country       string
        tel           string
    //    token         string
token      sql.NullString 
      //  avatar        sql.NullString 
	  avatar        sql.NullString 
	 describtion       sql.NullString 
       // describtion   string
        chats         []byte  
    )

    query := `
    SELECT username, country, tel, token, chats, avatar, describtion
    FROM user_data
    WHERE username = $1
    `

    row := DB.QueryRow(query, username)
    err := row.Scan(&foundUsername, &country, &tel, &token, &chats, &avatar, &describtion)
    if err != nil {
        if err == sql.ErrNoRows {
            return "", "", "", "", nil, "", "", fmt.Errorf("user with username '%s' not found", username)
        }
        return "", "", "", "", nil, "", "", fmt.Errorf("failed to find user: %v", err)
    }
 
    var avatarString string
    if avatar.Valid {
        avatarString = avatar.String
    }


	var tokenString string
    if token.Valid {
        tokenString = token.String
    }
	var describtionString string
    if describtion.Valid {
        describtionString = describtion.String
    }


    return foundUsername, country, tel, tokenString, chats, avatarString, describtionString, nil
}








 
func UpdateUserToken(username, token string) error {
	query := `UPDATE user_data SET token = $1 WHERE username = $2`
	_, err := DB.Exec(query, token, username)
	if err != nil {
		return fmt.Errorf("failed to update user token: %v", err)
	}
	return nil
}