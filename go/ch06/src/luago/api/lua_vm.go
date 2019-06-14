package api

type LuaVM interface {
	LuaState // 拓展了 LuaState
	// 新增了一些函数
	PC() int          // 返回当前 PC（仅用于测试）
	AddPC(n int)      // 修改当前 PC（用于实现跳转）
	Fetch() uint32    // 取出当前指令，PC 指向下一条指令
	GetConst(idx int) // 常量推入栈顶
	GetRK(rk int)     // 常量或值推入栈顶
}
