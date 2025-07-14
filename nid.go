package rfc8141

import (
	"github.com/reiver/go-utf8"
)

// PeekPrefixNID looks to see if `str` begins with an NID (Namespace Identifier),
// and if it does, then return its length.
//
// IETF RFC-8141 defines an NID (Namespace Identifier) as follows:
//
//	NID = (alphanum) 0*30(ldh) (alphanum)
//
//	ldh = alphanum / "-"
func PeekPrefixNID(str string) (n int, found bool) {

	if "" == str {
		return 0, false
	}

	const min int = 2
	const max int = 32

	var index int
	var r rune
	var prev rune

	const colon rune = ':'
	const lenColon int = len(string(colon))

	for index, r = range str {
		switch {
		case 0 == index:
			if !IsAlphaNum(r) {
				return 0, false
			}
			n += utf8.RuneLength(r)
			prev = r
		case colon == r:
			if index < min {
				return 0, false
			}
			if max < index {
				return 0, false
			}
			if !IsAlphaNum(prev) {
				return 0, false
			}
			return n, true
		default:
			if !IsAlphaNum(r) && '-' != r {
				return 0, false
			}
			n += utf8.RuneLength(r)
			prev = r
		}
	}

	return 0, false
}

// NormalizeNID returns the normalized form of an NID (Namespace Identifier).
//
// From IETF RFC-8141:
//
// "NIDs are case insensitive (e.g., "ISBN" and "isbn" are equivalent)."
//
// Normalization lower-cases chracter 'A' through 'Z' to 'a' through 'z'.
func NormalizeNID(nid string) string {
	var buffer [64]byte
	var p []byte = buffer[0:0]

	for _, r := range nid {
		if 'A' <= r && r <= 'Z' {
			r = r + ('a' - 'A')
		}
		p = append(p, string(r)...)
	}

	return string(p)
}
