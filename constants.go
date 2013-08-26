package ibapi

const (
	VerClient    = 60
	VerMinServer = 38
	SecTypeBag   = "BAG"
)

type FaMsgType int64

const (
	FaMsgTypeGroups FaMsgType = iota
	FaMsgTypeProfiles
	FaMsgTypeAliases
)

type LegOpenClose int64

const (
	LogOpenClosePosSame LegOpenClose = iota
	LogOpenClosePosOpen
	LogOpenClosePosClose
	LogOpenClosePosUnknown
)

func (t FaMsgType) String() string {
	switch t {
	case FaMsgTypeGroups:
		return "GROUPS"
	case FaMsgTypeProfiles:
		return "PROFILES"
	case FaMsgTypeAliases:
		return "ALIASES"
	default:
		return ""
	}
}

const (
	// incoming msg ids
	mInTickPrice              = 1
	mInTickSize               = 2
	mInOrderStatus            = 3
	mInErrorMessage           = 4
	mInOpenOrder              = 5
	mInAccountValue           = 6
	mInPortfolioValue         = 7
	mInAccountUpdateTime      = 8
	mInNextValidId            = 9
	mInContractData           = 10
	mInExecutionData          = 11
	mInMarketDepth            = 12
	mInMarketDepthL2          = 13
	mInNewsBulletins          = 14
	mInManagedAccounts        = 15
	mInReceiveFA              = 16
	mInHistoricalData         = 17
	mInBondContractData       = 18
	mInScannerParameters      = 19
	mInScannerData            = 20
	mInTickOptionComputation  = 21
	mInTickGeneric            = 45
	mInTickString             = 46
	mInTickEFP                = 47
	mInCurrentTime            = 49
	mInRealtimeBars           = 50
	mInFundamentalData        = 51
	mInContractDataEnd        = 52
	mInOpenOrderEnd           = 53
	mInAccountDownloadEnd     = 54
	mInExecutionDataEnd       = 55
	mInDeltaNeutralValidation = 56
	mInTickSnapshotEnd        = 57
	mInMarketDataType         = 58
	mInCommisionReport        = 59
	mInPosition               = 61
	mInPositionEnd            = 62
	mInAccountSummary         = 63
	mInAccountSummaryEnd      = 64
)

const (
	// outgoing message ids
	mOutRequestMarketData          = 1
	mOutCancelMarketData           = 2
	mOutPlaceOrder                 = 3
	mOutCancelOrder                = 4
	mOutRequestOpenOrders          = 5
	mOutRequestAccountData         = 6
	mOutRequestExecutions          = 7
	mOutRequestIds                 = 8
	mOutRequestContractData        = 9
	mOutRequestMarketDepth         = 10
	mOutCancelMarketDepth          = 11
	mOutRequestNewsBulletins       = 12
	mOutCancelNewsBulletins        = 13
	mOutSetServerLogLevel          = 14
	mOutRequestAutoOpenOrders      = 15
	mOutRequestAllOpenOrders       = 16
	mOutRequestManagedAccounts     = 17
	mOutRequestFA                  = 18
	mOutReplaceFA                  = 19
	mOutRequestHistoricalData      = 20
	mOutExerciseOptions            = 21
	mOutRequestScannerSubscription = 22
	mOutCancelScannerSubscription  = 23
	mOutRequestScannerParameters   = 24
	mOutCancelHistoricalData       = 25
	mOutRequestCurrentTime         = 49
	mOutRequestRealtimeBars        = 50
	mOutCancelRealtimeBars         = 51
	mOutRequestFundamentalData     = 52
	mOutCancelFundamentalData      = 53
	mOutRequestCalcImpliedVol      = 54
	mOutRequestCalcOptionPrice     = 55
	mOutCancelCalcImpliedVol       = 56
	mOutCancelCalcOptionPrice      = 57
	mOutRequestGlobalCancel        = 58
	mOutRequestMarketDataType      = 59
	mOutRequestPositions           = 61
	mOutRequestAccountSummary      = 62
	mOutCancelAccountSummary       = 63
	mOutCancelPositions            = 64
)

const (
	TickBidSize               = 0
	TickBid                   = 1
	TickAsk                   = 2
	TickAskSize               = 3
	TickLast                  = 4
	TickLastSize              = 5
	TickHigh                  = 6
	TickLow                   = 7
	TickVolume                = 8
	TickClose                 = 9
	TickBidOptionComputation  = 10
	TickAskOptionComputation  = 11
	TickLastOptionComputation = 12
	TickModelOption           = 13
	TickOpen                  = 14
	TickLow13Week             = 15
	TickHigh13Week            = 16
	TickLow26Week             = 17
	TickHigh26Week            = 18
	TickLow52Week             = 19
	TickHigh52Week            = 20
	TickAverageVolume         = 21
	TickOpenInterest          = 22
	TickOptionHistoricalVol   = 23
	TickOptionImpliedVol      = 24
	TickOptionBidExch         = 25
	TickOptionAskExch         = 26
	TickOptionCallOpenInt     = 27
	TickOptionPutOpenInt      = 28
	TickOptionCallVolume      = 29
	TickOptionPutVolume       = 30
	TickIndexFuturePremium    = 31
	TickBidExch               = 32
	TickAskExch               = 33
	TickAuctionVolume         = 34
	TickAuctionPrice          = 35
	TickAuctionImbalance      = 36
	TickMarkPrice             = 37
	TickBidEFPComputation     = 38
	TickAskEFPComputation     = 39
	TickLastEFPComputation    = 40
	TickOpenEFPComputation    = 41
	TickHighEFPComputation    = 42
	TickLowEFPComputation     = 43
	TickCloseEFPComputation   = 44
	TickLastTimestamp         = 45
	TickShortable             = 46
	TickFundamentalRations    = 47
	TickRTVolume              = 48
	TickHalted                = 49
	TickBidYield              = 50
	TickAskYield              = 51
	TickLastYield             = 52
	TickCustOptionComputation = 53
	TickTradeCount            = 54
	TickTradeRate             = 55
	TickVolumeRate            = 56
	TickLastRTHTrade          = 57
	TickNotSet                = 58
)
