package binchunk

import "encoding/binary"
import "math"

type reader struct {
	data []byte
}
// 读取一个字节
func (self *reader) readByte() byte {
	b := self.data[0]
	self.data = self.data[1:]
	return b
}

func (self *reader) readUint32() uint32 {
	i := binary.LittleEndian.Uint32(self.data)
	self.data = self.data[4:]
	return i
}

func (self *reader) readUint64() uint64 {
	i := binary.LittleEndian.Uint64(self.data)
	self.data = self.data[8:]
	return i
}

func (self *reader) readLuaInteger() int64 {
	return int64(self.readUint64())
}

func (self *reader) readLuaNumber() float64 {
	return math.Float64frombits(self.readUint64())
}

func (self *reader) readString() string {
	size := uint(self.readByte())
	if size == 0 {
		return ""
	}
	if size == 0xFF {
		size = uint(self.readUint64())
	}
	bytes := self.readBytes(size - 1)
	return string(bytes)
}

func (self *reader) readBytes(n uint) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}

// 读取头部信息，并进行核对，异常则终止运行
func (self *reader) checkHeader() {
	if string(self.readBytes(4)) != LUA_SIGNATURE {
		panic("not a precompiled chunk!")
	}else if self.readByte() != LUAC_VERSION {
		panic("version mismatch!")
	}else if self.readByte() != LUAC_FORMAT {
		panic("format mismatch!")
	}else if string(self.readBytes(6)) != LUAC_DATA {
		panic("corrupted!")
	}else if self.readByte() != CINT_SIZE {
		panic("int size mismatch!")
	}else if self.readByte() != CSIZET_SIZE {
		panic("size_t size mismatch!")
	}else if self.readByte() != INSTRUCTION_SIZE {
		panic("instruction size mismatch!")
	}else if self.readByte() != LUA_INTEGER_SIZE {
		panic("lua_Integer size mismatch!")
	}else if self.readByte() != LUA_NUMBER_SIZE {
		panic("lua_Number size mismatch!")
	}else if self.readLuaInteger() != LUAC_INT {
		panic("endianness mismatch!")
	}else if self.readLuaNumber() != LUAC_NUM {
		panic("float format mismatch!")
	}
}

func (self *reader) readProto(parentSource string) *Prototype {
	source := self.readString()
	// 只有主函数有文件名，因此所有的函数文件名需要从父原型中找。
	if source == "" {
		source = parentSource
	}
	return &Prototype {
		Source: source,
		LineDefined: self.readUint32(),
		LastLineDefined: self.readUint32(),
		NumParams: self.readByte(),
		IsVararg: self.readByte(),
		MaxStackSize: self.readByte(),
		Code: self.readCode(),
		Constants: self.readConstants(),
		Upvalues: self.readUpvalues(),
		Protos: self.readProtos(source),
		LineInfo: self.readLineInfo(),
		LocVars: self.readLocVars(),
		UpvalueNames: self.readUpvalueNames(),
	}
}

func (self *reader) readCode() []uint32 {
	code := make([]uint32, self.readUint32())
	for i := range code {
		code[i] = self.readUint32()
	}
	return code
}

func (self *reader) readConstants() []interface{} {
	constants := make([]interface{}, self.readUint32())
	for i := range constants {
		constants[i] = self.readConstant()
	}
	return constants
}

func (self *reader) readConstant() interface{} {
	switch self.readByte() {
	case TAG_NIL: return nil
	case TAG_BOOLEAN: return self.readByte()!= 0
	case TAG_INTEGER: return self.readLuaInteger
	case TAG_NUMBER: return self.readLuaNumber
	case TAG_SHORT_STR: return self.readString()
	case TAG_LONG_STR: return self.readString()
	default: panic("corrupted!")
	}
}

func (self *reader) readUpvalues() []Upvalue {
	upvalues := make([]Upvalue, self.readUint32())
	for i := range upvalues {
		upvalues[i] = Upvalue {
			Instack: self.readByte(),
			Idx: self.readByte(),
		}
	}
	return upvalues
}

func (self *reader) readProtos(parentSource string) []*Prototype {
	protos := make([]*Prototype, self.readUint32())
	for i := range protos {
		protos[i] = self.readProto(parentSource)
	}
	return protos
}

func (self *reader) readLineInfo() []uint32 {
	lineInfo := make([]uint32, self.readUint32())
	for i := range lineInfo {
		lineInfo[i] = self.readUint32()
	}
	return lineInfo
}

func (self *reader) readLocVars() []LocVar {
	locVars := make([]LocVar, self.readUint32())
	for i := range locVars {
		locVars[i] = LocVar{ 
			VarName: self.readString(),
			StartPC: self.readUint32(),
			EndPC: self.readUint32(),
		}
	}
	return locVars
}

func (self *reader) readUpvalueNames() []string {
	names := make([]string, self.readUint32())
	for i := range names {
		names[i] = self.readString()
	}
	return names
}

