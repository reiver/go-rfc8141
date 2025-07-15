package rfc8141_test

import (
	"testing"

	"strings"

	"github.com/reiver/go-rfc8141"
)

func TestPeekPrefixQComponent(t *testing.T) {

	tests := []struct{
		URN string
		ExpectedQComponent string
		ExpectedN int
	}{
		{
			URN: "urn:example:once/twice/thrice/fource?=ONE=1&TWO=22",
			ExpectedQComponent:                        "ONE=1&TWO=22",
			ExpectedN:                             len("ONE=1&TWO=22"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+ONE=1&TWO=22?=apple=11&banana=22",
			ExpectedQComponent:                                      "apple=11&banana=22",
			ExpectedN:                                           len("apple=11&banana=22"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?=ONE=1&TWO=22#here",
			ExpectedQComponent:                        "ONE=1&TWO=22",
			ExpectedN:                             len("ONE=1&TWO=22"),
		},



		{
			URN: "urn:example:once/twice/thrice/fource?=q",
			ExpectedQComponent:                        "q",
			ExpectedN:                             len("q"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+query?=w",
			ExpectedQComponent:                               "w",
			ExpectedN:                                    len("w"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?=q#here",
			ExpectedQComponent:                        "q",
			ExpectedN:                             len("q"),
		},



		{
			URN: "urn:example:once/twice/thrice/fource?=q/",
			ExpectedQComponent:                        "q/",
			ExpectedN:                             len("q/"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+query?=q/",
			ExpectedQComponent:                               "q/",
			ExpectedN:                                    len("q/"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?=q/#here",
			ExpectedQComponent:                        "q/",
			ExpectedN:                             len("q/"),
		},



		{
			URN: "urn:example:once/twice/thrice/fource?=q/r",
			ExpectedQComponent:                        "q/r",
			ExpectedN:                             len("q/r"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+query?=q/r",
			ExpectedQComponent:                               "q/r",
			ExpectedN:                                    len("q/r"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?=q/r#here",
			ExpectedQComponent:                        "q/r",
			ExpectedN:                             len("q/r"),
		},



		{
			URN: "urn:example:once/twice/thrice/fource?=q/r/",
			ExpectedQComponent:                        "q/r/",
			ExpectedN:                             len("q/r/"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+query?=q/r/",
			ExpectedQComponent:                               "q/r/",
			ExpectedN:                                    len("q/r/"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?=q/r/#here",
			ExpectedQComponent:                        "q/r/",
			ExpectedN:                             len("q/r/"),
		},



		{
			URN: "urn:example:once/twice/thrice/fource?=q/r/s",
			ExpectedQComponent:                        "q/r/s",
			ExpectedN:                             len("q/r/s"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+query?=q/r/s",
			ExpectedQComponent:                               "q/r/s",
			ExpectedN:                                    len("q/r/s"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?=q/r/s#here",
			ExpectedQComponent:                        "q/r/s",
			ExpectedN:                             len("q/r/s"),
		},



		{
			URN: "urn:example:once/twice/thrice/fource?=q/r/s?",
			ExpectedQComponent:                        "q/r/s?",
			ExpectedN:                             len("q/r/s?"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+query?=q/r/s?",
			ExpectedQComponent:                               "q/r/s?",
			ExpectedN:                                    len("q/r/s?"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?=q/r/s?#here",
			ExpectedQComponent:                        "q/r/s?",
			ExpectedN:                             len("q/r/s?"),
		},
	}

	for testNumber, test := range tests {

		var str string = test.URN[strings.Index(test.URN, "?="):]

		actualQComponent, actualN, actualFound := rfc8141.PeekPrefixQComponent(str)
		if !actualFound {
			t.Errorf("For test #%d, expected to 'find' an q-component but actually did not.", testNumber)
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
			expected := test.ExpectedQComponent
			actual := actualQComponent

			if expected != actual {
				t.Errorf("For test #%d, the actual 'q-component' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("URN:  %q", test.URN)
				continue
			}
		}
	}
}

func TestPeekPrefixQComponent_fail(t *testing.T) {

	tests := []struct{
		URN string
	}{
		{
			URN: "urn:example:once/twice/thrice/fource?=",
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+?=",
		},
		{
			URN: "urn:example:once/twice/thrice/fource?=#here",
		},
	}

	for testNumber, test := range tests {

		var str string = test.URN[strings.Index(test.URN, "?="):]

		actualNSS, actualN, actualFound := rfc8141.PeekPrefixQComponent(str)
		if actualFound {
			t.Errorf("For test #%d, did not expect to 'find' an q-component but actually did.", testNumber)
			t.Logf("URN:    %q", test.URN)
			t.Logf("ACTUAL-N: %d", actualN)
			t.Logf("ACTUAL-Q-COMPONENT: %q (encoded)", str[:actualN])
			t.Logf("ACTUAL-Q-COMPONENT: %q (decoded)", actualNSS)
			continue
		}
	}
}
