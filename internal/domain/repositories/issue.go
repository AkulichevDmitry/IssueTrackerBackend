package repositories

import (
	"issue_tracker/internal/domain/entities"
)

type IssueRepository interface {
	GetByID(id entities.IssueID) (*entities.IssueEntity, error)
	GetByProductID(productID entities.ProductID) ([]*entities.IssueEntity, error)
	GetByCompanyID(companyID entities.CompanyID) ([]*entities.IssueEntity, error)
	Create(issue *entities.IssueEntity) error
	Update(issue *entities.IssueEntity) error
	Delete(id entities.IssueID) error
}
