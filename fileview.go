package main

import (
	"github.com/nsf/termbox-go"
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

const W = 40
const H = 20

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func draw_box() {
	fg := termbox.ColorGreen
	bg := termbox.ColorDefault
	termbox.SetCell(    0,     0, '┌' , fg , bg)
	termbox.SetCell(    0, H + 1, '└' , fg , bg)
	termbox.SetCell(W + 1,     0, '┐' , fg , bg)
	termbox.SetCell(W + 1, H + 1, '┘' , fg , bg)
	for j := 1; j <= W; j++ {
		termbox.SetCell(j,     0, '─', fg, bg)
		termbox.SetCell(j, H + 1, '─', fg, bg)
	}
	for i := 1; i <= H; i++ {
		termbox.SetCell(    0, i,  '│', fg, bg)
		termbox.SetCell(W + 1, i,  '│', fg, bg)
	}
	termbox.Flush()
}

func show_text(text []string, first_line, first_column int) {
	fg := termbox.ColorBlue
	fg2 := termbox.ColorCyan
	bg := termbox.ColorDefault
	for row := 1; row <= H; row++ {
		nr_line := first_line + row - 1
		var line string
		if 0 <= nr_line && nr_line < len(text) {
			line = text[nr_line]
		} else {
			break
		}
		line = strings.Replace(line, "\t", "        ", -1)

		for col := 1; col <= W; col++ {
			nr_char := first_column + col - 1
			var ch rune = ' '
			if 0 <= nr_char && nr_char < len(line) {
				ch = rune(line[nr_char])
			}
			termbox.SetCell(col, row, ch, fg, bg)
		}
		if (len(line) - first_column > W) {
			termbox.SetCell(W, row, '…', fg2, bg)
		}
	}
	termbox.Flush()
}

func print_at(j, i int, s string, fg termbox.Attribute) {
	for c, r := range s {
		termbox.SetCell(j + c, i, r, fg, termbox.ColorDefault)
	}
}


func show_help() {
	fg := termbox.ColorRed
	col := W + 5
	print_at(col, 2, "h ←", fg)
	print_at(col, 3, "j ↓", fg)
	print_at(col, 4, "k ↑", fg)
	print_at(col, 5, "l →", fg)
	print_at(col, 6, "g ↟", fg)
	print_at(col, 7, "G ↡", fg)
	print_at(col, 8, "0 ↞", fg)
	print_at(col, 9, "q ☠", fg)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No file was provided.")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Could not open file.")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error while reading contents.")
		return
	}
	file.Close()

	contents := strings.Split(string(data), "\n")

	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	draw_box()
	show_help()

	row := 0
	col := 0

loop:	for {
		show_text(contents, row, col)
		event := termbox.PollEvent()
		switch event.Type {
		case termbox.EventError:
			panic(event.Err)
		case termbox.EventKey:
			switch event.Ch {
			case 'j':
				row = min(row + 1, len(contents) - H)
			case 'k':
				row = max(row - 1, 0)
			case 'h':
				col = max(col - 1, 0)
			case 'l':
				col = col + 1
			case 'g':
				row = 0
			case 'G':
				row = len(contents) - H
			case '0':
				col = 0
			case 'q':
				break loop
			}

		}
	}

	return

}

