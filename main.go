package main

import (
	"os"
	"fmt"
)

const version string = "0.0.1"

var options Options

func main () {
	processFlags()
	toto := NewToto()
	if options.numbers != "" {
		toto.ProcessNumbers(options.numbers)
	}
	if options.draws != "" {
		toto.ProcessDraws(options.draws)
	}
	if options.print_draws {
		fmt.Print("draws")
		toto.Print()
	}
}

func processFlags () {
	var fs = options.Init()
	fs.Parse(os.Args[1:])

	if options.version {
		fmt.Println("Toto version:", version)
		os.Exit(0)
	}

	if options.help {
		fmt.Println("Toto version:", version)
		fmt.Println("Usage:")
		fs.PrintDefaults()
		os.Exit(0)
	}
}
