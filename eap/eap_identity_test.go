package eap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	validEapIdentity = EapIdentity{
		IdentityData: []byte{
			0x7d, 0x09, 0x18, 0x42, 0x60, 0x9c, 0x9e, 0x20,
			0x56, 0x9f, 0xc0, 0x39, 0xda, 0x3f, 0x22, 0x2a,
			0xb8, 0x56, 0x81, 0x8a,
		},
	}

	validEapIdentityByte = []byte{
		0x01, 0x7d, 0x09, 0x18, 0x42, 0x60, 0x9c, 0x9e,
		0x20, 0x56, 0x9f, 0xc0, 0x39, 0xda, 0x3f, 0x22,
		0x2a, 0xb8, 0x56, 0x81, 0x8a,
	}
)

func TestEapIdentityMarshal(t *testing.T) {
	testcases := []struct {
		description string
		eap         EapIdentity
		expMarshal  []byte
		expErr      bool
	}{
		{
			description: "EAP identity is empty",
			eap: EapIdentity{
				IdentityData: nil,
			},
			expErr: true,
		},
		{
			description: "EapIdentity Marshal",
			eap:         validEapIdentity,
			expMarshal:  validEapIdentityByte,
			expErr:      false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := tc.eap.Marshal()
			if tc.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expMarshal, result)
			}
		})
	}
}

func TestEapIdentityUnmarshal(t *testing.T) {
	testcases := []struct {
		description string
		b           []byte
		expMarshal  EapIdentity
	}{
		{
			description: "EapIdentity Unmarshal",
			b:           validEapIdentityByte,
			expMarshal:  validEapIdentity,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			var eap EapIdentity
			err := eap.Unmarshal(tc.b)
			require.NoError(t, err)
			require.Equal(t, tc.expMarshal, eap)
		})
	}
}

func TestEapIdentitySetIdentityData(t *testing.T) {
	id := EapIdentity{}
	data := []byte{0x11, 0x22, 0x33}
	id.SetIdentityData(data)
	require.Equal(t, data, id.IdentityData)
	// Ensure it is a copy, not a reference
	data[0] = 0x99
	require.NotEqual(t, data, id.IdentityData)
}

func TestEapIdentitySetIdentityDataString(t *testing.T) {
	id := EapIdentity{}
	str := "user01"
	id.SetIdentityDataString(str)
	require.Equal(t, []byte(str), id.IdentityData)
}
