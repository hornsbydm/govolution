// The Carrier Package provides a framework for communicating with
// Carrier Infinity and Bryant Evolution Furnaces.

// The Carrier Package provides a framework for communicating with
// Carrier Infinity and Bryant Evolution Furnaces.
package Carrier

type frameOpr uint8

const (
	OPRRESPONSE = iota
	OPRREAD
	OPRWRITE
	OPRERROR
)

func newFrameOpr(a uint8) frameOpr {
	switch a {
	case 0x06:
		return OPRRESPONSE
	case 0x0b:
		return OPRREAD
	case 0x0c:
		return OPRWRITE
	case 0x15:
		return OPRERROR
	}
	return 0x00
}

func (o frameOpr) Encode() uint8 {
	switch o {
	case OPRRESPONSE:
		return 0x06
	case OPRREAD:
		return 0x0b
	case OPRWRITE:
		return 0x0c
	case OPRERROR:
		return 0x15
	}
	return 0x00
}
