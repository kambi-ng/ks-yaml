package main

import (
	"fmt"
	"io/ioutil"
	"log"

	ksyaml "github.com/bukan-kambing/ks-yaml/pkg"
)

func main() {
	b, err := ioutil.ReadFile("test.yml")
	if err != nil {
		log.Fatal(err)
	}
	c := ksyaml.NewConverter()
	ms := string(b)
	os, _ := c.Convert(ms)
	fmt.Println(os)
}
