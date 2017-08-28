package quantity

type Unit string

const (
	Each Unit = "個"
	Feet Unit = "フィート"
)

type Quantity struct {
	Amount int
	Unit   Unit
}

func (a Quantity) plus(b Quantity) Quantity {
	return Quantity{a.Amount + b.Amount, a.Unit}
}

func (a Quantity) minus(b Quantity) Quantity {
	return Quantity{a.Amount - b.Amount, a.Unit}
}

func (a Quantity) times(b int) Quantity {
	return Quantity{a.Amount * b, a.Unit}
}

func (a Quantity) dividedBy(b int) Quantity {
	return Quantity{a.Amount / b, a.Unit}
}

func (a Quantity) lessThan(b Quantity) bool {
	return a.Amount < b.Amount
}

func (a Quantity) moreThan(b Quantity) bool {
	return a.Amount > b.Amount
}

type SampleQuantity Quantity

func NewSampleQuantity(amount int) SampleQuantity {
	return SampleQuantity{amount, Each}
}

func (a SampleQuantity) lessThan(b SampleQuantity) bool {
	l, r := Quantity(a), Quantity(b)
	return l.lessThan(r)
}

