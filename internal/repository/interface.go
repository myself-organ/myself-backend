package repository

import "myself-backend/internal/domain"

type CVRepository interface {
	Save(cv domain.CV) error
	FindByID(id int) (*domain.CV, error)
	// ...other methods...
}

type CV struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Address string
}
