package vm

import . "luago/api"

func loadNil(i Instruction, vm LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	vm.PushNil()
	for i := a; i < a+b; i++ {
		vm.Copy(-1, i)
	}
	vm.Pop(1)
}
