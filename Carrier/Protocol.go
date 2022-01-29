// Copyright 2022 David M. Hornsby.
// Use of this source code file is governed by a license that
// can be found in the LICENSE file.

// The Carrier Package provides a framework for communicating with
// Carrier Infinity and Bryant Evolution Furnaces.
package Carrier

import (
	"log"
	"time"

	"github.com/tarm/serial"
)

type Protocol struct {
	Device string
	port   *serial.Port
}

func (p *Protocol) openSerial() {
	var e error
	log.Print("Attempting to open Serial")
	p.port, e = serial.OpenPort(&serial.Config{Name: p.Device, Baud: 38400, ReadTimeout: 5 * time.Second})
	if e != nil {
		log.Fatal("Error OPening Serial")
	}
}

func (p *Protocol) Run() {
	log.Print("Starting Protocol")
	msg := []byte{}
	buf := make([]byte, 1024)

	for { //daemonize

		if p.port == nil {
			log.Print("Need to Open Serial Port")
			msg = []byte{}
			log.Print(p.Device)
			p.openSerial()
		}

		n, err := p.port.Read(buf)
		if n == 0 || err != nil {
			log.Print("Runtime Serial error resetting...")
			if p.port != nil {
				p.port.Close()
			}
			p.port = nil
			continue
		}

		msg = append(msg, buf[:n]...)

		for { //parse msg
			if len(msg) < 10 { // If we don't have a full header
				break
			}
			l := int(msg[4]) + 10
			if len(msg) < l { // If we don't have a full frame
				break
			}

			f, err := NewFrame(msg[:l])
			if err != nil {
				msg = msg[:copy(msg, msg[1:])] // discard first byte and try again
			}
			if err == nil {
				dispatchFrame(f)
				msg = msg[:copy(msg, msg[l:])]
			}
		}
	}

}
