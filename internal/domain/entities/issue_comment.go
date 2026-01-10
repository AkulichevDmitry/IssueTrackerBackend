package entities

type IssueCommentID int64

type IssueComment struct {
	ID      IssueCommentID
	IssueID IssueID
	UserID  UserID
	Text    string
}
