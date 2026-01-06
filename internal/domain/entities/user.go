package entities

import (
	"time"
)

type UserID int64

type UserEntity struct {
	ID        UserID
	CompanyID CompanyID
	Username  string
	Password  string
	Rating    int64
	IsCompany bool
	IsBanned  bool
	Avatar    string
	CreatedAt time.Time
}
