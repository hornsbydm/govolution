// Copyright 2022 David M. Hornsby.
// Use of this source code file is governed by a license that
// can be found in the LICENSE file.

// The Carrier Package provides a framework for communicating with
// Carrier Infinity and Bryant Evolution Furnaces.
package Carrier

var frameHandlerRegistry []FrameHandler

type FrameHandler interface {
	Match(f *Frame) bool
	Handle(f *Frame)
}

func RegisterFrameHandler(fh FrameHandler) {
	frameHandlerRegistry = append(frameHandlerRegistry, fh)
}

func dispatchFrame(f *Frame) {
	for _, v := range frameHandlerRegistry {
		if v.Match(f) {
			go v.Handle(f)
		}
	}
}
