package entities

type Payment struct {
	ID string
	Currency string
	AmountPaid int64
	Metadata map[string]any
}
