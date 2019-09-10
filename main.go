//main.go

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

var db *sql.DB

//Credentials have information [email] & [password] for storing to database
type Credentials struct {
	email    string
	password string
}

//users for login API
type users struct {
	Email    string
	Password string `json:"-"`
}

func main() {

	ConnectDB()

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "7000"
	}

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", HomeHandler).Methods("GET")

	myRouter.HandleFunc("/login", LoginHandler).Methods("POST")
	myRouter.HandleFunc("/signup", SignupHandler).Methods("POST")

	log.Printf("Starting server on port %s \n", port)

	if err := http.ListenAndServe(":"+port, myRouter); err != nil {
		log.Fatalf("Could not start server: %s \n", err.Error())
	}
}

//HomeHandler link to HomePage API
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello World")
}

//ConnectDB open connection to database
func ConnectDB() {
	var err error

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))

	err = db.Ping()

	if err != nil {
		log.Fatal("Error: Could not establish connection to databse")
	}
}

/*LoginHandler get user's information for logging in
Return result user exist or not */
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	r.ParseForm()

	var Guest Credentials
	Guest.email, Guest.password = r.FormValue("email"), r.FormValue("password")

	result := FindUser(Guest.email, Guest.password)

	if result == 0 {
		w.WriteHeader(http.StatusUnauthorized)

		fmt.Fprintln(w, "Invalid email or password")
		return
	}

	User := users{}

	User.Email = Guest.email
	User.Password = Guest.password

	MyUser, err := json.Marshal(User)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(MyUser)
}

//SignupHandler get guest's information for Signing up to API
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	r.ParseForm()

	var Guest Credentials
	Guest.email, Guest.password = r.FormValue("email"), r.FormValue("password")

	if Guest.email == "" || Guest.password == "" {

		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(w, "Email or Password is not valid")

		return
	}

	result := getUser(Guest.email)

	if result == 1 {
		w.WriteHeader(http.StatusUnauthorized)

		fmt.Fprintln(w, "This Email has been Registered")

		return
	}

	AddUser(Guest.email, Guest.password)
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "SUCCESS")

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
