package quantity

import (
	"testing"
)

func Test(t *testing.T) {
	a, b := Quantity{6, Each}, Quantity{5, Each}

	t.Logf("%v + %v = %v \n", a, b, a.plus(b))
	t.Logf("%v - %v = %v \n", a, b, a.minus(b))
	t.Logf("%v * %v = %v \n", a, 3, a.times(3))
	t.Logf("%v / %v = %v \n", a, 3, a.dividedBy(3))
	t.Logf("%v < %v = %v \n", a, b, a.lessThan(b))
	t.Logf("%v > %v = %v \n", a, b, a.moreThan(b))
}
