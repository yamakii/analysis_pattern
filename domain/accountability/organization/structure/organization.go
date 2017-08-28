package hierarchy

import (
	"time"
)

type Organization struct {
}

type OrganizationStructureType int

const (
	Sales   OrganizationStructureType = 1
	Service OrganizationStructureType = 2
)

type TimePeriod struct {
	Start time.Time
	End   time.Time
}

type OrganizationStructure struct {
	OrganizationStructureType
	TimePeriod
	Parent     *Organization
	Subsidiary *Organization
}

type OperatingUnit struct {
	Organization
}

type Region struct {
	Organization
}

type Division struct {
	Organization
}

type SalesOffice struct {
	Organization
}

