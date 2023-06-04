package main

import (
	"flag"
	"github.com/senior-cyber/liquibase/xml"
	"log"
)

func init() {
	log.SetPrefix("[liquibase] : ")
}

func main() {

	var _input string
	var _output string

	flag.StringVar(&_input, "input", "", "")
	flag.StringVar(&_output, "output", "", "")

	flag.Parse()

	_xml := xml.New(_input, _output)
	_xml.ParseXml()
}
