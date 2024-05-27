package main

import (
	"compress/gzip"
	"flag"
	"io"
	"os"
)

type Decompress struct {
	Input  string
	Output string
}

func (d *Decompress) Parse(args []string) {
	fs := flag.NewFlagSet("decompress", flag.ExitOnError)
	fs.StringVar(&d.Input, "input", "", "input file")
	fs.StringVar(&d.Output, "output", "", "output file")
	fs.Parse(args)
}

func (d *Decompress) Run() {
	f, err := os.Open(d.Input)
	Must(err)
	g, err := gzip.NewReader(f)
	Must(err)
	data, err := io.ReadAll(g)
	Must(err)
	encoded, err := Encrypt(data)
	Must(err)
	Must(os.WriteFile(d.Output, encoded, 0644))
}
