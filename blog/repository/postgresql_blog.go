package repository

import (
	""
)

//ConnectDB open connection to database
func ConnectDB() {
	var err error

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))

	err = db.Ping()

	if err != nil {
		log.Fatal("Error: Could not establish connection to databse")
	}
}

//FindUser Looking for user' information storing on database
func FindUser(email, password string) int64 {

	QueryStmt := `
	SELECT * FROM users
	WHERE email=$1 AND password=$2
	`

	result, err := db.Exec(QueryStmt, email, password)

	if err != nil {
		panic(err)
	}

	rows, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	return rows
}

func getUser(email string) int64 {

	QueryStmt := `
	SELECT * FROM users
	WHERE email=$1
	`
	result, err := db.Exec(QueryStmt, email)

	if err != nil {
		panic(err)
	}

	rows, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	return rows
}

//AddUser make a SQL' query for Storing user's information to database
func AddUser(email, password string) {

	QueryStmt := `
	INSERT INTO users(email, password)
	VALUES($1,$2)
	`

	_, err := db.Exec(QueryStmt, email, password)

	if err != nil {
		panic(err)
	}
}