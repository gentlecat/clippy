package main

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	textWidth = 40
)

var (
	clippy = []string{
		// All the strings that are part of the clippy must be the same length.
		// Otherwise shit will break.
		" ╭──╮   ",
		" │  │   ",
		" @  @  ╭",
		" ││ ││ │",
		" ││ ││ ╯",
		" │╰─╯│  ",
		" ╰───╯  ",
	}
	clippyHeight = len(clippy)
	clippyWidth  = len(clippy[0])

	bubbleSide = "│"
	// Adding 2 to account for side characters
	bubbleTop    = "╭" + strings.Repeat("─", textWidth+2) + "╮"
	bubbleBottom = "╰" + strings.Repeat("─", textWidth+2) + "╯"
)

func splitString(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}
	// Adding missing spaces for the last one
	last := subs[len(subs)-1]
	subs[len(subs)-1] = last + strings.Repeat(" ", n-len(last))
	return subs
}

// TODO: Allow to provide a title
func generateTextBox(text string, width int) (rows []string) {
	rows = append(rows, bubbleTop)
	for _, substr := range splitString(text, width) {
		rows = append(rows, bubbleSide+" "+substr+" "+bubbleSide)
	}
	rows = append(rows, bubbleBottom)
	return rows
}

func combineWithClippy(rows []string) (output []string) {
	// Prepending an empty row to align with Clippy
	rows = append([]string{""}, rows...)

	// Outputting the textbox
	for i, row := range rows {
		if i > clippyHeight {
			row = strings.Repeat(" ", clippyWidth) + row
		} else {
			row = clippy[i] + row
		}
		output = append(output, row)
	}

	// Outputting the rest of Clippy if the textbox is not high enough
	for i := len(rows); i < clippyHeight; i++ {
		output = append(output, clippy[i])
	}

	return output
}

func main() {
	text := ""
	for _, row := range combineWithClippy(generateTextBox(text, textWidth)) {
		fmt.Println(row)
	}
}
