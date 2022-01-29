// Copyright 2022 David M. Hornsby.
// Use of this source code file is governed by a license that
// can be found in the LICENSE file.

package main

import (
	"govolution/Carrier"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	log.Print("Test")
	p := &Carrier.Protocol{Device: "/dev/ttyUSB0"}
	go p.Run()
	wg.Add(1)

	wg.Wait()

}
