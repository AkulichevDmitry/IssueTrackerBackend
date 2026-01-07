package entities

import (
	"time"
)

type CompanyID int64
type CompanyName string

type Company struct {
	ID        CompanyID
	Name      CompanyName
	CreatedAt time.Time
}
