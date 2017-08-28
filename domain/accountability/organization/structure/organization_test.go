package hierarchy

import (
	"testing"
)

func Test(t *testing.T) {
	operationUnitStructure := OrganizationStructure{Sales, TimePeriod{}, nil, &Region{}.Organization}
	t.Log(operationUnitStructure)
	regionStructure := OrganizationStructure{Sales, TimePeriod{}, &OperatingUnit{}.Organization, &Division{}.Organization}
	t.Log(regionStructure)
	service := OrganizationStructure{Service, TimePeriod{}, nil, &Region{}.Organization}
	t.Log(service)
}
