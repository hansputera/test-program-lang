package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Printf("Usage: %s <script_path>\n", os.Args[0])
		os.Exit(1)
	}

	err := ReadScript(args[0])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
