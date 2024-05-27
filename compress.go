package main

import (
	"compress/gzip"
	"flag"
	"os"
)

type Compress struct {
	Input  string
	Output string
}

func (c *Compress) Parse(args []string) {
	fs := flag.NewFlagSet("compress", flag.ExitOnError)
	fs.StringVar(&c.Input, "input", "", "input file")
	fs.StringVar(&c.Output, "output", "", "output file")
	fs.Parse(args)
}

func (c *Compress) Run() {
	data, err := os.ReadFile(c.Input)
	Must(err)
	decoded, err := Decrypt(data)
	Must(err)
	f, err := os.Create(c.Output)
	Must(err)
	defer f.Close()
	w := gzip.NewWriter(f)
	_, err = w.Write(decoded)
	Must(err)
	Must(w.Close())
}
