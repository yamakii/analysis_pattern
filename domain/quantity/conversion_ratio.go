package quantity

type ConversionMap map[Unit]map[Unit]ConversionRatio
type ConversionRatio struct {
	From  Unit
	To    Unit
	Ratio int
}

const (
	feet   Unit = "feet"
	inches Unit = "inches"
)

var feetToInches ConversionRatio = ConversionRatio{feet, inches, 12}

func NewConversionMap() ConversionMap {
	result := make(ConversionMap)
	result[feet] = map[Unit]ConversionRatio{inches: feetToInches}
	return result
}

func (q Quantity) convertTo(unit Unit) Quantity {
	ratio := NewConversionMap()[q.Unit][unit]
	return Quantity{q.Amount * ratio.Ratio, unit}
}
