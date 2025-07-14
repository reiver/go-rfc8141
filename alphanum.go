package rfc8141

import (
	"github.com/reiver/go-rfc3986"
)

// IsAlphaNum return 'true' if the rune is an "alphanum"
// as defined by IETF RFC-8141, and returns 'false' otherwise.
//
// IETF RFC-8141 imports the definition of "alphanum" from IETF RFC-3986.
//
// For example:
//
//	result := rfc8141.IsAlphaNum("E")
//	// result == true
//
// And also, for example:
//
//	result := rfc8141.IsAlphaNum("-")
//	// result == false
func IsAlphaNum(r rune) bool {
	return rfc3986.IsAlphaNum(r)
}
