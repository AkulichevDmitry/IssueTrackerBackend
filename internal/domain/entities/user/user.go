package user

import (
	"issue_tracker/internal/domain/entities/company"
)

type User struct {
	ID        ID
	CompanyID company.ID
	Username  Username
	Email     Email
	Password  Password
	Rating    Rating
	IsCompany bool
	IsBanned  bool
	Avatar    Avatar
}
