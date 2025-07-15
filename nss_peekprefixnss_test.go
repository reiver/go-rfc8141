package rfc8141_test

import (
	"testing"

	"strings"

	"github.com/reiver/go-rfc8141"
)

func TestPeekPrefixNSS(t *testing.T) {

	tests := []struct{
		URN string
		ExpectedNSS string
		ExpectedN int
	}{
		{
			URN:     "urn:example:something",
			ExpectedNSS:         "something",
			ExpectedN:       len("something"),
		},



		{
			URN:     "urn:ietf:rfc:8141",
			ExpectedNSS:      "rfc:8141",
			ExpectedN:    len("rfc:8141"),
		},



		{
			URN:     "urn:sha1:FINYVGHENTHSMNDSQQYDNLPONVBZTICF",
			ExpectedNSS:      "FINYVGHENTHSMNDSQQYDNLPONVBZTICF",
			ExpectedN:    len("FINYVGHENTHSMNDSQQYDNLPONVBZTICF"),
		},



		{
			URN:     "urn:uuid:347a4f5f-9a01-4d56-9a45-86cce48e60c9",
			ExpectedNSS:      "347a4f5f-9a01-4d56-9a45-86cce48e60c9",
			ExpectedN:    len("347a4f5f-9a01-4d56-9a45-86cce48e60c9"),
		},



		{
			URN:     "urn:xmpp:joeblow",
			ExpectedNSS:      "joeblow",
			ExpectedN:    len("joeblow"),
		},



		{
			URN:     "urn:0123456789aBcDeFgHiJkLmNoPqRsTuV:Example",
			ExpectedNSS:                                  "Example",
			ExpectedN:                                len("Example"),
		},



		{
			URN:     "urn:0------------------------------V:Example",
			ExpectedNSS:                                  "Example",
			ExpectedN:                                len("Example"),
		},



		{
			URN:     "urn:example:%2F",
			ExpectedNSS:         "/",
			ExpectedN:       len("%2F"),
		},
		{
			URN:     "urn:example:%F0%9F%98%88",
			ExpectedNSS:         "😈",
			ExpectedN:       len("%F0%9F%98%88"),
		},
		{
			URN:     "urn:example:Hello%20world!%20%F0%9F%99%82",
			ExpectedNSS:         "Hello world! 🙂",
			ExpectedN:       len("Hello%20world!%20%F0%9F%99%82"),
		},



		{
			URN:     "urn:example:!",
			ExpectedNSS:         "!",
			ExpectedN:       len("!"),
		},
		{
			URN:     "urn:example:$",
			ExpectedNSS:         "$",
			ExpectedN:       len("$"),
		},
		{
			URN:     "urn:example:&",
			ExpectedNSS:         "&",
			ExpectedN:       len("&"),
		},
		{
			URN:     "urn:example:'",
			ExpectedNSS:         "'",
			ExpectedN:       len("'"),
		},
		{
			URN:     "urn:example:(",
			ExpectedNSS:         "(",
			ExpectedN:       len("("),
		},
		{
			URN:     "urn:example:)",
			ExpectedNSS:         ")",
			ExpectedN:       len(")"),
		},
		{
			URN:     "urn:example:*",
			ExpectedNSS:         "*",
			ExpectedN:       len("*"),
		},
		{
			URN:     "urn:example:+",
			ExpectedNSS:         "+",
			ExpectedN:       len("+"),
		},
		{
			URN:     "urn:example:,",
			ExpectedNSS:         ",",
			ExpectedN:       len(","),
		},
		{
			URN:     "urn:example:;",
			ExpectedNSS:         ";",
			ExpectedN:       len(";"),
		},
		{
			URN:     "urn:example:=",
			ExpectedNSS:         "=",
			ExpectedN:       len("="),
		},



		{
			URN:     "urn:example::",
			ExpectedNSS:         ":",
			ExpectedN:       len(":"),
		},
		{
			URN:     "urn:example:@",
			ExpectedNSS:         "@",
			ExpectedN:       len("@"),
		},



		{
			URN:     "urn:example:once",
			ExpectedNSS:         "once",
			ExpectedN:       len("once"),
		},
		{
			URN:     "urn:example:once/twice",
			ExpectedNSS:         "once/twice",
			ExpectedN:       len("once/twice"),
		},
		{
			URN:     "urn:example:once/twice/thrice",
			ExpectedNSS:         "once/twice/thrice",
			ExpectedN:       len("once/twice/thrice"),
		},
		{
			URN:     "urn:example:once/twice/thrice/fource",
			ExpectedNSS:         "once/twice/thrice/fource",
			ExpectedN:       len("once/twice/thrice/fource"),
		},
	}

	for testNumber, test := range tests {

		const prefix string = rfc8141.Scheme+":" // "urn:"
		var str string = test.URN[len(prefix):]
		str = str[1+strings.Index(str, ":"):]

		actualNSS, actualN, actualFound := rfc8141.PeekPrefixNSS(str)
		if !actualFound {
			t.Errorf("For test #%d, expected to 'find' an NSS but actually did not.", testNumber)
			t.Logf("URN: %q", test.URN)
			continue
		}

		{
			expected := test.ExpectedN
			actual := actualN

			if expected != actual {
				t.Errorf("For test #%d, the actual 'n' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				t.Logf("URN: %q", test.URN)
				continue
			}
		}

		{
			expected := test.ExpectedNSS
			actual := actualNSS

			if expected != actual {
				t.Errorf("For test #%d, the actual 'NID' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("URN:  %q", test.URN)
				continue
			}
		}
	}
}

func TestPeekPrefixNSS_fail(t *testing.T) {

	tests := []struct{
		URN string
	}{

		{
			URN: "urn:example:/",
		},
		{
			URN: "urn:example:/path",
		},
		{
			URN: "urn:example:/path/to",
		},
		{
			URN: "urn:example:/path/to/it",
		},



		{
			URN: "urn:example:😈",
		},
	}

	for testNumber, test := range tests {

		const prefix string = rfc8141.Scheme+":" // "urn:"
		var str string = test.URN[len(prefix):]
		str = str[1+strings.Index(str, ":"):]

		actualNSS, actualN, actualFound := rfc8141.PeekPrefixNSS(str)
		if actualFound {
			t.Errorf("For test #%d, did not expect to 'find' an NSS but actually did.", testNumber)
			t.Logf("URN:    %q", test.URN)
			t.Logf("ACTUAL-N: %d", actualN)
			t.Logf("ACTUAL-NSS: %q (encoded)", str[:actualN])
			t.Logf("ACTUAL-NSS: %q (decoded)", actualNSS)
			continue
		}
	}
}
