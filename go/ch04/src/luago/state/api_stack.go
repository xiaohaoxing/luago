package state

func (self *luaState) GetTop() int {
	return self.stack.top
}

func (self *luaState) AbsIndex(idx int) int {
	return self.stack.absIndex(idx)
}

func (self *luaState) CheckStack(n int) bool {
	self.stack.check(n)
	return true
}

func (self *luaState) Pop(n int) {
	self.SetTop(-n - 1)
}

/*
*
将指定 idx 的数据复制到指定 idx
*/
func (self *luaState) Copy(fromIdx, toIdx int) {
	val := self.stack.get(fromIdx)
	self.stack.set(toIdx, val)
}

/*
*
将指定 idx 位置的数据写入到栈顶
*/
func (self *luaState) PushValue(idx int) {
	val := self.stack.get(idx)
	self.stack.push(val)
}

/*
*
将栈顶的数据弹出写入到指定 idx
*/
func (self *luaState) Replace(idx int) {
	val := self.stack.pop()
	self.stack.set(idx, val)
}

/*
*
将栈顶的数据弹出写入到指定 idx，并将后面所有数据向上移动一格（rotate 的一种特殊情况）
*/
func (self *luaState) Insert(idx int) {
	self.Rotate(idx, 1)
}

/*
*
删除指定位置的数据，并将后面所有数据都向下移动一格
*/
func (self *luaState) Remove(idx int) {
	self.Rotate(idx, -1)
	self.Pop(1)
}

func (self *luaState) Rotate(idx, n int) {
	t := self.stack.top - 1
	p := self.stack.absIndex(idx) - 1
	var m int
	if n >= 0 {
		m = t - n
	} else {
		m = p - n - 1
	}
	self.stack.reverse(p, m)
	self.stack.reverse(m+1, t)
	self.stack.reverse(p, t)

}

/*
*
设置栈的大小，如果 idx > top，则填充 nil，如果 idx < top， 则 pop 掉多余的元素
*/
func (self *luaState) SetTop(idx int) {
	newTop := self.stack.absIndex(idx)
	if newTop < 0 {
		panic("stack underflow!")
	}
	n := self.stack.top - newTop
	if n > 0 {
		for i := 0; i < n; i++ {
			self.stack.pop()
		}
	} else if n < 0 {
		for i := 0; i > n; i-- {
			self.stack.push(nil)
		}
	}
}
