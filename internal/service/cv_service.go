package service

import (
	"myself-backend/internal/domain"
	"myself-backend/internal/repository"
)

type CVService struct {
	repo repository.CVRepository
}

func NewCVService(repo repository.CVRepository) *CVService {
	return &CVService{repo: repo}
}

func (s *CVService) CreateCV(cv domain.CV) error {
	return s.repo.Save(cv)
}

// ...other business logic methods...
