package main

import "flag"

type Options struct {
	help bool
	version bool
	print_draws bool
	freq bool

	numbers string
	draws string
}

func (o *Options) Init() *flag.FlagSet {
	var fs = flag.NewFlagSet("Options", flag.ContinueOnError)
	fs.BoolVar(&(options.help), "help", false, "Shows this message")
	fs.BoolVar(&(options.version), "version", false, "Show program version")
	fs.BoolVar(&(options.print_draws), "print-draws", false, "Print loaded draws")
	fs.BoolVar(&(options.freq), "freq", false, "Frequency analysis")

	fs.StringVar(&(options.numbers), "numbers", "", "Numbers to test");
	fs.StringVar(&(options.draws), "draws", "", "Draws file name");
	return fs
}
