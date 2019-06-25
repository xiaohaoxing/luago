package vm

// 编码模式：
const (
	IABC = iota
	IABx
	IAsBx
	IAx
)

// 共 46 个指令：
const (
	OP_MOVE = iota
	OP_LOCDK
	OP_LOADKX
	OP_LOADBOOL
	OP_LOADNIL
	OP_GETUPVAL
	OP_GETTABUP
	OP_GETTABLE
	OP_SETTABUP
	OP_SETUPVAL
	OP_SETTABLE
	OP_NEWTABLE
	OP_SELF
	OP_ADD
	OP_SUB
	OP_MUL
	OP_MOD
	OP_POW
	OP_DIV
	OP_IDIV
	OP_BAND
	OP_BOR
	OP_BXOR
	OP_SHL
	OP_SHR
	OP_UNM
	OP_BNOT
	OP_NOT
	OP_LEN
	OP_CONCAT
	OP_JMP
	OP_EQ
	OP_LT
	OP_LE
	OP_TEST
	OP_TESTSET
	OP_CALL
	OP_TAILCALL
	OP_RETURN
	OP_FORLOOP
	OP_FORPREP
	OP_TFORCALL
	OP_TFORLOOP
	OP_SETLIST
	OP_CLOSURE
	OP_VARARG
	OP_EXTRAARG
)

const (
	OpArgN = iota
	OpArgU
	OpArgR
	OpArgK
)

type opcode struct {
	testFlag byte
	setAFlag byte
	argBMode byte
	argCMode byte
	opMode   byte
	name     string
}

var opcodes = []opcode{
	{0, 1, OpArgR, OpArgN, IABC, "MOVE    "},
	{0, 1, OpArgK, OpArgN, IABx, "LOADK   "},
	{0, 1, OpArgN, OpArgN, IABx, "LOADKX  "},
	{0, 1, OpArgU, OpArgU, IABC, "LOADBOOL"},
	{0, 1, OpArgU, OpArgN, IABC, "LOADNIL "},
	{0, 1, OpArgU, OpArgN, IABC, "GETUPVAL"},
	{0, 1, OpArgU, OpArgK, IABC, "GETTABUP"},
	{0, 1, OpArgR, OpArgK, IABC, "GETTABLE"},
	{0, 0, OpArgK, OpArgK, IABC, "SETTABUP"},
	{0, 0, OpArgU, OpArgN, IABC, "SETUPVAL"},
	{0, 0, OpArgK, OpArgK, IABC, "SETTABLE"},
	{0, 1, OpArgU, OpArgU, IABC, "NEWTABLE"},
	{0, 1, OpArgR, OpArgK, IABC, "SELF    "},
	{0, 1, OpArgK, OpArgK, IABC, "ADD     "},
	{0, 1, OpArgK, OpArgK, IABC, "SUB     "},
	{0, 1, OpArgK, OpArgK, IABC, "MUL     "},
	{0, 1, OpArgK, OpArgK, IABC, "MOD     "},
	{0, 1, OpArgK, OpArgK, IABC, "POW     "},
	{0, 1, OpArgK, OpArgK, IABC, "DIV     "},
	{0, 1, OpArgK, OpArgK, IABC, "IDIV    "},
	{0, 1, OpArgK, OpArgK, IABC, "BAND    "},
	{0, 1, OpArgK, OpArgK, IABC, "BOR     "},
	{0, 1, OpArgK, OpArgK, IABC, "BXOR    "},
	{0, 1, OpArgK, OpArgK, IABC, "SHL     "},
	{0, 1, OpArgK, OpArgK, IABC, "SHR     "},
	{0, 1, OpArgR, OpArgN, IABC, "UNM     "},
	{0, 1, OpArgR, OpArgN, IABC, "BNOT    "},
	{0, 1, OpArgR, OpArgN, IABC, "NOT     "},
	{0, 1, OpArgR, OpArgN, IABC, "LEN     "},
	{0, 1, OpArgR, OpArgR, IABC, "CONCAT  "},
	{0, 0, OpArgR, OpArgN, IAsBx, "JMP     "},
	{1, 0, OpArgK, OpArgK, IABC, "EQ      "},
	{1, 0, OpArgK, OpArgK, IABC, "LT      "},
	{1, 0, OpArgK, OpArgK, IABC, "LE      "},
	{1, 0, OpArgN, OpArgU, IABC, "TEST    "},
	{1, 1, OpArgR, OpArgU, IABC, "TESTSET "},
	{0, 1, OpArgU, OpArgU, IABC, "CALL    "},
	{0, 1, OpArgU, OpArgU, IABC, "TAILCALL"},
	{0, 0, OpArgU, OpArgN, IABC, "RETURN  "},
	{0, 1, OpArgR, OpArgN, IAsBx, "FORLOOP "},
	{0, 1, OpArgR, OpArgN, IAsBx, "FORPREP "},
	{0, 0, OpArgN, OpArgU, IABC, "TFORCALL"},
	{0, 1, OpArgR, OpArgN, IAsBx, "TFORLOOP"},
	{0, 0, OpArgU, OpArgU, IABC, "SETLIST "},
	{0, 1, OpArgU, OpArgN, IABx, "CLOSURE "},
	{0, 1, OpArgU, OpArgN, IABC, "VARARG  "},
	{0, 0, OpArgU, OpArgU, IAx, "EXTRAARG"},
}
