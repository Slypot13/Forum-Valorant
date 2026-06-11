package models

type Message struct {
	Id        int
	Content   string
	ThreadId  int
	UserId    int
	CreatedAt string
	Score     int
}