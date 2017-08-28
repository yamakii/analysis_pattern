package associated

type ObservationConcept struct {
	name string
}

type AssociativeFunctions struct {
	arguments []ObservationConcept
	product   ObservationConcept
}

type Observation struct {
	ObservationConcept
}

type AssociatedObservation struct {
	Observation
	AssociativeFunctions
	evidence []Observation
}

// 使用例
var (
	// 知識レベル
	weightLoss = ObservationConcept{"体重減少"}
	thirst     = ObservationConcept{"喉の渇き"}
	polyuria   = ObservationConcept{"多尿症"}
	diabetes   = ObservationConcept{"糖尿病"}
	fun = AssociativeFunctions{[]ObservationConcept{weightLoss, thirst, polyuria}, diabetes}

	// 操作レベル
	evi = []Observation{{weightLoss}, {thirst}, {polyuria}}
	// 糖尿病の連想観測
	predDiabetes = AssociatedObservation{
		Observation{diabetes}, fun, evi}
)
