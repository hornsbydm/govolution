// Copyright 2022 David M. Hornsby.
// Use of this source code file is governed by a license that
// can be found in the LICENSE file
package CarrierTest

import (
	"bytes"
	"govolution/Carrier"
	"testing"
)

var frameEncodeCases = []struct {
	f *Carrier.Frame
	b []byte
	e error
}{
	{
		b: []byte{0x20, 0x01, 0x92, 0x01, 0x03, 0x00, 0x00, 0x0b, 0x00, 0x3b, 0x03, 0xec, 0xba},
	},
	{
		b: []byte{0x20, 0x01, 0x40, 0x01, 0x11, 0x00, 0x00, 0x06, 0x00, 0x03, 0x0a, 0x01, 0x0c, 0x0c, 0x02, 0x00, 0x1f, 0x00, 0x5d, 0x00, 0x00, 0x10, 0x10, 0x01, 0x14, 0x34, 0xf5},
	},
}

func TestFrame(t *testing.T) {
	for i, v := range frameEncodeCases {
		gotFrame, _ := Carrier.NewFrame(v.b)
		if gotBytes, _ := gotFrame.Encode(); !bytes.Equal(gotBytes, v.b) {
			t.Errorf("Failed Test #%v", i)

		}

	}
}
