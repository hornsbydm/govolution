// Copyright 2022 David M. Hornsby.
// Use of this source code file is governed by a license that
// can be found in the LICENSE file.

// The Carrier Package provides a framework for communicating with
// Carrier Infinity and Bryant Evolution Furnaces.
package Carrier

import (
	"bytes"
	"encoding/binary"

	"github.com/npat-efault/crc16"
)

// Frame provides a structure to capture the Carrier Framing.
// See Docs/framing.md for details.
type Frame struct {
	Dst devAddr  `json:"dest"` // Bytes 0:1 Destination Address
	Src devAddr  `json:"srce"` // Bytes 2:3 Source Address
	Len uint8    // Byte 4 Data Length
	Opr frameOpr // Byte 7 Operation
	Tab uint8    `json:"table"` // Byte 9 Table
	Row uint8    `json:"row"`   // Byte 10 Row
	Dat []byte   `json:"data"`  //Bytes 11:-1 Row contents
}

// encode Frame into a slice of bytes for the protocol.
func (f *Frame) Encode() (a []byte, e error) {
	// b := make([]byte, 10 + len(frame.data))
	if len(f.Dat) > 255 {
		return a, ERRBIGFRAME
	}

	var b bytes.Buffer

	binary.Write(&b, binary.BigEndian, f.Dst.Encode()) //Bytes 0 and 1 Destination Address
	binary.Write(&b, binary.BigEndian, f.Src.Encode()) //Bytes 2 and 3 Source Address
	b.WriteByte(byte(len(f.Dat) + 3))                  // byte 4 Data Length
	b.WriteByte(0)                                     //byte 5 unknown
	b.WriteByte(0)                                     //byte 6 unknown
	b.WriteByte(f.Opr.Encode())                        //byte 7 operation
	b.WriteByte(0)                                     //byte 8 is usually 0x00
	b.WriteByte(f.Tab)                                 //byte 9 is the table
	b.WriteByte(f.Row)                                 //byte 10 is the row
	b.Write(f.Dat)                                     // Row Data
	b.Write(crc(b.Bytes()))                            //Last 2 bytes crc checksum

	return b.Bytes(), nil
}

// Create a new frame from raw slice of bytes.
func NewFrame(buf []byte) (*Frame, error) {

	// Validate CRC
	if !bytes.Equal(crc(buf[:len(buf)-2]), buf[len(buf)-2:]) {
		return &Frame{}, ERRBADCRC
	}

	f := &Frame{}

	// Bytes 0:1 - Destination Device Address
	f.Dst = devAddrFromUint16(binary.BigEndian.Uint16(buf[0:2]))
	// Bytes 2:3 - Source Device Address
	f.Src = devAddrFromUint16(binary.BigEndian.Uint16(buf[2:4]))
	// Byte 4 - Data Length
	f.Len = buf[4] - 0x03
	// Bytes 5:6 - Unknown
	// Byte 7 - Operation Code
	f.Opr = newFrameOpr(buf[7])
	// TODO: Add logic for pulling out register information
	// Byte 8 - Unknown
	// Byte 9 - Table Address
	f.Tab = buf[9]
	// Byte 10 - Row Address
	f.Row = buf[10]
	// Bytes 11:-2 - Row Data
	if len(buf) > 13 {
		f.Dat = buf[11 : len(buf)-2]
	}
	// Bytes -2: - CRC Checksum

	return f, nil
}

func crc(b []byte) []byte {
	s := crc16.New(&crc16.Conf{
		Poly:   0x8005,
		BitRev: true,
		IniVal: 0x0,
		FinVal: 0x0,
		BigEnd: false,
	})
	s.Write(b)
	return s.Sum(nil)
}
