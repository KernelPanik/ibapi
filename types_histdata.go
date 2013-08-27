package ibapi

import (
	"time"
)

type MsgInHistDataItem struct {
	Date string
	Open float64
	High float64
	Low float64
	Close float64
	Volume int64
	WAP float64
	HasGaps string
	BarCount int64 `minVer:"3"`
}

type MsgInHistData struct {
	ReqId int64
	StartDate string
	EndDate string
	Items []MsgInHistDataItem
}

type MsgOutReqHistData struct {
	ReqId int64
	Symbol string
	SecurityType string
	Expiry string
	Strike float64
	Right string
	Multiplier string
	Exchange string
	PrimaryExchange string
	Currency string
	LocalSymbol string
	IncludeExpired string `minVer:"31"`
	EndDateTime time.Time `minVer:"20"`
	BarSizeSetting string `minVer:"20"`
	Duration string // format is "x S|D|W|M|Y" where x is integer, sec/day/week/mon/year
	UseRTH bool
	WhatToShow string
	FormatDate int64 `minVer:"16"`
}

const (
	FormatDateString = 1
	FormatDateSeconds = 2
	WhatToShowTrades = "TRADES"
	WhatToShowMidpoint = "MIDPOINT"
	WhatToShowBid = "BID"
	WhatToShowAsk = "ASK"
	WhatToShowBidAsk = "BID_ASK"
	WhatToShowHistVol = "HISTORICAL_VOLATILITY"
	WhatToShowImpVol = "OPTION_IMPLIED_VOLATILITY"
	BarSize1Sec = "1 sec"
	BarSize5Sec = "5 secs"
	BarSize15Sec = "15 secs"
	BarSize30Sec = "30 secs"
	BarSize1Min = "1 min"
	BarSize2Min = "2 mins"
	BarSize3Min = "3 mins"
	BarSize5Min = "5 mins"
	BarSize15Min = "15 mins"
	BarSize30Min = "30 mins"
	BarSize1Hr = "1 hour"
	BarSize1Day = "1 day"
)
