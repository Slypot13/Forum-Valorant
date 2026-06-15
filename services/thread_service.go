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
func (s *ThreadService) CreateThread(title string, content string, userId int, tagId int) error {
	if title == "" {
		return errors.New("le titre est obligatoire")
	}

	if content == "" {
		return errors.New("le contenu est obligatoire")
	}

	if tagId <= 0 {
		return errors.New("la catégorie est obligatoire")
	}

	thread := models.Thread{
		Title:   title,
		Content: content,
		Status:  "ouvert",
		UserId:  userId,
	}

	threadId, err := s.threadRepository.CreateThread(thread)

	if err != nil {
		return err
	}

	return s.threadRepository.AddTagToThread(threadId, tagId)
}

// liste les sujets actifs.
func (s *ThreadService) GetVisibleThreads() ([]models.Thread, error) {
	return s.threadRepository.ReadVisibleThreads()
}

// liste les sujets actifs avec pagination.
func (s *ThreadService) GetVisibleThreadsPaginated(limit int, offset int) ([]models.Thread, error) {
	return s.threadRepository.ReadVisibleThreadsPaginated(limit, offset)
}

// liste les sujets d'une catégorie avec pagination.
func (s *ThreadService) GetVisibleThreadsByTagPaginated(tagId int, limit int, offset int) ([]models.Thread, error) {
	return s.threadRepository.ReadVisibleThreadsByTagPaginated(tagId, limit, offset)
}

// retourne un sujet par son ID.
func (s *ThreadService) GetThreadById(id int) (models.Thread, error) {
	return s.threadRepository.ReadById(id)
}

// vérifie les droits et modifie le sujet.
// Pour l'instant, seul le propriétaire peut modifier son sujet.
// La logique admin sera ajoutée plus tard dans la partie administration.
func (s *ThreadService) UpdateThread(id int, title string, content string, userId int, role string) error {
	thread, err := s.threadRepository.ReadById(id)

	if err != nil {
		return errors.New("fil introuvable")
	}

	if thread.UserId != userId {
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
// Pour l'instant, seul le propriétaire peut supprimer son sujet.
// La logique admin sera ajoutée plus tard dans la partie administration.
func (s *ThreadService) DeleteThread(id int, userId int, role string) error {
	thread, err := s.threadRepository.ReadById(id)

	if err != nil {
		return errors.New("fil introuvable")
	}

	if thread.UserId != userId {
		return errors.New("vous n'avez pas le droit de supprimer ce fil")
	}

	return s.threadRepository.DeleteThread(id)
}

// recherche des sujets par titre ou catégorie.
func (s *ThreadService) SearchVisibleThreads(search string, limit int, offset int) ([]models.Thread, error) {
	if search == "" {
		return s.threadRepository.ReadVisibleThreadsPaginated(limit, offset)
	}

	return s.threadRepository.SearchVisibleThreads(search, limit, offset)
}