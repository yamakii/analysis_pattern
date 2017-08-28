package observation

type Person struct {
}

type Observation struct {
	Person
}

type Measurement struct {
	Observation
	PhenomenonType
	Quantity
}

type PhenomenonType struct {
}

type Quantity struct {
	Amount int
	Unit
}

type Unit string

type CategoryObservation struct {
	Observation
	ObservationConcept
}

type Phenomenon struct {
	PhenomenonType
	*ObservationConcept
	Value string
}

type ObservationConcept struct {
	*ObservationConcept
	Value string
}

type Absence struct {
	CategoryObservation
}
type Presence struct {
	CategoryObservation
}

// 使用例
var (
	John     = Person{}
	o        = Observation{John}
	diabetes = ObservationConcept{Value: "糖尿病"}
	// Johnのは糖尿病
	JothnBloodGroup = Presence{CategoryObservation{o, diabetes}}
)
