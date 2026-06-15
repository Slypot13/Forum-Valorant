package models

// représente une réaction (like/dislike) d'un utilisateur sur un message.
type Reaction struct {
	Id        int
	MessageId int
	UserId    int
	Type      string
}