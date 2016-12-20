package goyhfin

import "fmt"

func GetTickerData(ticker, rangeStr, intervalStr string, includePrePostTradingPeriods bool) (ChartQueryResponse, error) {
	data := YahooFinanceResponse{}
	err := getJson("https://query1.finance.yahoo.com/v7/finance/chart/"+ticker+"?range="+rangeStr+"&interval="+intervalStr+"&indicators=quote&includeTimestamps=true&includePrePost="+fmt.Sprint(includePrePostTradingPeriods)+"&corsDomain=finance.yahoo.com", &data)
	if err != nil {
		return ChartQueryResponse{}, err
	}

	return data.GetFormattedOutput()
}
