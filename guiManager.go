package main

import (
	"encoding/json"
	"fmt"
	ui "github.com/gizak/termui"
	// "strconv"
)

func getCoinsList() []string {
	var btc Coin
	var ltc Coin
	var eth Coin
	err := json.Unmarshal(CallApi("bitcoin"), &btc)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(CallApi("litecoin"), &ltc)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(CallApi("ethereum"), &eth)
	if err != nil {
		panic(err)
	}

	coins := []string{
		fmt.Sprintf("Ƀ BTC\n"),
		fmt.Sprintf("\t\t Ƀ %v  \n", btc[0].PriceBtc),
		fmt.Sprintf("\t\t $ %v", btc[0].PriceUsd),
		fmt.Sprintf("\t\t %v %% change(1h)", btc[0].PercentChange1H),
		fmt.Sprintf("\t\t %v %% change(24h)", btc[0].PercentChange24H),
		"",
		fmt.Sprintf("Ξ ETH\n"),
		fmt.Sprintf("\t\t Ƀ %v  \n", eth[0].PriceBtc),
		fmt.Sprintf("\t\t $ %v", eth[0].PriceUsd),
		fmt.Sprintf("\t\t %v %% change(1h)", eth[0].PercentChange1H),
		fmt.Sprintf("\t\t %v %% change(24h)", eth[0].PercentChange24H),
		"",
		fmt.Sprintf("Ł LTC\n"),
		fmt.Sprintf("\t\t Ƀ %v  \n", ltc[0].PriceBtc),
		fmt.Sprintf("\t\t $ %v", ltc[0].PriceUsd),
		fmt.Sprintf("\t\t %v %% change(1h)", ltc[0].PercentChange1H),
		fmt.Sprintf("\t\t %v %% change(24h)", ltc[0].PercentChange24H),
		"", "", "",
		fmt.Sprintf("Last update: %v", btc[0].LastUpdated),
	}

	return coins
}
func getTickerList() *ui.List {
	ls := ui.NewList()
	ls.Items = getCoinsList()
	ls.ItemFgColor = ui.ColorYellow
	ls.BorderLabel = "Trackers"
	ls.Height = 25
	ls.Width = 30
	ls.Y = 0

	return ls
}
