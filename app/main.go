package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	asciiLogo1 := figure.NewColorFigure(
		"BraiNiac", "elite", "green", true)
	fmt.Print("\033[H\033[2J")
	fmt.Print("\n")
	asciiLogo1.Print()
	fmt.Print("\n")

}
