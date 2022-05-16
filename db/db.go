package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Db *sql.DB

const (
	host     = "localhost"
	port     = 5444
	user     = "abubakirkais"
	password = ""
	dbname   = "abubakirkais"
)

func SetpUpDbConnection() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	Db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// insert tables
	insertUsersTableCmd := `CREATE TABLE IF NOT EXISTS users (
		id serial PRIMARY KEY,
		uid UUID UNIQUE NOT NULL,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		profile_image_url VARCHAR(255),
		profile_public_id VARCHAR(255),
		active BOOL DEFAULT true,
		verified_at TIMESTAMP WITH TIME ZONE,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE,
		deleted_at TIMESTAMP WITH TIME ZONE
	 	);`

	// insert users table
	_, err = Db.Exec(insertUsersTableCmd)
	CheckError(err)

	// close database
	defer Db.Close()

	// check db
	err = Db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func GetDb() *sql.DB {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	Db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	return Db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
