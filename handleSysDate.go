package main

import (
	"encoding/json"
	"govolution/Carrier"
	"log"
)

func init() {
	Carrier.RegisterFrameHandler(TSTATDATE{})
}

type TSTATDATE struct {
}

func (t TSTATDATE) Match(f *Carrier.Frame) bool {
	return true
}

func (t TSTATDATE) Handle(f *Carrier.Frame) {
	s, _ := json.Marshal(f)
	log.Print(string(s))
}
