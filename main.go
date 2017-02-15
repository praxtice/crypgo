package main

import (
	//"fmt"
	ui "github.com/gizak/termui"
	"strconv"
	"time"
)

// CoinGui The GUI manage struct for coins
type CoinGui struct {
	Coin *Coin
}

func main() {
	var ss int
	// InitHistData()
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	var Gbtc CoinGui
	var Gltc CoinGui
	var Geth CoinGui

	var btc Coin
	var ltc Coin
	var eth Coin
	err = json.Unmarshal(CallApi("bitcoin"), &btc)
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
	btcLabel.Width = 2
	btcLabel.TextFgColor = ui.ColorWhite

	ethLabel := ui.NewPar("Ξ ETH")
	ethLabel.Height = 3
	ethLabel.Width = 2
	ethLabel.TextFgColor = ui.ColorWhite

	ltcLabel := ui.NewPar("Ł LTC")
	ltcLabel.Height = 3
	ltcLabel.Width = 2
	ltcLabel.TextFgColor = ui.ColorWhite

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(2, 5, btcLabel),
		),
		ui.NewRow(
			ui.NewCol(0, 3), // Price BTC Label
			ui.NewCol(0, 3), // Price BTC Value
			ui.NewCol(0, 3), // Price USD Label
			ui.NewCol(0, 3), // Price USD Value
			ui.NewCol(0, 3), // Percentage Change 1 h
			ui.NewCol(0, 3), // Percentage Change 24h
		),

		ui.NewRow(
			ui.NewCol(2, 5, ethLabel),
		),
		ui.NewRow(
			ui.NewCol(0, 3), // Price BTC Label
			ui.NewCol(0, 3), // Price BTC Value
			ui.NewCol(0, 3), // Price USD Label
			ui.NewCol(0, 3), // Price USD Value
			ui.NewCol(0, 3), // Percentage Change 1 h
			ui.NewCol(0, 3), // Percentage Change 24h
		),

		ui.NewRow(
			ui.NewCol(2, 5, ltcLabel),
		),
		ui.NewRow(
			ui.NewCol(0, 3), // Price BTC Label
			ui.NewCol(0, 3), // Price BTC Value
			ui.NewCol(0, 3), // Price USD Label
			ui.NewCol(0, 3), // Price USD Value
			ui.NewCol(0, 3), // Percentage Change 1 h
			ui.NewCol(0, 3), // Percentage Change 24h
		),
	)

	ui.Body.Align()
	ui.Render(ui.Body)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})
	// 1486923240
	ui.Merge("6s", ui.NewTimerCh(time.Second*6))
	ui.Handle("/timer/6s", func(e ui.Event) {
		ss++
		par.Text = strconv.Itoa(ss)
		tickerList.Items = getCoinsList()
		ui.Body.Align()
		ui.Clear()
		ui.Render(ui.Body)
	})

	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		tickerList = getTickerList()
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Clear()
		ui.Render(ui.Body)
	})

	ui.Loop()

}
