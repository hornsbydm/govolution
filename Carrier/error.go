// Copyright 2022 David M. Hornsby.
// Use of this source code file is governed by a license that
// can be found in the LICENSE file.

// The Carrier Package provides a framework for communicating with
// Carrier Infinity and Bryant Evolution Furnaces.
package Carrier

type CarrierError int

const (
	ERRBADCRC CarrierError = iota
	ERRBIGFRAME
)

func (e CarrierError) Error() string {
	switch e {
	case ERRBADCRC:
		return "Invalid CRC"
	case ERRBIGFRAME:
		return "Frame too large"
	}
	return "Unknown Error"
}
