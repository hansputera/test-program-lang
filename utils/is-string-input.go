package utils

func IsStrSym(sym byte) bool {
	return string(sym) == "'" || string(sym) == "\""
}

func IsStringInput(str string) bool {
	return IsStrSym(str[0]) && IsStrSym(str[len(str)-1])
}
