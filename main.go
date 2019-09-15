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
