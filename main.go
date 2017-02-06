package main

import (
	ui "github.com/gizak/termui"
)

var ss int = 0

func main() {
	// InitHistData()
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	tickerList := getTickerList()

	par0 := ui.NewPar(string(ss))
	par0.Height = 10
	par0.Width = 20
	par0.Y = 1
	par0.Border = false

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, tickerList),
		),
		ui.NewRow(
			ui.NewCol(6, 0, par0),
		),
	)

	ui.Body.Align()
	ui.Render(ui.Body)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.StopLoop()
	})

	ui.Handle("/timer/1s", func(e ui.Event) {
		ss++
		ui.Body.Align()
		tickerList = getTickerList()
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
