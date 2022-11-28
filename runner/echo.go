package runner

import (
	"fmt"
	"testlang/globals"
	"testlang/structs"
	"testlang/utils"

	"github.com/maja42/goval"
)

func Echo(s *structs.ScriptStruct) {
	if utils.IsStringInput(s.Val) {
		PrintMethod(utils.CleanupStr(s.Val))
	} else if s.Val == "true" || s.Val == "false" {
		PrintMethod(fmt.Sprintf("Boolean: %s", s.Val))
	} else {
		eval, err := goval.NewEvaluator().Evaluate(s.Val, globals.MemoizedVars, nil)
		if err != nil {
			utils.BuildError(s.Path, s.Location.Line, s.Location.CharIndex, err.Error())
		} else {
			fmt.Println(eval)
		}
	}
}
