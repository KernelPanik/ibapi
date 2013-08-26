package ibapi

type ComboLeg struct {
	ContractId int64
	Ratio      int64
	Action     string
	Exchange   string
}

type UnderComp struct {
	ContractId int64
	Delta      float64
	Price      float64
}

type TagValue struct {
	Tag   string
	Value string
}
