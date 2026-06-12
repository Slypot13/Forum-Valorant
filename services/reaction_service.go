package services

import (
	"errors"

	"forum-valorant/models"
	"forum-valorant/repositories"
)

// contient la logique des réactions.
type ReactionService struct {
	reactionRepository *repositories.ReactionRepository
}

// initialise le service.
func InitReactionService(reactionRepository *repositories.ReactionRepository) *ReactionService {
	return &ReactionService{
		reactionRepository: reactionRepository,
	}
}

// enregistre un like ou dislike.
func (s *ReactionService) ReactToMessage(messageId int, userId int, reactionType string) error {
	if reactionType != "like" && reactionType != "dislike" {
		return errors.New("réaction invalide")
	}

	reaction := models.Reaction{
		MessageId: messageId,
		UserId:    userId,
		Type:      reactionType,
	}

	return s.reactionRepository.SaveReaction(reaction)
}