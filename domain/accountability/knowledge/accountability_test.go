package knowledge

import "testing"

func TestRegion(t *testing.T) {
	const (
		operationUnit PartyType = PartyType(1)
		region        PartyType = PartyType(2)
		division      PartyType = PartyType(3)
		salesOffice   PartyType = PartyType(4)
	)
	// commissionerになれるPartyType
	commissioners := PartyTypes{operationUnit, region, division}
	// responsibleになれるPartyType
	responsibles := PartyTypes{region, division, salesOffice}

	var (
		salesOperation   Party = Party{partyType: operationUnit}
		productOperation Party = Party{partyType: operationUnit}
		kanto            Party = Party{partyType: region}
		todaSalesOffice  Party = Party{partyType: salesOffice}
	)

	regionStruct := AccountabilityType{commissioners, responsibles}

	if regionStruct.IsSatisfiedWith(salesOperation, kanto) {
		t.Log("salesOperationはcommisionerになれる, kantoはresponsibleになれる")
	}

	if !regionStruct.IsSatisfiedWith(todaSalesOffice, kanto) {
		t.Log("todaSalesOfficeはcommisionerになれない")
	}

	if !regionStruct.IsSatisfiedWith(kanto, productOperation) {
		t.Log("productOperationはresponsibleになれない")
	}

	// この構造では階層構造などの制約は実現できない(もっと詳細なモデルが必要)
}
