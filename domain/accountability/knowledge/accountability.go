package knowledge

import "time"

type PartyType int

type Party struct {
	partyType   PartyType
	tel         Telephone
	address     Address
	mailAddress MailAddress
}

type Telephone int
type Address string
type MailAddress string

func (p Party) Type() PartyType {
	return p.partyType
}
func (p Party) Telephone() Telephone {
	return p.tel
}
func (p Party) Address() Address {
	return p.address
}
func (p Party) MailAddress() MailAddress {
	return p.mailAddress
}

type AccountabilityType struct {
	Commissioners PartyTypes
	Responsibles  PartyTypes
}

func (a AccountabilityType) IsSatisfiedWith(commissioner Party, responsible Party) bool {
	return a.Commissioners.In(commissioner.Type()) && a.Responsibles.In(responsible.Type())
}

type PartyTypes []PartyType

func (p PartyTypes) In(target PartyType) bool {
	for _, partyType := range p {
		if partyType == target {
			return true
		}
	}
	return false
}

type TimePeriod struct {
	Start time.Time
	End   time.Time
}

type Accountability struct {
	Type         AccountabilityType
	TimePeriod
	Commissioner Party
	Responsible  Party
	Actions
}

type Action struct {}
type Actions []Action
