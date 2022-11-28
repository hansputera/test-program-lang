package utils

import "strings"

func CleanupStr(str string) string {
	return strings.TrimSuffix(strings.TrimSuffix(strings.TrimPrefix(strings.TrimPrefix(str, "'"), "\""), "'"), "\"")
}
