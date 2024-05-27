package main

import (
	"flag"
	"os"
)

type Encode struct {
	Input  string
	Output string
}

func (e *Encode) Parse(args []string) {
	fs := flag.NewFlagSet("encode", flag.ExitOnError)
	fs.StringVar(&e.Input, "input", "", "input file")
	fs.StringVar(&e.Output, "output", "", "output file")
	fs.Parse(args)
}

func (e *Encode) Run() {
	data, err := os.ReadFile(e.Input)
	Must(err)
	encoded, err := Encrypt(data)
	Must(err)
	Must(os.WriteFile(e.Output, encoded, 0644))
}
