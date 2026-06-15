package models

// représente un message posté dans un fil de discussion.
type Message struct {
	Id        int
	Content   string
	ThreadId  int
	UserId    int
	Username  string
	CreatedAt string
	Score     int
}