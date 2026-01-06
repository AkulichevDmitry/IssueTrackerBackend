package repositories

import (
	"issue_tracker/internal/domain/entities"
)

type CompanyRepository interface {
	GetByID(id entities.CompanyID) (*entities.CompanyEntity, error)
	Create(company *entities.CompanyEntity) error
	Update(company *entities.CompanyEntity) error
	Delete(id entities.CompanyID) error
}
