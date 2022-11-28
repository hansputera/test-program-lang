package runner

import (
	"strings"
	"testlang/globals"
	"testlang/structs"
	"testlang/utils"
)

func RunStruct(s *structs.ScriptStruct) {
	if len(s.Key) < 1 && len(s.Val) < 1 {
		return
	}

	switch s.Mode {
	case 1:
	case 0:
		utils.BuildError(s.Path, s.Location.Line, s.Location.CharIndex, "Missing semicolon!")
	case 3:
		s.Val = strings.TrimSpace(s.Val)
		// MEMO
		if s.Memo {
			if utils.IsStringInput(s.Val) {
				s.Val = utils.CleanupStr(s.Val)
			}
			globals.MemoizedVars[s.Key] = s.Val
		}
		// PRINT method
		if s.Key == "echo" {
			Echo(s)
		}
	}
}
