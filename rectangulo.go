package main

import (
	"github.com/nsf/termbox-go"
)

func main() {
	termbox.Init()
	defer termbox.Close()

	w, _ := termbox.Size()

	fg := termbox.ColorWhite
	bg := termbox.ColorBlue

	for y := 5; y < 10; y++ {
		for x := 10; x < 15; x++ {
			termbox.SetCell(x, y, 'X', fg, bg)
		}
		for x := 20; x < 25; x++ {
			termbox.SetCell(x, y, ' ', fg, bg)
		}
		for x := 30; x < 35; x++ {
			termbox.CellBuffer()[x + y * w].Bg = termbox.ColorRed
		}
	}
	termbox.Flush()

	termbox.PollEvent()
}
