package main

import (
	//"fmt"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	ui "github.com/gizak/termui"
)

func label(s string) *ui.Par {
	l := ui.NewPar(s)
	l.Height = 5
	l.Width = 3
	l.Border = false
	return l
}

func main() {
	ss := 0
	// InitHistData()?
	// Read old hist data
	// create new struct for old data

	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()
	var btc Coin
	var ltc Coin
	var eth Coin
	err = json.Unmarshal(CallAPI("bitcoin"), &btc)
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

	Gbtc := newCoinGui(btc)
	Gltc := newCoinGui(ltc)
	Geth := newCoinGui(eth)
	// ui.Body.AddRows(
	// 	ui.NewRow(
	// 		ui.NewCol(12, 0, tickerList),
	// 	),
	// 	ui.NewRow(
	// 		ui.NewCol(12, 0, par),
	// 	),
	// )

	btcLabel := ui.NewPar("Ƀ BTC")
	btcLabel.Height = 3
	btcLabel.Width = 3
	btcLabel.TextFgColor = ui.ColorWhite

	ethLabel := ui.NewPar("Ξ ETH")
	ethLabel.Height = 3
	ethLabel.Width = 3
	ethLabel.TextFgColor = ui.ColorWhite

	ltcLabel := ui.NewPar("Ł LTC")
	ltcLabel.Height = 3
	ltcLabel.Width = 3
	ltcLabel.TextFgColor = ui.ColorWhite

	ssLabel := ui.NewPar("refresh counter: 0")
	ssLabel.Height = 3
	ssLabel.Width = 3
	ssLabel.TextFgColor = ui.ColorWhite
	ssLabel.Border = false

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(2, 5, btcLabel),
		),
		// ui.NewRow(
		// 	ui.NewCol(4, 0, label("Price BTC")), // Price BTC Label
		// 	ui.NewCol(4, 0, label("Price USD")), // Price USD Label
		// 	ui.NewCol(4, 0, label("Change 1h")), // Percentage Change 1 h
		// 	ui.NewCol(4, 0, label("Change 24h")),
		// ),
		ui.NewRow(
			ui.NewCol(4, 2, Gbtc.PriceBtc), // Price BTC Value
			ui.NewCol(4, 0, Gbtc.PriceUsd), // Price USD Value
		),
		ui.NewRow(
			ui.NewCol(4, 2, Gbtc.Change1h),  // Percentage Change 1h
			ui.NewCol(4, 0, Gbtc.Change24h), // Percentage Change 24h

		),
		ui.NewRow(
			ui.NewCol(2, 5, ethLabel),
		),
		// ui.NewRow(
		// 	ui.NewCol(4, 0, label("Price BTC")), // Price BTC Label
		// 	ui.NewCol(4, 0, label("Price USD")), // Price USD Label
		// 	ui.NewCol(4, 0, label("Change 1h")), // Percentage Change 1 h
		// 	ui.NewCol(4, 0, label("Change 24h")),
		// ),
		ui.NewRow(
			ui.NewCol(4, 2, Geth.PriceBtc), // Price BTC Value
			ui.NewCol(4, 0, Geth.PriceUsd), // Price USD Value
		),
		ui.NewRow(
			ui.NewCol(4, 2, Geth.Change1h),  // Percentage Change 1h
			ui.NewCol(4, 0, Geth.Change24h), // Percentage Change 24h

		),

		ui.NewRow(
			ui.NewCol(2, 5, ltcLabel),
		),
		// ui.NewRow(
		// 	ui.NewCol(4, 0, label("Price BTC")), // Price BTC Label
		// 	ui.NewCol(4, 0, label("Price USD")), // Price USD Label
		// 	ui.NewCol(4, 0, label("Change 1h")), // Percentage Change 1 h
		// 	ui.NewCol(4, 0, label("Change 24h")),
		// ),
		ui.NewRow(
			ui.NewCol(4, 2, Gltc.PriceBtc), // Price BTC Value
			ui.NewCol(4, 0, Gltc.PriceUsd), // Price USD Value
		),
		ui.NewRow(
			ui.NewCol(4, 2, Gltc.Change1h),  // Percentage Change 1h
			ui.NewCol(4, 0, Gltc.Change24h), // Percentage Change 24h

		),
		ui.NewRow(
			ui.NewCol(8, 2, Gbtc.LastUpdate),
		),
		ui.NewRow(
			ui.NewCol(4, 2, ssLabel),
		),
	)

	ui.Body.Align()
	ui.Render(ui.Body)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Merge("6s", ui.NewTimerCh(time.Second*6))
	ui.Handle("/timer/6s", func(e ui.Event) {
		ss++
		ssLabel.Text = fmt.Sprintf("refresh counter: %v", strconv.Itoa(ss))
		err = json.Unmarshal(CallAPI("bitcoin"), &btc)
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

		Gbtc = newCoinGui(btc)
		Gltc = newCoinGui(ltc)
		Geth = newCoinGui(eth)
		//tickerList.Items = getCoinsList()
		ui.Body.Align()
		ui.Clear()
		ui.Render(ui.Body)
	})

	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		//tickerList = getTickerList()
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Clear()
		ui.Render(ui.Body)
	})

	ui.Loop()

}
