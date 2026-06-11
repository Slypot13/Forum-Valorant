package models

type Reaction struct {
	Id        int
	MessageId int
	UserId    int
	Type      string
}