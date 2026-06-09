package services

import (
	"errors"

	"forum-valorant/models"
	"forum-valorant/repositories"
)

type MessageService struct {
	messageRepository *repositories.MessageRepository
	threadRepository  *repositories.ThreadRepository
}

func InitMessageService(messageRepository *repositories.MessageRepository, threadRepository *repositories.ThreadRepository) *MessageService {
	return &MessageService{
		messageRepository: messageRepository,
		threadRepository:  threadRepository,
	}
}

func (s *MessageService) CreateMessage(content string, threadId int, userId int) error {
	if content == "" {
		return errors.New("le message ne peut pas être vide")
	}

	thread, err := s.threadRepository.ReadById(threadId)

	if err != nil {
		return errors.New("le fil de discussion n'existe pas")
	}

	if thread.Status != "ouvert" {
		return errors.New("ce fil n'accepte plus de nouveaux messages")
	}

	message := models.Message{
		Content:  content,
		ThreadId: threadId,
		UserId:   userId,
	}

	return s.messageRepository.CreateMessage(message)
}

func (s *MessageService) GetMessagesByThreadId(threadId int) ([]models.Message, error) {
	return s.messageRepository.ReadByThreadId(threadId)
}