package entities

import (
	"time"
)

type IssueID int64

type IssueEntity struct {
	ID          IssueID
	ProductID   ProductID
	Title       string
	Description string
	UserID      UserID
	Solved      bool
	CreatedAt   time.Time
}
