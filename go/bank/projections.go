package bank

type Count struct {
	EventCount    int
}

func NewCount() *Count {
	return &Count{}
}

func (c *Count) Accept(e Event) {
	switch e.(type) {
	case *MoneyWasDeposited:

	case *MoneyWasWithdrawn:
	}
}
