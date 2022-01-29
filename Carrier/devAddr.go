// Copyright 2022 David M. Hornsby.
// Use of this source code file is governed by a license that
// can be found in the LICENSE file.

// The Carrier Package provides a framework for communicating with
// Carrier Infinity and Bryant Evolution Furnaces.
package Carrier

type devAddr int

const (
	TSTATADDR = iota
	INDOORADDR
	SAMADDR
	BROADCASTADDR
)

func devAddrFromUint16(a uint16) devAddr {
	switch a {
	case 0x2001:
		return TSTATADDR
	case 0x4001:
		return INDOORADDR
	case 0x9201:
		return SAMADDR
	}
	return BROADCASTADDR
}

func (d devAddr) Encode() uint16 {
	switch d {
	case TSTATADDR:
		return 0x2001
	case INDOORADDR:
		return 0x4001
	case SAMADDR:
		return 0x9201
	case BROADCASTADDR:
		return 0xf1f1

	}
	return 0x00
}
