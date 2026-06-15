package services

import (
	"errors"

	"forum-valorant/models"
	"forum-valorant/repositories"
)

type AdminService struct {
	adminRepository *repositories.AdminRepository
}

func InitAdminService(adminRepository *repositories.AdminRepository) *AdminService {
	return &AdminService{
		adminRepository: adminRepository,
	}
}

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

func (s *AdminService) UpdateThreadStatus(threadId int, status string) error {
	if status != "ouvert" && status != "fermé" && status != "archivé" {
		return errors.New("état invalide")
	}

	return s.adminRepository.UpdateThreadStatus(threadId, status)
}

func (s *AdminService) DeleteThread(threadId int) error {
	return s.adminRepository.DeleteThread(threadId)
}

func (s *AdminService) DeleteMessage(messageId int) error {
	return s.adminRepository.DeleteMessage(messageId)
}

func (s *AdminService) BanUser(userId int) error {
	return s.adminRepository.BanUser(userId)
}