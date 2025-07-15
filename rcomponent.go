package rfc8141

import (
	"strings"
)

// PeekPrefixRComponent looks to see if `str` begins with an r-component (without the "?+" prefix),
// and if it does, then returns the decoded r-component and its original (encoded) length.
//
// Notet that PeekPrefixRComponent expectd the "?+" to be there.
//
// IETF RFC-8141 defines an r-component as follows:
//
//	rq-components = [ "?+" r-component ]
//	                [ "?=" q-component ]
//	
//	r-component   = pchar *( pchar / "/" / "?" )
func PeekPrefixRComponent(str string) (rcomponent string, n int, found bool) {

	var s string = str

	{
		const prefix string = "?+"

		if !strings.HasPrefix(s, prefix) {
			var nada string = ""
			return nada, 0, false
		}

		s = s[len(prefix):]
	}

	var buffer [256]byte
	var p []byte = buffer[0:0]

	{
		b, length, found := PeekPrefixPChar(s)
		if !found {
			return "", 0, false
		}

		p = append(p, b)
		n += length
		s = s[length:]
	}

	const qComponentPrefix string = "?="

	loop: for {
		switch {
		case len(s) <= 0:
			break loop
		case strings.HasPrefix(s, qComponentPrefix):
			break loop
		case '/' == s[0]:
			length := len("/")

			p = append(p, '/')
			n += length
			s = s[length:]
			continue loop
		case '?' == s[0]:
			length := len("?")

			p = append(p, '?')
			n += length
			s = s[length:]
			continue loop
		default:
			b, length, found := PeekPrefixPChar(s)
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
