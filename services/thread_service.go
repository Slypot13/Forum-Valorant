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

func (s *ThreadService) GetVisibleThreads() ([]models.Thread, error) {
	return s.threadRepository.ReadVisibleThreads()
}

func (s *ThreadService) GetThreadById(id int) (models.Thread, error) {
	return s.threadRepository.ReadById(id)
}

func (s *ThreadService) UpdateThread(id int, title string, content string, userId int, role string) error {
	thread, err := s.threadRepository.ReadById(id)

	if err != nil {
		return errors.New("fil introuvable")
	}

	if thread.UserId != userId && role != "admin" {
		return errors.New("vous n'avez pas le droit de modifier ce fil")
	}

	if title == "" || content == "" {
		return errors.New("le titre et le contenu sont obligatoires")
	}

	thread.Title = title
	thread.Content = content

	return s.threadRepository.UpdateThread(thread)
}

func (s *ThreadService) DeleteThread(id int, userId int, role string) error {
	thread, err := s.threadRepository.ReadById(id)

	if err != nil {
		return errors.New("fil introuvable")
	}

	if thread.UserId != userId && role != "admin" {
		return errors.New("vous n'avez pas le droit de supprimer ce fil")
	}

	return s.threadRepository.DeleteThread(id)
}