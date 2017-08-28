package basic

import "testing"

func TestEmployment(t *testing.T) {
	employment := AccountabilityType(1)
	company := Organization{}
	smith := Person{}
	period := TimePeriod{}

	// companyはsmithに対して雇用(employment)の責任関係を持つ
	t.Log(Accountability{employment, period, &Party(company), &Party(smith)})
}

func TestManager(t *testing.T) {
	manager := AccountabilityType(2)
	serviceTeam := Organization{}
	smith := Person{}
	period := TimePeriod{}

	// smithはservice teamのmanager
	// service teamはsmithに対してmanagerの責任関係を持つ
	t.Log(Accountability{manager, period, &Party(serviceTeam), &Party(smith)})
}

func TestContract(t *testing.T) {
	contract := AccountabilityType(3)
	service := Organization{}
	smith := Person{}
	period := TimePeriod{}

	// smithはserviceの契約をした
	// serviceはsmithに対してcontractの責任関係を持つ
	t.Log(Accountability{contract, period, &Party(service), &Party(smith)})
}

func TestRegion(t *testing.T) {
	contract := AccountabilityType(3)
	service := Organization{}
	smith := Person{}
	period := TimePeriod{}

	// smithはserviceの契約をした
	// serviceはsmithに対してcontractの責任関係を持つ
	t.Log(Accountability{contract, period, &Party(service), &Party(smith)})
}
