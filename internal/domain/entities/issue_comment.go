package entities

import (
	"time"
)

type IssueCommentID int64

type IssueCommentEntity struct {
	ID        IssueCommentID
	IssueID   IssueID
	UserID    UserID
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
