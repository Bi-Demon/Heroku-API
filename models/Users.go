package models

//Users for login API
type Users struct {
	Email    string
	Password string `json:"-"`
}
