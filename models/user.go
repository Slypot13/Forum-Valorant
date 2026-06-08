package models

type User struct {
	Id        int
	Username  string
	Email     string
	Password  string
	Role      string
	IsBanned  bool
	CreatedAt string
}