package http

import (
	"github.com/Bi-Demon/Heroku-API/blog"
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
