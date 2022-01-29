// Copyright 2022 David M. Hornsby.
// Use of this source code file is governed by a license that
// can be found in the LICENSE file.
package mqtt

// MQTTAction is an enum type that describes the use of the action in mqtt messages.

type MQTTAction uint8

const (
	POLL MQTTAction = iota
	READ
	WRITE
)

func (a MQTTAction) String() string {
	switch a {
	case POLL:
		return "POLL"
	case READ:
		return "READ"
	case WRITE:
		return "WRITE"
	}
	return "unknown"
}
