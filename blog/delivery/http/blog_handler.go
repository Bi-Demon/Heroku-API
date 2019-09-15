package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	repo "github.com/Bi-Demon/Heroku-API/blog/repository"
	"github.com/Bi-Demon/Heroku-API/models"
)

//HomeHandler link to HomePage API
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello World")
}

/*LoginHandler get user's information for logging in
Return result user exist or not */
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	r.ParseForm()

	var Guest models.Credentials

	Guest.Email, Guest.Password = r.FormValue("email"), r.FormValue("password")

	result := repo.FindUser(Guest.Email, Guest.Password)

	if result == 0 {
		w.WriteHeader(http.StatusUnauthorized)

		fmt.Fprintln(w, "Invalid email or password")
		return
	}

	User := models.Users{}

	User.Email = Guest.Email
	User.Password = Guest.Password

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

	var Guest models.Credentials
	Guest.Email, Guest.Password = r.FormValue("email"), r.FormValue("password")

	if Guest.Email == "" || Guest.Password == "" {

		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprintln(w, "Email or Password is not valid")

		return
	}

	result := repo.GetUser(Guest.Email)

	if result == 1 {
		w.WriteHeader(http.StatusUnauthorized)

		fmt.Fprintln(w, "This Email has been Registered")

		return
	}

	repo.AddUser(Guest.Email, Guest.Password)
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "SUCCESS")

}
