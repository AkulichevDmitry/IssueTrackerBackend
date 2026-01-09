package repositories

import (
	"issue_tracker/internal/domain/entities/company"
)

type CompanyRepository interface {
	GetByID(id company.ID) (*company.Company, error)
	Create(company *company.Company) error
	Update(company *company.Company) error
	Delete(id company.ID) error
}
