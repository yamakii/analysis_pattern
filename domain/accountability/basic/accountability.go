package basic

import "time"

type Party interface {
	Telephone() Telephone
	Address() Address
	MailAddress() MailAddress
}

type PartyType int

type party struct {
	tel         Telephone
	address     Address
	mailAddress MailAddress
}

type Telephone int
type Address string
type MailAddress string

func (p party) Telephone() Telephone {
	return p.tel
}
func (p party) Address() Address {
	return p.address
}
func (p party) MailAddress() MailAddress {
	return p.mailAddress
}

type Person struct {
	party
}
type Organization struct {
	party
}

type AccountabilityType int

type TimePeriod struct {
	Start time.Time
	End   time.Time
}

type Accountability struct {
	AccountabilityType
	TimePeriod
	Commissioner *Party
	Responsible  *Party
}
