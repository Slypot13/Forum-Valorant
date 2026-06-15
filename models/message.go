package models

type Message struct {
	Id        int
	Content   string
	ThreadId  int
	UserId    int
	Username  string
	CreatedAt string
	Score     int
}