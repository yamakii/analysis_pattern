package basic

type PhenomenonType struct {
}

type Person struct {
}

type Observation struct {
	PhenomenonType
	Person
}

type Measurement struct {
	Observation
	Quantity
}

type Quantity struct {
	Amount int
	Unit
}

type Unit string

type CategoryObservation struct {
	Observation
	Category
}

type Category string

// 使用例
var (
	bloodGroup = PhenomenonType{}
	John = Person{}
	o = Observation{bloodGroup, John}
	// Johnの血液型の測定結果
	JothnBloodGroup  = CategoryObservation{o, Category("A")}
)
