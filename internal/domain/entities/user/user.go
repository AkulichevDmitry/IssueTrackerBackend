package user

import (
	"issue_tracker/internal/domain/entities/company"
)

type ID int64
type Email string

type User struct {
	ID        ID
	CompanyID company.ID
	Username  string
	Email     Email
	Password  string
	Rating    int64
	IsCompany bool
	IsBanned  bool
	Avatar    string
}
