package entities

import "issue_tracker/internal/domain/entities/company"

type ProductID int64
type ProductName string
type ProductDescription string

type Product struct {
	ID          ProductID
	CompanyID   company.ID
	Name        ProductName
	Description ProductDescription
}
