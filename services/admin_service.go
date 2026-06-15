package services

import (
	"errors"

	"forum-valorant/models"
	"forum-valorant/repositories"
)

// gère la logique métier de l'administration.
type AdminService struct {
	adminRepository *repositories.AdminRepository
}

// initialise le service d'administration.
func InitAdminService(adminRepository *repositories.AdminRepository) *AdminService {
	return &AdminService{
		adminRepository: adminRepository,
	}
}

// récupère l'ensemble des données pour le tableau de bord (sujets, messages, utilisateurs).
func (s *AdminService) GetDashboardData() ([]models.Thread, []models.Message, []models.User, error) {
	threads, err := s.adminRepository.ReadAllThreads()

	if err != nil {
		return nil, nil, nil, err
	}

	messages, err := s.adminRepository.ReadAllMessages()

	if err != nil {
		return nil, nil, nil, err
	}

	users, err := s.adminRepository.ReadAllUsers()

	if err != nil {
		return nil, nil, nil, err
	}

	return threads, messages, users, nil
}

// valide et met à jour le statut d'un sujet.
func (s *AdminService) UpdateThreadStatus(threadId int, status string) error {
	if status != "ouvert" && status != "fermé" && status != "archivé" {
		return errors.New("état invalide")
	}

	return s.adminRepository.UpdateThreadStatus(threadId, status)
}

// supprime un sujet.
func (s *AdminService) DeleteThread(threadId int) error {
	return s.adminRepository.DeleteThread(threadId)
}

// supprime un message.
func (s *AdminService) DeleteMessage(messageId int) error {
	return s.adminRepository.DeleteMessage(messageId)
}

// bannit un utilisateur.
func (s *AdminService) BanUser(userId int) error {
	return s.adminRepository.BanUser(userId)
}