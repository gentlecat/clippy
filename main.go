package main

import (
	"fmt"
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
		"        ",
	}
	clippyWidth = len(clippy[0])
)

// TODO: Allow to provide a title
func generateTextBox(text string, width int) (out string) {
	const (
		side   = "│"
		top    = "╭─╮"
		bottom = "╰─╯"
	)
	return
}

func main() {
	fmt.Println("WIP")
}
