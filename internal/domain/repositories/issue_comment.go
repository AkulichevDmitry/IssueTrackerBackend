package repositories

import "issue_tracker/internal/domain/entities"

type IssueCommentRepository interface {
	GetByID(id entities.IssueCommentID) (*entities.IssueComment, error)
	Create(comment *entities.IssueComment) error
	Update(id entities.IssueCommentID) error
	Delete(id entities.IssueCommentID) error
}
