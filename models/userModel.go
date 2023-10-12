package models

type User struct {
	Id       int    `json:id`
	Name     string `json:name`
	Username string `json:username`
	Email    string `json:email`
}

type UserList struct {
	Users []User `json:users`
}

type NewUser struct {
	Name     string `json:name`
	Username string `json:username`
	Email    string `json:email`
}
