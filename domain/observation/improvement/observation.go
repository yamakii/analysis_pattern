package improvement

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
	Phenomenon
}

type Phenomenon struct {
	PhenomenonType
	Value string
}

// 使用例
var (
	bloodGroup = PhenomenonType{}
	bloodGroupA = Phenomenon{bloodGroup, "A"}
	John = Person{}
	o = Observation{John}
	// Johnの血液型の測定結果
	JothnBloodGroup  = CategoryObservation{o, bloodGroupA}

	car = Person{}
	oilLevel = PhenomenonType{}
	low = Phenomenon{oilLevel, "low"}
	lowOilLevelCar  = CategoryObservation{Observation{car}, low}
)
