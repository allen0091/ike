package message

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIdentificationResponderMarshal(t *testing.T) {
	testcases := []struct {
		description string
		id          IdentificationResponder
		expMarshal  []byte
	}{
		{
			description: "IdentificationResponder marshal",
			id: IdentificationResponder{
				IDType: ID_KEY_ID,
				IDData: []byte{
					0x55, 0x45,
				},
			},
			expMarshal: []byte{
				0xb, 0x0, 0x0, 0x0, 0x55, 0x45,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := tc.id.marshal()
			require.NoError(t, err)
			require.Equal(t, tc.expMarshal, result)
		})
	}
}

func TestIdentificationResponderUnmarshal(t *testing.T) {
	testcases := []struct {
		description string
		b           []byte
		expErr      bool
		expMarshal  IdentificationResponder
	}{
		{
			description: "No sufficient bytes to decode next identification",
			b: []byte{
				0x01, 0x02, 0x03,
			},
			expErr: true,
		},
		{
			description: "IdentificationResponder Unmarshal",
			b: []byte{
				0xb, 0x0, 0x0, 0x0, 0x55, 0x45,
			},
			expMarshal: IdentificationResponder{
				IDType: ID_KEY_ID,
				IDData: []byte{
					0x55, 0x45,
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			var id IdentificationResponder
			err := id.unmarshal(tc.b)
			if tc.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expMarshal, id)
			}
		})
	}
}
