package main

import (
	"encoding/json"
	"fmt"
	"strings"

	ui "github.com/gizak/termui"
	// "strconv"
)

// CoinGui The GUI manage struct for coins
type CoinGui struct {
	Coin       Coin
	PriceBtc   *ui.Par
	PriceUsd   *ui.Par
	Change1h   *ui.Par
	Change24h  *ui.Par
	LastUpdate *ui.Par
}

func newCoinGui(c Coin) CoinGui {
	Usd := ui.NewPar(fmt.Sprintf("\t\t $ %v", c[0].PriceUsd))
	Usd.Height = 3
	Usd.Width = 4

	Btc := ui.NewPar(strings.TrimSpace(fmt.Sprintf("\t\t Ƀ %v", c[0].PriceBtc)))
	Btc.Height = 3
	Btc.Width = 4

	Change1Hour := ui.NewPar(fmt.Sprintf("change(1h)\n%v %%", c[0].PercentChange24H))
	Change1Hour.Height = 4
	Change1Hour.Width = 4

	Change24Hour := ui.NewPar(fmt.Sprintf("change(24h)\n%v %%", c[0].PercentChange24H))
	Change24Hour.Height = 4
	Change24Hour.Width = 4

	LastUp := ui.NewPar(fmt.Sprintf("Last Update: %v", c[0].LastUpdated))
	LastUp.Height = 3
	LastUp.Width = 4

	GC := CoinGui{Coin: c, PriceBtc: Btc, PriceUsd: Usd, Change1h: Change1Hour, Change24h: Change24Hour, LastUpdate: LastUp}
	return GC
}

func getCoinsList() []string {
	var btc Coin
	var ltc Coin
	var eth Coin
	err := json.Unmarshal(CallAPI("bitcoin"), &btc)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(CallAPI("litecoin"), &ltc)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(CallAPI("ethereum"), &eth)
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
