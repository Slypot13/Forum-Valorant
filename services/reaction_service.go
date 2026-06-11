package services

import (
	"errors"

	"forum-valorant/models"
	"forum-valorant/repositories"
)

type ReactionService struct {
	reactionRepository *repositories.ReactionRepository
}

func InitReactionService(reactionRepository *repositories.ReactionRepository) *ReactionService {
	return &ReactionService{
		reactionRepository: reactionRepository,
	}
}

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