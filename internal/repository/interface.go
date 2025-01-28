package repository

import "myself-backend/internal/domain"

type CVRepository interface {
	Save(cv domain.CV) error
	FindById(id int) (*domain.CV, error)
	// ...other methods...
}
