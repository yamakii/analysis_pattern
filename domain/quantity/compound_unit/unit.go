package compound_unit

type Unit struct {
}

type AtomicUnit struct {
	Unit
}

type CompoundUnitBasic struct {
	Unit
}

type UnitReference struct {
	*AtomicUnit
	CompoundUnitBasic
	Power int
}

// こんな感じで使う
var (
	meter             = AtomicUnit{}
	square            = CompoundUnitBasic{}
	squareRef         = UnitReference{&meter, square, 2}
	second            = AtomicUnit{}
	perSecond         = CompoundUnitBasic{}
	perSecondRef      = UnitReference{&second, perSecond, -1}
	meterPerSecond    = CompoundUnitBasic{}
	meterPerSecondRef = []UnitReference{
		UnitReference{&meter, meterPerSecond, 1},
		UnitReference{&second, meterPerSecond, -1}}
)

type CompoundUnit struct {
	Unit
	direct  []AtomicUnit
	inverse []AtomicUnit
}

// こんな感じで使う
var (
	gravity            = CompoundUnit{Unit{}, []AtomicUnit{meter}, []AtomicUnit{second, second}}
)
