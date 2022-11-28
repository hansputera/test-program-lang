package utils

import "fmt"

func BuildError(path string, line int, index int, message string) {
	fmt.Printf("Error at %s:%d:%d\n", path, line, index)
	fmt.Println("Message:", message)
}
