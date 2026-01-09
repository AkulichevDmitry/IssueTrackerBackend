package company

import "time"

type Name string
type ID int64

type Company struct {
	ID        ID
	Name      Name
	CreatedAt time.Time
}
