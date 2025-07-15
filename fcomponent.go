package rfc8141

import (
	"github.com/reiver/go-rfc3986"
)

// PeekPrefixFComponent looks at `str` and if it starts (and ends) with
// a fragment, as defined by IETF RFC-3986, then it returns it
// (without the '#' prefix).
//
// PeekPrefixFComponent expects `str` to start with a '#', but the '#' is
// not returns as part of `fragment`. `n` includes the length of '#'.
//
// IETF RFC-3986 defines 'fragment' as follows:
//
//      A fragment identifier component is indicated by
//      the presence of a number sign ("#") character and
//      terminated by the end of the URI.
//
//      fragment = *( pchar / "/" / "?" )
func PeekPrefixFComponent(str string) (fcomponent string, n int, found bool) {
	return rfc3986.PeekPrefixFragment(str)
}
