package repositories

import "issue_tracker/internal/domain/entities"

type ProductRepository interface {
	GetByID(id entities.IssueID) (*entities.Issue, error)
	GetByCompanyID(companyID entities.CompanyID) (*entities.Issue, error)
	Create(issue *entities.Issue) error
	Update(issue *entities.Issue) error
	Delete(id entities.IssueID) error
}
