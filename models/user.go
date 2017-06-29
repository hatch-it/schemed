package models

type User struct {
	schemed.Model

	email    string
	password string
}
