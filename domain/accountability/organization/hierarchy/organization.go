package hierarchy

import a "github.com/yamakii/analysis_pattern/domain/accountability"

type Organization struct {
	a.party
	Parent *Organization
}

type OperatingUnit struct {
	Organization
}

func NewOperatingUnit(party a.party) OperatingUnit {
	return OperatingUnit{Organization{party, nil}}
}

type Region struct {
	Organization
}

func NewRegion(party a.party, unit OperatingUnit) Region {
	return Region{Organization{party, &unit.Organization}}
}

type Division struct {
	Organization
}

func NewDivision(party a.party, region Region) Division {
	return Division{Organization{party, &region.Organization}}
}

type SalesOffice struct {
	Organization
}

func NewSalesOffice(party a.party, division Division) SalesOffice {
	return SalesOffice{Organization{party, &division.Organization}}
}
