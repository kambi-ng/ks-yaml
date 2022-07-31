package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	ksyaml "github.com/bukan-kambing/ks-yaml/pkg"
)

type opts struct {
	inFile   string
	toStdout bool
	outFile  string
	indent   int
}

var opt opts

func main() {
	b, err := ioutil.ReadFile(opt.inFile)
	if err != nil {
		log.Fatal(err)
	}
	c := ksyaml.NewConverter(ksyaml.WithIndentation(opt.indent))
	ms := string(b)
	os, _ := c.Convert(ms)

	if opt.toStdout {
		fmt.Println(os)
		return
	}

	if err := ioutil.WriteFile(opt.outFile, []byte(os), 0644); err != nil {
		log.Fatal(err)
	}
}

func init() {
	indent := flag.Int("i", 4, "indent")
	outFile := flag.String("o", "", "output file")
	inFile := flag.String("f", "", "input file")
	flag.Parse()
	opt.indent = *indent
	opt.inFile = *inFile
	if *outFile != "" {
		opt.outFile = *outFile
	} else {
		opt.toStdout = true
	}
}
