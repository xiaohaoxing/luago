package state

import "luago/binchunk"
import . "luago/api"

type closure struct {
	proto  *binchunk.Prototype
	goFunc GoFunction
}

func newLuaClosure(proto *binchunk.Prototype) *closure {
	return &closure{proto: proto}
}

func newGoClosure(f GoFunction) *closure {
	return &closure{goFunc: f}
}
