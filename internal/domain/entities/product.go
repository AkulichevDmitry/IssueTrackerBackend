package entities

type ProductID int64
type ProductName string
type ProductDescription string

type Product struct {
	ID          ProductID
	CompanyID   CompanyID
	Name        ProductName
	Description ProductDescription
}
