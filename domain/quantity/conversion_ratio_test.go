package quantity

import "testing"

func TestConvert(t *testing.T) {
	q := Quantity{10, feet}

	if q.convertTo(inches) == (Quantity{120, inches}) {
		t.Log("feet * 12 = inces")
	}
}

