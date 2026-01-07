package repositories

import "issue_tracker/internal/domain/entities"

type ProductRepository interface {
	GetByID(id entities.IssueID) (*entities.IssueEntity, error)
	GetByCompanyID(companyID entities.CompanyID) (*entities.IssueEntity, error)
	Create(issue *entities.IssueEntity) error
	Update(issue *entities.IssueEntity) error
	Delete(id entities.IssueID) error
}
