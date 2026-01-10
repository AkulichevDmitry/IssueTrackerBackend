package entities

type IssueID int64

type Issue struct {
	ID          IssueID
	ProductID   ProductID
	Title       string
	Description string
	UserID      UserID
	Solved      bool
}
