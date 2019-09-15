package models

//users for login API
type users struct {
	Email    string
	Password string `json:"-"`
}
