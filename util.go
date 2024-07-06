package fmtlimit

import (
	"unicode/utf8"
)

type limitReached struct{}

func truncateString(b []byte, bytesLimit int) string {
	if bytesLimit <= 0 || len(b) == 0 {
		return ""
	}

	if bytesLimit >= len(b) {
		return string(b)
	}

	for i := bytesLimit; i > 0; i-- {
		r, _ := utf8.DecodeLastRune(b[:i])
		if r != utf8.RuneError {
			return string(b[:i])
		}
	}

	return ""
}
