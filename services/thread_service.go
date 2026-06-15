package services

import (
	"errors"

	"forum-valorant/models"
	"forum-valorant/repositories"
)

// contient la logique des sujets.
type ThreadService struct {
	threadRepository *repositories.ThreadRepository
}

// initialise le service.
func InitThreadService(threadRepository *repositories.ThreadRepository) *ThreadService {
	return &ThreadService{
		threadRepository: threadRepository,
	}
}

// valide et crée un sujet.
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

// liste les sujets actifs.
func (s *ThreadService) GetVisibleThreads() ([]models.Thread, error) {
	return s.threadRepository.ReadVisibleThreads()
}

// liste les sujets actifs avec pagination.
func (s *ThreadService) GetVisibleThreadsPaginated(limit int, offset int) ([]models.Thread, error) {
	return s.threadRepository.ReadVisibleThreadsPaginated(limit, offset)
}

// retourne un sujet par son ID.
func (s *ThreadService) GetThreadById(id int) (models.Thread, error) {
	return s.threadRepository.ReadById(id)
}

// vérifie les droits et modifie le sujet.
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

// vérifie les droits et supprime le sujet.
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