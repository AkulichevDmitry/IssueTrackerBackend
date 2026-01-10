package entities

import "issue_tracker/internal/domain/entities/user"

type IssueCommentID int64

type IssueComment struct {
	ID      IssueCommentID
	IssueID IssueID
	UserID  user.ID
	Text    string
}
