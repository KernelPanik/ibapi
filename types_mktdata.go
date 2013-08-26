package ibapi

// Inbound Messages
type MsgInTickPrice struct {
	TickerId       int64
	Type           int64
	Price          float64
	Size           int64 `minVer:"2"`
	CanAutoExecute bool  `minVer:"3"`
}

type MsgInTickSize struct {
	TickerId int64
	Type     int64
	Size     int64
}

type MsgInTickOptionComputation struct {
	TickerId    int64
	Type        int64
	ImpliedVol  float64 // > 0
	Delta       float64 // 0 <= delta <= 1
	OptionPrice float64
	PvDividend  float64
	Gamma       float64
	Vega        float64
	Theta       float64
	SpotPrice   float64
}

type MsgInTickGeneric struct {
	TickerId int64
	Type     int64
	Value    float64
}

type MsgInTickString struct {
	TickerId int64
	Type     int64
	Value    string
}

type MsgInTickEFP struct {
	TickerId             int64
	Type                 int64
	BasisPoints          float64
	FormattedBasisPoints string
	ImpliedFuturesPrice  float64
	HoldDays             int64
	FuturesExpiry        string
	DividendImpact       float64
	DividendsToExpiry    float64
}

type MsgInTickSnapshotEnd struct {
	TickerId int64
}

type MsgInMarketDataType struct {
	TickerId int64
	Type     int64
}

// Outbound Messages

type MsgOutReqMktData struct {
	TickerId        int64
	ContractId      int64 `minVer:"37"`
	Symbol          string
	SecurityType    string
	Expiry          string
	Strike          float64
	Right           string
	Multiplier      string
	Exchange        string
	PrimaryExchange string
	Currency        string
	LocalSymbol     string
	ComboLegs       []ComboLeg
	Comp            *UnderComp
	GenericTickList string
	Snapshot        bool
}

func (m *MsgOutReqMktData) RequestEncode(b *requestBytes) {
	b.writeInt(m.TickerId)

	if b.verServer >= 47 {
		b.writeInt(m.ContractId)
	}
	b.writeString(m.Symbol)
	b.writeString(m.SecurityType)
	b.writeString(m.Expiry)
	b.writeFloat(m.Strike)
	b.writeString(m.Right)
	if b.verServer >= 15 {
		b.writeString(m.Multiplier)
	}
	b.writeString(m.Exchange)
	if b.verServer >= 14 {
		b.writeString(m.PrimaryExchange)
	}
	b.writeString(m.Currency)
	if b.verServer >= 2 {
		b.writeString(m.LocalSymbol)
	}

	if b.verServer >= 8 && m.SecurityType == SecTypeBag {
		if m.ComboLegs != nil {
			b.writeInt(int64(len(m.ComboLegs)))
			for _, e := range m.ComboLegs {
				b.writeStruct(e)
			}
		} else {
			b.writeInt(0)
		}
	}

	if b.verServer >= 40 {
		if m.Comp != nil {
			b.writeBool(true)
			b.writeStruct(m.Comp)
		} else {
			b.writeBool(false)
		}
	}

	if b.verServer >= 31 {
		b.writeString(m.GenericTickList)
	}

	if b.verServer >= 35 {
		b.writeBool(m.Snapshot)
	}
}

type MsgOutCxlMktData struct {
	TickerId int64
}

type MsgOutCalcImpVol struct {
}

type MsgOutCxlCalcImpVol struct {
}

type MsgOutCalcOptPx struct {
}

type MsgOutCxlCalcOptPx struct {
}

type MsgOutReqMktDataType struct {
	Type int64
}
