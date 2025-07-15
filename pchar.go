package rfc8141

import (
	"github.com/reiver/go-rfc3986"
)

// PeekPrefixPChar peeks a string, checks to see if it starts with a 'pchar' (as defined in IETF RFC-3986),
// and returns it if it does.
//
// 'pchar' is defined in IETF RFC-3986 as:
//
//	pchar = unreserved / pct-encoded / sub-delims / ":" / "@"
func PeekPrefixPChar(str string) (byte, int, bool) {
	return rfc3986.PeekPrefixPChar(str)
}
