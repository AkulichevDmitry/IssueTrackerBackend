package entities

import (
	"issue_tracker/internal/domain/entities/company"
)

type UserID int64
type UserEmail string

type User struct {
	ID        UserID
	CompanyID company.ID
	Username  string
	Email     UserEmail
	Password  string
	Rating    int64
	IsCompany bool
	IsBanned  bool
	Avatar    string
}
