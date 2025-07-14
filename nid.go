package rfc8141

import (
	"github.com/reiver/go-utf8"
)

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
