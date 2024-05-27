package db

import (
	"database/sql"
//	"encoding/json"
	"fmt"
	"log"
	"github.com/google/uuid"
    "time"
    "encoding/json"
    "math/rand"
	_ "github.com/lib/pq"
)
var (
	DB *sql.DB
)
/*
type Avatar struct {
    Text  string `json:"text"`
    Color string `json:"color"`
    URL   string `json:"url"`
}

type User struct {
    username    string
    country     string
    tel         string
    token       string
    chats       []byte
    avatar      Avatar
    description string
}
*/

type Avatar struct {
    Text  string `json:"text"`
    Color string `json:"color"`
    URL   string `json:"url"`
}

type User struct {
    Username    string `json:"username"`
    Country     string `json:"country"`
    Tel         string `json:"tel"`
    Token       string `json:"token"`
    Chats       []byte `json:"chats"`
    Avatar      Avatar `json:"avatar"`
    Description string `json:"description"`
}



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

 
/*
 func AddUser(username, country, tel string) error {
    if DB == nil {
        return fmt.Errorf("database connection is not established. Call Connect function first")
    }

    
    userID := uuid.New().String()

    query := `
    INSERT INTO user_data (user_id, username, country, tel)
    VALUES ($1, $2, $3, $4)
    `

    _, err := DB.Exec(query, userID, username, country, tel)
    if err != nil {
        return fmt.Errorf("failed to add user: %v", err)
    }

    return nil
}
*/


func generateRandomColor() string {
    rand.Seed(time.Now().UnixNano())
    // Define color ranges for orange, blue, and green
    colors := []string{
        "orange", "orangered", "darkorange", // Orange shades
        "blue", "deepskyblue", "dodgerblue", // Blue shades
        "green", "limegreen", "seagreen",    // Green shades
    }
    return colors[rand.Intn(len(colors))]
}


/*
func AddUser(username, country, tel string) error {
    if DB == nil {
        return fmt.Errorf("database connection is not established. Call Connect function first")
    }

    userID := uuid.New().String()

    color := generateRandomColor()
avatarHTML :=    fmt.Sprintf(
    `color: %s; text: %c`,
    color, username[0],
)
    query := `
    INSERT INTO user_data (user_id, username, country, tel, avatar)
    VALUES ($1, $2, $3, $4, $5)
    `

    _, err := DB.Exec(query, userID, username, country, tel, avatarHTML)
    if err != nil {
        return fmt.Errorf("failed to add user: %v", err)
    }

    return nil
}

*/



func AddUser(username, country, tel string) error {
    if DB == nil {
        return fmt.Errorf("database connection is not established. Call Connect function first")
    }

    userID := uuid.New().String()
url:= ""
    color := generateRandomColor()
    avatar := Avatar{
        Color: color,
        Text:  string(username[0]),
        URL: string(url),
    }

 
    avatarJSON, err := json.Marshal(avatar)
    if err != nil {
        return fmt.Errorf("failed to marshal avatar: %v", err)
    }

    query := `
    INSERT INTO user_data (user_id, username, country, tel, avatar)
    VALUES ($1, $2, $3, $4, $5)
    `

    _, err = DB.Exec(query, userID, username, country, tel, string(avatarJSON))
    if err != nil {
        return fmt.Errorf("failed to add user: %v", err)
    }

    return nil
}


func CreateTableChats() {
    fmt.Println("START")
    if DB == nil {
        log.Fatal("Database connection is not established. Call Connect function first.")
    }
    query := `
    CREATE TABLE IF NOT EXISTS user_data (
        
        chats JSONB,
        groups JSONB
    );
    `
    _, err := DB.Exec(query)
    if err != nil {
        log.Fatalf("Failed to create table: %v", err)
    }
    fmt.Println("Table user_data created successfully.")
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
func FindUsersDataByUsername(username string) ([]User, error) {
    if DB == nil {
        return nil, fmt.Errorf("database connection is not established. Call Connect function first")
    }

    query := `
    SELECT username, country, tel, token, chats, avatar, describtion
    FROM user_data
    WHERE username LIKE '%' || $1 || '%'
    `

    rows, err := DB.Query(query, username)
    if err != nil {
        return nil, fmt.Errorf("failed to query users: %v", err)
    }
    defer rows.Close()

    var users []User

    for rows.Next() {
        var (
            foundUsername string
            country       string
            tel           string
            token         sql.NullString
            chats         []byte
            avatar        sql.NullString
            description   sql.NullString
        )

        err := rows.Scan(&foundUsername, &country, &tel, &token, &chats, &avatar, &description)
        if err != nil {
            return nil, fmt.Errorf("failed to scan user: %v", err)
        }

        user := User{
            Username:    foundUsername,
            Country:     country,
            Tel:         tel,
            Chats:       chats,
        }

        if token.Valid {
            user.Token = token.String
        }

        if avatar.Valid {
            err := json.Unmarshal([]byte(avatar.String), &user.Avatar)
            if err != nil {
                return nil, fmt.Errorf("failed to unmarshal avatar: %v", err)
            }
        }

        if description.Valid {
            user.Description = description.String
        }
        fmt.Println("user")
        fmt.Println(user)
        users = append(users, user)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating over rows: %v", err)
    }

    return users, nil
}


 
func UpdateUserToken(username, token string) error {
	query := `UPDATE user_data SET token = $1 WHERE username = $2`
	_, err := DB.Exec(query, token, username)
	if err != nil {
		return fmt.Errorf("failed to update user token: %v", err)
	}
	return nil
}


func UpdateAccessToken(username, accessToken string) error {

	fmt.Println(("ACCESS TOKEN IN UPFATE" +accessToken))
	query := "UPDATE user_data SET access_token = $1 WHERE username = $2"
	_, err := DB.Exec(query, accessToken, username)
	return err
}