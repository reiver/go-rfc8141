package rfc8141_test

import (
	"testing"

	"github.com/reiver/go-rfc8141"
)

func TestPeekPrefixNID(t *testing.T) {

	tests := []struct{
		URN string
		ExpectedNID string
	}{
		{
			URN:     "urn:example:something",
			ExpectedNID: "example",
		},



		{
			URN:     "urn:ietf:rfc:8141",
			ExpectedNID: "ietf",
		},



		{
			URN:     "urn:sha1:FINYVGHENTHSMNDSQQYDNLPONVBZTICF",
			ExpectedNID: "sha1",
		},



		{
			URN:     "urn:uuid:347a4f5f-9a01-4d56-9a45-86cce48e60c9",
			ExpectedNID: "uuid",
		},



		{
			URN:     "urn:xmpp:joeblow",
			ExpectedNID: "xmpp",
		},
		{
			URN:     "urn:xmpP:joeblow",
			ExpectedNID: "xmpP",
		},
		{
			URN:     "urn:xmPp:joeblow",
			ExpectedNID: "xmPp",
		},
		{
			URN:     "urn:xmPP:joeblow",
			ExpectedNID: "xmPP",
		},
		{
			URN:     "urn:xMpp:joeblow",
			ExpectedNID: "xMpp",
		},
		{
			URN:     "urn:xMpP:joeblow",
			ExpectedNID: "xMpP",
		},
		{
			URN:     "urn:xMPp:joeblow",
			ExpectedNID: "xMPp",
		},
		{
			URN:     "urn:xMPP:joeblow",
			ExpectedNID: "xMPP",
		},
		{
			URN:     "urn:Xmpp:joeblow",
			ExpectedNID: "Xmpp",
		},
		{
			URN:     "urn:XmpP:joeblow",
			ExpectedNID: "XmpP",
		},
		{
			URN:     "urn:XmPp:joeblow",
			ExpectedNID: "XmPp",
		},
		{
			URN:     "urn:XmPP:joeblow",
			ExpectedNID: "XmPP",
		},
		{
			URN:     "urn:XMpp:joeblow",
			ExpectedNID: "XMpp",
		},
		{
			URN:     "urn:XMpP:joeblow",
			ExpectedNID: "XMpP",
		},
		{
			URN:     "urn:XMPp:joeblow",
			ExpectedNID: "XMPp",
		},
		{
			URN:     "urn:XMPP:joeblow",
			ExpectedNID: "XMPP",
		},



		{
			URN:     "urn:0:Example",
			ExpectedNID: "0",
		},
		{
			URN:     "urn:01:Example",
			ExpectedNID: "01",
		},
		{
			URN:     "urn:012:Example",
			ExpectedNID: "012",
		},
		{
			URN:     "urn:0123:Example",
			ExpectedNID: "0123",
		},
		{
			URN:     "urn:01234:Example",
			ExpectedNID: "01234",
		},
		{
			URN:     "urn:012345:Example",
			ExpectedNID: "012345",
		},
		{
			URN:     "urn:0123456:Example",
			ExpectedNID: "0123456",
		},
		{
			URN:     "urn:01234567:Example",
			ExpectedNID: "01234567",
		},
		{
			URN:     "urn:012345678:Example",
			ExpectedNID: "012345678",
		},
		{
			URN:     "urn:0123456789:Example",
			ExpectedNID: "0123456789",
		},
		{
			URN:     "urn:0123456789a:Example",
			ExpectedNID: "0123456789a",
		},
		{
			URN:     "urn:0123456789aB:Example",
			ExpectedNID: "0123456789aB",
		},
		{
			URN:     "urn:0123456789aBc:Example",
			ExpectedNID: "0123456789aBc",
		},
		{
			URN:     "urn:0123456789aBcD:Example",
			ExpectedNID: "0123456789aBcD",
		},
		{
			URN:     "urn:0123456789aBcDe:Example",
			ExpectedNID: "0123456789aBcDe",
		},
		{
			URN:     "urn:0123456789aBcDeF:Example",
			ExpectedNID: "0123456789aBcDeF",
		},
		{
			URN:     "urn:0123456789aBcDeFg:Example",
			ExpectedNID: "0123456789aBcDeFg",
		},
		{
			URN:     "urn:0123456789aBcDeFgH:Example",
			ExpectedNID: "0123456789aBcDeFgH",
		},
		{
			URN:     "urn:0123456789aBcDeFgHi:Example",
			ExpectedNID: "0123456789aBcDeFgHi",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJ:Example",
			ExpectedNID: "0123456789aBcDeFgHiJ",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJk:Example",
			ExpectedNID: "0123456789aBcDeFgHiJk",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkL:Example",
			ExpectedNID: "0123456789aBcDeFgHiJkL",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLm:Example",
			ExpectedNID: "0123456789aBcDeFgHiJkLm",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLmN:Example",
			ExpectedNID: "0123456789aBcDeFgHiJkLmN",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLmNo:Example",
			ExpectedNID: "0123456789aBcDeFgHiJkLmNo",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLmNoP:Example",
			ExpectedNID: "0123456789aBcDeFgHiJkLmNoP",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLmNoPq:Example",
			ExpectedNID: "0123456789aBcDeFgHiJkLmNoPq",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLmNoPqR:Example",
			ExpectedNID: "0123456789aBcDeFgHiJkLmNoPqR",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLmNoPqRs:Example",
			ExpectedNID: "0123456789aBcDeFgHiJkLmNoPqRs",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLmNoPqRsT:Example",
			ExpectedNID: "0123456789aBcDeFgHiJkLmNoPqRsT",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLmNoPqRsTu:Example",
			ExpectedNID: "0123456789aBcDeFgHiJkLmNoPqRsTu",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLmNoPqRsTuV:Example",
			ExpectedNID: "0123456789aBcDeFgHiJkLmNoPqRsTuV",
		},



		{
			URN:     "urn:0-2:Example",
			ExpectedNID: "0-2",
		},
		{
			URN:     "urn:0--3:Example",
			ExpectedNID: "0--3",
		},
		{
			URN:     "urn:0---4:Example",
			ExpectedNID: "0---4",
		},
		{
			URN:     "urn:0----5:Example",
			ExpectedNID: "0----5",
		},
		{
			URN:     "urn:0-----6:Example",
			ExpectedNID: "0-----6",
		},
		{
			URN:     "urn:0------7:Example",
			ExpectedNID: "0------7",
		},
		{
			URN:     "urn:0-------8:Example",
			ExpectedNID: "0-------8",
		},
		{
			URN:     "urn:0--------9:Example",
			ExpectedNID: "0--------9",
		},
		{
			URN:     "urn:0---------a:Example",
			ExpectedNID: "0---------a",
		},
		{
			URN:     "urn:0----------B:Example",
			ExpectedNID: "0----------B",
		},
		{
			URN:     "urn:0-----------c:Example",
			ExpectedNID: "0-----------c",
		},
		{
			URN:     "urn:0------------D:Example",
			ExpectedNID: "0------------D",
		},
		{
			URN:     "urn:0-------------e:Example",
			ExpectedNID: "0-------------e",
		},
		{
			URN:     "urn:0--------------F:Example",
			ExpectedNID: "0--------------F",
		},
		{
			URN:     "urn:0---------------g:Example",
			ExpectedNID: "0---------------g",
		},
		{
			URN:     "urn:0----------------H:Example",
			ExpectedNID: "0----------------H",
		},
		{
			URN:     "urn:0-----------------i:Example",
			ExpectedNID: "0-----------------i",
		},
		{
			URN:     "urn:0------------------J:Example",
			ExpectedNID: "0------------------J",
		},
		{
			URN:     "urn:0-------------------k:Example",
			ExpectedNID: "0-------------------k",
		},
		{
			URN:     "urn:0--------------------L:Example",
			ExpectedNID: "0--------------------L",
		},
		{
			URN:     "urn:0---------------------m:Example",
			ExpectedNID: "0---------------------m",
		},
		{
			URN:     "urn:0----------------------N:Example",
			ExpectedNID: "0----------------------N",
		},
		{
			URN:     "urn:0-----------------------o:Example",
			ExpectedNID: "0-----------------------o",
		},
		{
			URN:     "urn:0------------------------P:Example",
			ExpectedNID: "0------------------------P",
		},
		{
			URN:     "urn:0-------------------------q:Example",
			ExpectedNID: "0-------------------------q",
		},
		{
			URN:     "urn:0--------------------------R:Example",
			ExpectedNID: "0--------------------------R",
		},
		{
			URN:     "urn:0---------------------------s:Example",
			ExpectedNID: "0---------------------------s",
		},
		{
			URN:     "urn:0----------------------------T:Example",
			ExpectedNID: "0----------------------------T",
		},
		{
			URN:     "urn:0-----------------------------u:Example",
			ExpectedNID: "0-----------------------------u",
		},
		{
			URN:     "urn:0------------------------------V:Example",
			ExpectedNID: "0------------------------------V",
		},
	}

	for testNumber, test := range tests {

		const prefix string = rfc8141.Scheme+":" // "urn:"
		var str string = test.URN[len(prefix):]

		actualN, actualFound := rfc8141.PeekPrefixNID(str)
		if !actualFound {
			t.Errorf("For test #%d, expected to 'find' an NID but actually did not.", testNumber)
			t.Logf("URN: %q", test.URN)
			continue
		}

		{
			expected := len(test.ExpectedNID)
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
			expected := test.ExpectedNID
			actual := str[:actualN]

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

func TestPeekPrefixNID_fail(t *testing.T) {

	tests := []struct{
		URN string
	}{
		{
			URN: "urn:-:something",
		},



		{
			URN: "urn:+:something",
		},



		{
			URN: "urn:0123456789aBcDeFgHiJkLmNoPqRsTuVw:Example",
		},
		{
			URN: "urn:0123456789aBcDeFgHiJkLmNoPqRsTuVwX:Example",
		},
		{
			URN: "urn:0123456789aBcDeFgHiJkLmNoPqRsTuVwXy:Example",
		},
		{
			URN: "urn:0123456789aBcDeFgHiJkLmNoPqRsTuVwXyZ:Example",
		},



		{
			URN: "urn:0-------------------------------w:Example",
		},
		{
			URN: "urn:0--------------------------------X:Example",
		},
		{
			URN: "urn:0---------------------------------y:Example",
		},
		{
			URN: "urn:0----------------------------------Z:Example",
		},



		{
			URN:     "urn:-1:Example",
		},
		{
			URN:     "urn:-12:Example",
		},
		{
			URN:     "urn:-123:Example",
		},
		{
			URN:     "urn:-1234:Example",
		},
		{
			URN:     "urn:-12345:Example",
		},
		{
			URN:     "urn:-123456:Example",
		},
		{
			URN:     "urn:-1234567:Example",
		},
		{
			URN:     "urn:-12345678:Example",
		},
		{
			URN:     "urn:-123456789:Example",
		},
		{
			URN:     "urn:-123456789a:Example",
		},
		{
			URN:     "urn:-123456789aB:Example",
		},
		{
			URN:     "urn:-123456789aBc:Example",
		},
		{
			URN:     "urn:-123456789aBcD:Example",
		},
		{
			URN:     "urn:-123456789aBcDe:Example",
		},
		{
			URN:     "urn:-123456789aBcDeF:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFg:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgH:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHi:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJ:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJk:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkL:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLm:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLmN:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLmNo:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLmNoP:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLmNoPq:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLmNoPqR:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLmNoPqRs:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLmNoPqRsT:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLmNoPqRsTu:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLmNoPqRsTuV:Example",
		},



		{
			URN:     "urn:0-:Example",
		},
		{
			URN:     "urn:01-:Example",
		},
		{
			URN:     "urn:012-:Example",
		},
		{
			URN:     "urn:0123-:Example",
		},
		{
			URN:     "urn:01234-:Example",
		},
		{
			URN:     "urn:012345-:Example",
		},
		{
			URN:     "urn:0123456-:Example",
		},
		{
			URN:     "urn:01234567-:Example",
		},
		{
			URN:     "urn:012345678-:Example",
		},
		{
			URN:     "urn:0123456789-:Example",
		},
		{
			URN:     "urn:0123456789a-:Example",
		},
		{
			URN:     "urn:0123456789aB-:Example",
		},
		{
			URN:     "urn:0123456789aBc-:Example",
		},
		{
			URN:     "urn:0123456789aBcD-:Example",
		},
		{
			URN:     "urn:0123456789aBcDe-:Example",
		},
		{
			URN:     "urn:0123456789aBcDeF-:Example",
		},
		{
			URN:     "urn:0123456789aBcDeFg-:Example",
		},
		{
			URN:     "urn:0123456789aBcDeFgH-:Example",
		},
		{
			URN:     "urn:0123456789aBcDeFgHi-:Example",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJ-:Example",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJk-:Example",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkL-:Example",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLm-:Example",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLmN-:Example",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLmNo-:Example",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLmNoP-:Example",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLmNoPq-:Example",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLmNoPqR-:Example",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLmNoPqRs-:Example",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLmNoPqRsT-:Example",
		},
		{
			URN:     "urn:0123456789aBcDeFgHiJkLmNoPqRsTu-:Example",
		},



		{
			URN:     "urn:0-:Example",
		},
		{
			URN:     "urn:-1-:Example",
		},
		{
			URN:     "urn:-12-:Example",
		},
		{
			URN:     "urn:-123-:Example",
		},
		{
			URN:     "urn:-1234-:Example",
		},
		{
			URN:     "urn:-12345-:Example",
		},
		{
			URN:     "urn:-123456-:Example",
		},
		{
			URN:     "urn:-1234567-:Example",
		},
		{
			URN:     "urn:-12345678-:Example",
		},
		{
			URN:     "urn:-123456789-:Example",
		},
		{
			URN:     "urn:-123456789a-:Example",
		},
		{
			URN:     "urn:-123456789aB-:Example",
		},
		{
			URN:     "urn:-123456789aBc-:Example",
		},
		{
			URN:     "urn:-123456789aBcD-:Example",
		},
		{
			URN:     "urn:-123456789aBcDe-:Example",
		},
		{
			URN:     "urn:-123456789aBcDeF-:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFg-:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgH-:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHi-:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJ-:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJk-:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkL-:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLm-:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLmN-:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLmNo-:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLmNoP-:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLmNoPq-:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLmNoPqR-:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLmNoPqRs-:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLmNoPqRsT-:Example",
		},
		{
			URN:     "urn:-123456789aBcDeFgHiJkLmNoPqRsTu-:Example",
		},
	}

	for testNumber, test := range tests {

		const prefix string = rfc8141.Scheme+":" // "urn:"
		var str string = test.URN[len(prefix):]

		actualN, actualFound := rfc8141.PeekPrefixNID(str)
		if actualFound {
			t.Errorf("For test #%d, did not expect to 'find' an NID but actually did.", testNumber)
			t.Logf("URN:    %q", test.URN)
			t.Logf("ACTUAL-N: %d", actualN)
			t.Logf("ACTUAL-NID: %q", str[:actualN])
			continue
		}
	}
}
