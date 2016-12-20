package goyhfin

import (
	"time"
)

type YahooFinanceResponse struct {
	Chart YahooFinanceChart `json:"chart"`
}

func (resp *YahooFinanceResponse) GetFormattedOutput() (ChartQueryResponse, error) {
	if len(resp.Chart.Result) < 1 {
		return ChartQueryResponse{}, InvalidYahooFinanceResponseLengthError
	}
	out := ChartQueryResponse{
		Currency:       resp.Chart.Result[0].Meta.Currency,
		Symbol:         resp.Chart.Result[0].Meta.Symbol,
		ExchangeName:   resp.Chart.Result[0].Meta.ExchangeName,
		InstrumentType: resp.Chart.Result[0].Meta.InstrumentType,
		FirstTradeDate: time.Unix(resp.Chart.Result[0].Meta.FirstTradeDate, 0),
		GMTOffset:      time.Duration(resp.Chart.Result[0].Meta.GMTOffset) * time.Second,
		Timezone:       resp.Chart.Result[0].Meta.Timezone,
	}

	out.Quotes = make([]Quote, len(resp.Chart.Result[0].Timestamp))
	if len(resp.Chart.Result[0].Timestamp) < 2 {
		return ChartQueryResponse{}, InvalidYahooFinanceResponseNotEnoughDataError
	}
	periodSeconds := resp.Chart.Result[0].Timestamp[1] - resp.Chart.Result[0].Timestamp[0]
	for ind := range resp.Chart.Result[0].Timestamp {
		out.Quotes[ind] = Quote{
			OpensAt:  time.Unix(resp.Chart.Result[0].Timestamp[ind], 0),
			ClosesAt: time.Unix(resp.Chart.Result[0].Timestamp[ind]+periodSeconds, 0),
			Period:   time.Second * time.Duration(periodSeconds),
			Open:     resp.Chart.Result[0].Indicators.Quote[0].Open[ind],
			High:     resp.Chart.Result[0].Indicators.Quote[0].High[ind],
			Low:      resp.Chart.Result[0].Indicators.Quote[0].Low[ind],
			Close:    resp.Chart.Result[0].Indicators.Quote[0].Close[ind],
			Volume:   resp.Chart.Result[0].Indicators.Quote[0].Volume[ind],
		}
	}

	out.CurrentTradingPeriod.Pre = TradingPeriod{
		Timezone:  resp.Chart.Result[0].Meta.CurrentTradingPeriod.Pre.Timezone,
		GMTOffset: time.Second * time.Duration(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Pre.GMTOffset),
		Start:     time.Unix(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Pre.Start, 0),
		End:       time.Unix(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Pre.End, 0),
	}
	out.CurrentTradingPeriod.Regular = TradingPeriod{
		Timezone:  resp.Chart.Result[0].Meta.CurrentTradingPeriod.Regular.Timezone,
		GMTOffset: time.Second * time.Duration(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Regular.GMTOffset),
		Start:     time.Unix(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Regular.Start, 0),
		End:       time.Unix(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Regular.End, 0),
	}
	out.CurrentTradingPeriod.Post = TradingPeriod{
		Timezone:  resp.Chart.Result[0].Meta.CurrentTradingPeriod.Post.Timezone,
		GMTOffset: time.Second * time.Duration(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Post.GMTOffset),
		Start:     time.Unix(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Post.Start, 0),
		End:       time.Unix(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Post.End, 0),
	}

	out.TradingPeriods = make([][]TradingPeriod, len(resp.Chart.Result[0].Meta.TradingPeriods))
	for dayInd := range resp.Chart.Result[0].Meta.TradingPeriods {
		out.TradingPeriods[dayInd] = make([]TradingPeriod, len(resp.Chart.Result[0].Meta.TradingPeriods[dayInd]))
		for periodInd := range resp.Chart.Result[0].Meta.TradingPeriods[dayInd] {
			out.TradingPeriods[dayInd][periodInd] = TradingPeriod{
				Timezone:  resp.Chart.Result[0].Meta.TradingPeriods[dayInd][periodInd].Timezone,
				GMTOffset: time.Second * time.Duration(resp.Chart.Result[0].Meta.TradingPeriods[dayInd][periodInd].GMTOffset),
				Start:     time.Unix(resp.Chart.Result[0].Meta.TradingPeriods[dayInd][periodInd].Start, 0),
				End:       time.Unix(resp.Chart.Result[0].Meta.TradingPeriods[dayInd][periodInd].End, 0),
			}
		}
	}

	return out, nil
}

type YahooFinanceChart struct {
	Result []YahooFinanceResult `json:"result"`
	Error  interface{}          `json:"error"`
}

type YahooFinanceResult struct {
	Meta       YahooFinanceMeta       `json:"meta"`
	Timestamp  []int64                `json:"timestamp"`
	Indicators YahooFinanceIndicators `json:"indicators"`
}

type YahooFinanceMeta struct {
	Currency             string                           `json:"currency"`
	Symbol               string                           `json:"symbol"`
	InstrumentType       string                           `json:"instrumentType"`
	ExchangeName         string                           `json:"exchangeName"`
	FirstTradeDate       int64                            `json:"firstTradeDate"`
	GMTOffset            int64                            `json:"gmtoffset"`
	Timezone             string                           `json:"timezone"`
	CurrentTradingPeriod YahooFinanceCurrentTradingPeriod `json:"currentTradingPeriod"`
	TradingPeriods       [][]YahooFinanceTradingPeriod    `json:"tradingPeriods"`
}

type YahooFinanceCurrentTradingPeriod struct {
	Pre     YahooFinanceTradingPeriod `json:"pre"`
	Regular YahooFinanceTradingPeriod `json:"regular"`
	Post    YahooFinanceTradingPeriod `json:"post"`
}

type YahooFinanceTradingPeriod struct {
	Timezone  string `json:"timezone"`
	Start     int64  `json:"start"`
	End       int64  `json:"end"`
	GMTOffset int64  `json:"gmtoffset"`
}

type YahooFinanceIndicators struct {
	Quote []YahooFinanceQuote `json:"quote"`
}

type YahooFinanceQuote struct {
	High   []float64 `json:"high"`
	Open   []float64 `json:"open"`
	Low    []float64 `json:"low"`
	Close  []float64 `json:"close"`
	Volume []float64 `json:"volume"`
}

type ChartQueryResponse struct {
	Currency             string
	Symbol               string
	ExchangeName         string
	InstrumentType       string
	FirstTradeDate       time.Time
	GMTOffset            time.Duration
	Timezone             string
	CurrentTradingPeriod CurrentTradingPeriod
	TradingPeriods       [][]TradingPeriod
	Quotes               []Quote
}

type Quote struct {
	OpensAt  time.Time
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   float64
	ClosesAt time.Time
	Period   time.Duration
}

type CurrentTradingPeriod struct {
	Pre     TradingPeriod
	Regular TradingPeriod
	Post    TradingPeriod
}

type TradingPeriod struct {
	Timezone  string
	Start     time.Time
	End       time.Time
	GMTOffset time.Duration
}
