package main

import (
	"os"
)

func main() { // HOW TO: left click to select and commit a move, right click to deselect. close the window or press esc to exit the program
	if err := run(); err != 0 {
		os.Exit(1)
	}
}
