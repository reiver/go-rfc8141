package rfc8141

import (
	"github.com/reiver/go-rfc3986"
)

// PeekPrefixNSS looks to see if `str` begins with an NSS (Namespace Specific String),
// and if it does, then returns the decoded NSS (Namespace Specific String) its original (encoded) length.
//
// IETF RFC-8141 defines an NSS (Namespace Specific String) as follows:
//
//	NSS = pchar *(pchar / "/")
func PeekPrefixNSS(str string) (nss string, n int, found bool) {

	if "" == str {
		return "", 0, false
	}

	var buffer [256]byte
	var p []byte = buffer[0:0]

	var s string = str

	{
		b, length, found := rfc3986.PeekPrefixPChar(s)
		if !found {
			return "", 0, false
		}

		p = append(p, b)
		n += length
		s = s[length:]
	}

	loop: for {
		switch {
		case len(s) <= 0:
			break loop
		case '/' == s[0]:
			length := len("/")

			p = append(p, '/')
			n += length
			s = s[length:]
			continue loop
		default:
			b, length, found := rfc3986.PeekPrefixPChar(s)
			if !found {
				break loop
			}

			p = append(p, b)
			n += length
			s = s[length:]
			continue loop
		}
	}

	return string(p), n, true
}
