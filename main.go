package main

import (
	//"fmt"
	ui "github.com/gizak/termui"
	"strconv"
	"time"
)

func main() {
	var ss int = 0
	// InitHistData()
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	tickerList := getTickerList()

	par := ui.NewPar("")
	par.Height = 3
	par.Width = 17
	par.BorderLabel = "Label"
	par.TextFgColor = ui.ColorGreen
	par.TextBgColor = ui.ColorMagenta

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, tickerList),
		),
		ui.NewRow(
			ui.NewCol(12, 0, par),
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
		ss += 1
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
