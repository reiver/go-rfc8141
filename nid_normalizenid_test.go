package rfc8141_test

import (
	"testing"

	"github.com/reiver/go-rfc8141"
)

func TestNormalizeNID(t *testing.T) {
	tests := []struct{
		NID string
		ExpectedNID string
	}{
		{
			NID:         "ietf",
			ExpectedNID: "ietf",
		},
		{
			NID:         "ietF",
			ExpectedNID: "ietf",
		},
		{
			NID:         "ieTf",
			ExpectedNID: "ietf",
		},
		{
			NID:         "ieTF",
			ExpectedNID: "ietf",
		},
		{
			NID:         "iEtf",
			ExpectedNID: "ietf",
		},
		{
			NID:         "iEtF",
			ExpectedNID: "ietf",
		},
		{
			NID:         "iETf",
			ExpectedNID: "ietf",
		},
		{
			NID:         "iETF",
			ExpectedNID: "ietf",
		},
		{
			NID:         "Ietf",
			ExpectedNID: "ietf",
		},
		{
			NID:         "IetF",
			ExpectedNID: "ietf",
		},
		{
			NID:         "IeTf",
			ExpectedNID: "ietf",
		},
		{
			NID:         "IeTF",
			ExpectedNID: "ietf",
		},
		{
			NID:         "IEtf",
			ExpectedNID: "ietf",
		},
		{
			NID:         "IEtF",
			ExpectedNID: "ietf",
		},
		{
			NID:         "IETf",
			ExpectedNID: "ietf",
		},
		{
			NID:         "IETF",
			ExpectedNID: "ietf",
		},



		{
			NID:         "sha1",
			ExpectedNID: "sha1",
		},
		{
			NID:         "shA1",
			ExpectedNID: "sha1",
		},
		{
			NID:         "sHa1",
			ExpectedNID: "sha1",
		},
		{
			NID:         "sHA1",
			ExpectedNID: "sha1",
		},
		{
			NID:         "Sha1",
			ExpectedNID: "sha1",
		},
		{
			NID:         "ShA1",
			ExpectedNID: "sha1",
		},
		{
			NID:         "SHa1",
			ExpectedNID: "sha1",
		},
		{
			NID:         "SHA1",
			ExpectedNID: "sha1",
		},



		{
			NID:         "uuid",
			ExpectedNID: "uuid",
		},
		{
			NID:         "uuiD",
			ExpectedNID: "uuid",
		},
		{
			NID:         "uuId",
			ExpectedNID: "uuid",
		},
		{
			NID:         "uuID",
			ExpectedNID: "uuid",
		},
		{
			NID:         "uUid",
			ExpectedNID: "uuid",
		},
		{
			NID:         "uUiD",
			ExpectedNID: "uuid",
		},
		{
			NID:         "uUId",
			ExpectedNID: "uuid",
		},
		{
			NID:         "uUID",
			ExpectedNID: "uuid",
		},
		{
			NID:         "Uuid",
			ExpectedNID: "uuid",
		},
		{
			NID:         "UuiD",
			ExpectedNID: "uuid",
		},
		{
			NID:         "UuId",
			ExpectedNID: "uuid",
		},
		{
			NID:         "UuID",
			ExpectedNID: "uuid",
		},
		{
			NID:         "UUid",
			ExpectedNID: "uuid",
		},
		{
			NID:         "UUiD",
			ExpectedNID: "uuid",
		},
		{
			NID:         "UUId",
			ExpectedNID: "uuid",
		},
		{
			NID:         "UUID",
			ExpectedNID: "uuid",
		},



		{
			NID:         "xmpp",
			ExpectedNID: "xmpp",
		},
		{
			NID:         "xmpP",
			ExpectedNID: "xmpp",
		},
		{
			NID:         "xmPp",
			ExpectedNID: "xmpp",
		},
		{
			NID:         "xmPP",
			ExpectedNID: "xmpp",
		},
		{
			NID:         "xMpp",
			ExpectedNID: "xmpp",
		},
		{
			NID:         "xMpP",
			ExpectedNID: "xmpp",
		},
		{
			NID:         "xMPp",
			ExpectedNID: "xmpp",
		},
		{
			NID:         "xMPP",
			ExpectedNID: "xmpp",
		},
		{
			NID:         "Xmpp",
			ExpectedNID: "xmpp",
		},
		{
			NID:         "XmpP",
			ExpectedNID: "xmpp",
		},
		{
			NID:         "XmPp",
			ExpectedNID: "xmpp",
		},
		{
			NID:         "XmPP",
			ExpectedNID: "xmpp",
		},
		{
			NID:         "XMpp",
			ExpectedNID: "xmpp",
		},
		{
			NID:         "XMpP",
			ExpectedNID: "xmpp",
		},
		{
			NID:         "XMPp",
			ExpectedNID: "xmpp",
		},
		{
			NID:         "XMPP",
			ExpectedNID: "xmpp",
		},
	}

	for testNumber, test := range tests {

		actual := rfc8141.NormalizeNID(test.NID)

		expected := test.ExpectedNID

		if expected != actual {
			t.Errorf("For test #%d, the actual normalized NID is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("NID:      %q", test.NID)
			continue
		}
	}
}
