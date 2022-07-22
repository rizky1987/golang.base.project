package commonHelpers

import (
	"strings"
)

func StringSplitToArrayString(inputString, delimiter string) []string {
	return strings.Split(TrimWhiteSpace(inputString), TrimWhiteSpace(delimiter))
}

func TrimWhiteSpace(inputString string) string {
	return strings.TrimSpace(inputString)
}

func StringReplace(original, from, to string) string {
	return strings.ReplaceAll(original, from, to)
}
