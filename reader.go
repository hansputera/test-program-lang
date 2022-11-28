package main

import (
	"bufio"
	"io"
	"os"
	"testlang/runner"
	"testlang/structs"
)

func ReadScript(path string) error {
	file, err := os.Open(path)
	tmp := &structs.ScriptStruct{
		Path: path,
	}
	tmp.Location = &structs.Loc{
		Line:      1,
		CharIndex: 0,
	}
	tmp.Mode = 0

	if err != nil {
		return err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		r, _, err := reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		char := string(r)

		tmp.Location.CharIndex++
		switch char {
		case "\n":
			runner.RunStruct(tmp)

			tmp.Key = ""
			tmp.Val = ""
			tmp.Mode = 0
			tmp.Stop = false
			tmp.Memo = false

			tmp.Location.Line++
			tmp.Location.CharIndex = 0
		case "#":
			runner.RunStruct(tmp)
			tmp.Stop = true
			continue
		case "(":
			tmp.Mode = 1
		case ")":
			tmp.Mode = 0
		case ";":
			tmp.Mode = 3
			runner.RunStruct(tmp)
			continue
		case ":":
			tmp.Memo = true
			continue
		default:
			if tmp.Stop {
				continue
			}

			if tmp.Mode == 1 || tmp.Memo {
				tmp.Val += char
			} else {
				tmp.Key += char
			}
		}
	}

	return nil
}
