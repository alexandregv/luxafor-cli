package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/zshift/luxafor"
)

func main() {
	luxs := luxafor.Enumerate()
	if len(luxs) == 0 {
		fmt.Println("No attached devices. Exiting.")
		os.Exit(0)
	}

	lux := luxs[1]

	if len(os.Args) <= 3 {
		fmt.Println("Usage: luxafor-cli <red> <green> <blue>\nExample: luxafor-cli 0 255 0")
		os.Exit(1)
	}

	r, err := strconv.ParseUint(os.Args[1], 10, 8)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	g, err := strconv.ParseUint(os.Args[2], 10, 8)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	b, err := strconv.ParseUint(os.Args[3], 10, 8)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = lux.Solid(uint8(r), uint8(g), uint8(b))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
