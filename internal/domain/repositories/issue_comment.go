package repositories

import "issue_tracker/internal/domain/entities"

type IssueCommentRepository interface {
	GetByID(id entities.IssueCommentID) (*entities.IssueCommentEntity, error)
	Create(comment *entities.IssueCommentEntity) error
	Update(id entities.IssueCommentID) error
	Delete(id entities.IssueCommentID) error
}
