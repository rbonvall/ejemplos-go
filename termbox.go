package main

import (
	//"fmt"
	tb "github.com/nsf/termbox-go"
)


func main() {
	err := tb.Init()
	if err != nil {
		panic(err)
	}
	defer tb.Close()

	tb.SetInputMode(tb.InputEsc)
	tb.Clear(tb.ColorCyan, tb.ColorBlue)
	tb.Flush()

loop:	for i := 0; ; i++ {
		ev := tb.PollEvent()
		switch ev.Type {
		case tb.EventKey:
			if ev.Key == tb.KeyCtrlC {
				break loop
			}
			tb.SetCell(i, i, ev.Ch, tb.ColorGreen, tb.ColorBlack)
			tb.Flush()
		case tb.EventError:
			panic(ev.Err)
		}
	}
}

