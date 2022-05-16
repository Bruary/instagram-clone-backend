package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "instagram"
)

func SetpUpDbConnection() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// insert tables
	insertUsersTableCmd := `CREATE TABLE IF NOT EXISTS users (
		id uuid NOT NULL,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		profile_image_url VARCHAR(255),
		profile_public_id VARCHAR(255),
		active BOOL DEFAULT true,
		verified_at TIMESTAMP WITH TIME ZONE,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP Now(),
		updated_at TIMESTAMP WITH TIME ZONE,
		deleted_at TIMESTAMP WITH TIME ZONE
	 );`

	// insert users table
	db.Exec(insertUsersTableCmd)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
