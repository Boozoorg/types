package types

type PAN string

type Money int64

type Currency string

const (
	Somoni = "TJS"
	Rubl   = "RUB"
	Dollar = "USD"
	Euro   = "EUR"
)

type Color string

const (
	Platinum = "Star Platinum"
	Silver   = "Silver"
	Gold     = "Gold"
)

type Active bool

const (
	On  = true
	Off = false
)

type Card struct {
	ID       int
	PAN      PAN
	Balance  Money
	Currency Currency
	Color    Color
	Name     string
	Active   bool
}
