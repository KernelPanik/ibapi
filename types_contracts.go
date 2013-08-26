package ibapi

type MsgInContractData struct {
	ReqId           int64
	Symbol          string
	SecurityType    string
	Expiry          string
	Strike          float64
	Right           string
	Exchange        string
	Currency        string
	LocalSymbol     string
	MarketName      string
	TradingClass    string
	ContractId      int64
	MinTick         float64
	Multiplier      string
	OrderTypes      string
	ValidExchanges  string
	PriceMagnifier  int64
	SpotContractId  int64
	LongName        string
	PrimaryExchange string
	ContractMonth   string
	Industry        string
	Category        string
	Subcategory     string
	TimezoneId      string
	TradingHours    string
	LiquidHours     string
	EvRule          string
	EvMultiplier    float64
	SecIds          []TagValue
}

type MsgInBondContractData struct {
	ReqId             int64
	Symbol            string
	SecType           string
	Cusip             string
	Coupon            float64
	Maturity          string
	IssueDate         string
	Ratings           string
	BondType          string
	CouponType        string
	Convertible       bool
	Callable          bool
	Putable           bool
	DescAppend        string
	Exchange          string
	Currency          string
	MarketName        string
	TradingClass      string
	ContractId        int64
	MinTick           float64
	OrderTypes        string
	ValidExchanges    string
	NextOptionDate    string
	NextOptionType    string
	NextOptionPartial bool
	Notes             string
	LongName          string
	EvRule            string
	EvMultiplier      float64
	SecIds            []TagValue
}

type MsgInContractDataEnd struct {
	ReqId int64
}

type MsgOutReqContractData struct {
	ReqId          int64
	ContractId     int64
	Symbol         string
	SecurityType   string
	Expiry         string
	Strike         float64
	Right          string
	Multiplier     string
	Exchange       string
	Currency       string
	LocalSymbol    string
	IncludeExpired bool
	SecurityIdType string
	SecurityId     string
}
