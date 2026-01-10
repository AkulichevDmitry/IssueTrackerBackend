package entities

import "issue_tracker/internal/domain/entities/user"

type IssueID int64

type Issue struct {
	ID          IssueID
	ProductID   ProductID
	Title       string
	Description string
	UserID      user.ID
	Solved      bool
}
