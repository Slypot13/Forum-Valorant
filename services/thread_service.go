package services

import (
	"errors"
	"forum-valorant/models"
	"forum-valorant/repositories"
)

type ThreadService struct {
	threadRepository *repositories.ThreadRepository
}

func InitThreadService(threadRepository *repositories.ThreadRepository) *ThreadService {
	return &ThreadService{
		threadRepository: threadRepository,
	}
}

func (s *ThreadService) CreateThread(title string, content string, userId int) error {
	if title == "" {
		return errors.New("le titre est obligatoire")
	}

	if content == "" {
		return errors.New("le contenu est obligatoire")
	}

	thread := models.Thread{
		Title:   title,
		Content: content,
		Status:  "ouvert",
		UserId:  userId,
	}

	_, err := s.threadRepository.CreateThread(thread)

	return err
}