package hierarchy

import (
	"testing"
	. "github.com/yamakii/analysis_pattern/domain/accountability"
)

func Test(t *testing.T) {
	p := party{}
	unit := NewOperatingUnit(p)
	region := NewRegion(p, unit)
	div := NewDivision(p, region)
	sales := NewSalesOffice(p, div)

	t.Logf("%v", sales)
}
