package models

type Tokens struct {
	ID             uint
	CaseID         uint
	Type           string
	Symbol         string
	Name           string
	Price          float64
	IssuerNumber   uint
	CompanyArea    string
	CompanyCapital float64
	Description    string
}
