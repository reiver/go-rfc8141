package rfc8141_test

import (
	"testing"

	"strings"

	"github.com/reiver/go-rfc8141"
)

func TestPeekPrefixRComponent(t *testing.T) {

	tests := []struct{
		URN string
		ExpectedRComponent string
		ExpectedN int
	}{
		{
			URN: "urn:example:once/twice/thrice/fource?+ONE=1&TWO=22",
			ExpectedRComponent:                        "ONE=1&TWO=22",
			ExpectedN:                             len("ONE=1&TWO=22"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+ONE=1&TWO=22?=apple=11&banana=22",
			ExpectedRComponent:                        "ONE=1&TWO=22",
			ExpectedN:                             len("ONE=1&TWO=22"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+ONE=1&TWO=22#here",
			ExpectedRComponent:                        "ONE=1&TWO=22",
			ExpectedN:                             len("ONE=1&TWO=22"),
		},



		{
			URN: "urn:example:once/twice/thrice/fource?+q",
			ExpectedRComponent:                        "q",
			ExpectedN:                             len("q"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+q?=apple=11&banana=22",
			ExpectedRComponent:                        "q",
			ExpectedN:                             len("q"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+q#here",
			ExpectedRComponent:                        "q",
			ExpectedN:                             len("q"),
		},



		{
			URN: "urn:example:once/twice/thrice/fource?+q/",
			ExpectedRComponent:                        "q/",
			ExpectedN:                             len("q/"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+q/?=apple=11&banana=22",
			ExpectedRComponent:                        "q/",
			ExpectedN:                             len("q/"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+q/#here",
			ExpectedRComponent:                        "q/",
			ExpectedN:                             len("q/"),
		},



		{
			URN: "urn:example:once/twice/thrice/fource?+q/r",
			ExpectedRComponent:                        "q/r",
			ExpectedN:                             len("q/r"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+q/r?=apple=11&banana=22",
			ExpectedRComponent:                        "q/r",
			ExpectedN:                             len("q/r"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+q/r#here",
			ExpectedRComponent:                        "q/r",
			ExpectedN:                             len("q/r"),
		},



		{
			URN: "urn:example:once/twice/thrice/fource?+q/r/",
			ExpectedRComponent:                        "q/r/",
			ExpectedN:                             len("q/r/"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+q/r/?=apple=11&banana=22",
			ExpectedRComponent:                        "q/r/",
			ExpectedN:                             len("q/r/"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+q/r/#here",
			ExpectedRComponent:                        "q/r/",
			ExpectedN:                             len("q/r/"),
		},



		{
			URN: "urn:example:once/twice/thrice/fource?+q/r/s",
			ExpectedRComponent:                        "q/r/s",
			ExpectedN:                             len("q/r/s"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+q/r/s?=apple=11&banana=22",
			ExpectedRComponent:                        "q/r/s",
			ExpectedN:                             len("q/r/s"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+q/r/s#here",
			ExpectedRComponent:                        "q/r/s",
			ExpectedN:                             len("q/r/s"),
		},



		{
			URN: "urn:example:once/twice/thrice/fource?+q/r/s?",
			ExpectedRComponent:                        "q/r/s?",
			ExpectedN:                             len("q/r/s?"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+q/r/s??=apple=11&banana=22",
			ExpectedRComponent:                        "q/r/s?",
			ExpectedN:                             len("q/r/s?"),
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+q/r/s?#here",
			ExpectedRComponent:                        "q/r/s?",
			ExpectedN:                             len("q/r/s?"),
		},
	}

	for testNumber, test := range tests {

		var str string = test.URN[strings.Index(test.URN, "?+"):]

		actualRComponent, actualN, actualFound := rfc8141.PeekPrefixRComponent(str)
		if !actualFound {
			t.Errorf("For test #%d, expected to 'find' an r-component but actually did not.", testNumber)
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
			expected := test.ExpectedRComponent
			actual := actualRComponent

			if expected != actual {
				t.Errorf("For test #%d, the actual 'r-component' is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("URN:  %q", test.URN)
				continue
			}
		}
	}
}

func TestPeekPrefixRComponent_fail(t *testing.T) {

	tests := []struct{
		URN string
	}{
		{
			URN: "urn:example:once/twice/thrice/fource?+",
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+?=apple=11&banana=22",
		},
		{
			URN: "urn:example:once/twice/thrice/fource?+#here",
		},
	}

	for testNumber, test := range tests {

		var str string = test.URN[strings.Index(test.URN, "?+"):]

		actualNSS, actualN, actualFound := rfc8141.PeekPrefixRComponent(str)
		if actualFound {
			t.Errorf("For test #%d, did not expect to 'find' an r-component but actually did.", testNumber)
			t.Logf("URN:    %q", test.URN)
			t.Logf("ACTUAL-N: %d", actualN)
			t.Logf("ACTUAL-R-COMPONENT: %q (encoded)", str[:actualN])
			t.Logf("ACTUAL-R-COMPONENT: %q (decoded)", actualNSS)
			continue
		}
	}
}
