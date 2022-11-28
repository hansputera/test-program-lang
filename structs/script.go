package structs

type Loc struct {
	Line      int
	CharIndex int
}

type ScriptStruct struct {
	Path string
	Key  string
	Val  string
	Stop bool
	Memo bool

	// op state
	Mode uint8

	// Mode 1 = VALUE
	// Mode 0 = NORMAL
	// Mode 3 = EOS (End Of Script)

	Location *Loc
}
