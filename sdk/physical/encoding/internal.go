package encoding

import (
	"strings"
	"unicode"
)

func (e *StorageEncoding) containsNonPrintableChars(key string) bool {
	idx := strings.IndexFunc(key, func(c rune) bool {
		return !unicode.IsPrint(c)
	})
	return idx != -1
}
