package models

// représente un fil de discussion (sujet).
type Thread struct {
	Id        int
	Title     string
	Content   string
	Status    string
	UserId    int
	Username  string
	CreatedAt string
}