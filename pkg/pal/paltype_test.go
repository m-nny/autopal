package pal

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTypeJsonSuccess(t *testing.T) {
	for _, wantType := range TypeAll {
		t.Run(string(wantType), func(t *testing.T) {
			require := require.New(t)

			gotBuf, err := json.Marshal(wantType)
			require.NoError(err)

			wantBuf := []byte(fmt.Sprintf(`"%s"`, wantType))
			require.Equal(wantBuf, gotBuf)

			var gotType Type
			err = json.Unmarshal(gotBuf, &gotType)
			require.NoError(err)
			require.Equal(wantType, gotType)
		})
	}
}

func TestTypeUnmarshallJsonErr(t *testing.T) {
	testCases := [][]byte{
		[]byte(`invalid_json`),
		[]byte(`"not_type"`),
	}
	for _, testBuf := range testCases {
		t.Run(string(testBuf), func(t *testing.T) {
			require := require.New(t)

			var gotType Type
			err := json.Unmarshal(testBuf, &gotType)
			require.Error(err)
		})
	}
}
