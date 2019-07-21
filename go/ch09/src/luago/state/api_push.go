package state

import . "luago/api"

func (self *luaState) PushNil()                    { self.stack.push(nil) }
func (self *luaState) PushBoolean(b bool)          { self.stack.push(b) }
func (self *luaState) PushInteger(n int64)         { self.stack.push(n) }
func (self *luaState) PushNumber(n float64)        { self.stack.push(n) }
func (self *luaState) PushString(s string)         { self.stack.push(s) }
func (self *luaState) PushGoFunction(f GoFunction) { self.stack.push(newGoClosure(f)) }

func (self *luaState) PushGlobalTable() {
	global := self.registry.get(LUA_RIDX_GLOBALS)
	self.stack.push(global)
}
