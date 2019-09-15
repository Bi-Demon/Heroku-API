//main.go

package main

import (
	"log"
	"net/http"
	"os"

	blog "github.com/Bi-Demon/Heroku-API/blog/delivery/http"
	repo "github.com/Bi-Demon/Heroku-API/blog/repository"
	"github.com/gorilla/mux"
)

func main() {

	repo.ConnectDB()

	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "7000"
	}

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", blog.HomeHandler).Methods("GET")

	myRouter.HandleFunc("/login", blog.LoginHandler).Methods("POST")
	myRouter.HandleFunc("/signup", blog.SignupHandler).Methods("POST")

	log.Printf("Starting server on port %s \n", port)

	if err := http.ListenAndServe(":"+port, myRouter); err != nil {
		log.Fatalf("Could not start server: %s \n", err.Error())
	}
}
