package measurement

import q "github.com/yamakii/analysis_pattern/domain/quantity"

type PhenomenonType struct {
}

type Person struct {
}

type Measurement struct {
	PhenomenonType
	Person
	q.Quantity
}

// 使用例
var (
	John = Person{}
	height = PhenomenonType{}
	// Johnのheightの測定結果
	JothnHeight = Measurement{height, John, q.Quantity{6, q.Feet}}
)
