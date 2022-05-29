package domain

import (
	"simplifica/src/support"
	"strings"
)

type Stock struct {
	Ticker string
	Amount float64
}

func BuildStocks(info map[string]float64) []Stock {
	var stocks []Stock

	for ticker, amount := range info {
		stocks = append(stocks, Stock{
			Ticker: ticker,
			Amount: amount,
		})
	}

	return stocks
}

func CollectStocks(stockNames []string, stockDividends []string) map[string]float64 {
	info := make(map[string]float64)

	for i, name := range stockNames {
		ticker := strings.Fields(name)[0]

		amount, alreadyExists := info[ticker]

		if alreadyExists {
			info[ticker] = support.Round(amount + support.ConvertAmountToFloat(stockDividends[i*2+1]))
		} else {
			info[ticker] = support.ConvertAmountToFloat(stockDividends[i*2+1])
		}
	}

	return info
}
