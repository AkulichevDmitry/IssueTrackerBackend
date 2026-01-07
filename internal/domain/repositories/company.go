package repositories

import (
	"issue_tracker/internal/domain/entities"
)

type CompanyRepository interface {
	GetByID(id entities.CompanyID) (*entities.Company, error)
	Create(company *entities.Company) error
	Update(company *entities.Company) error
	Delete(id entities.CompanyID) error
}
