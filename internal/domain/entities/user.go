package entities

import (
	"time"
)

type UserID int64
type UserEmail string

type UserEntity struct {
	ID        UserID
	CompanyID CompanyID
	Username  string
	Email     UserEmail
	Password  string
	Rating    int64
	IsCompany bool
	IsBanned  bool
	Avatar    string
	CreatedAt time.Time
}
