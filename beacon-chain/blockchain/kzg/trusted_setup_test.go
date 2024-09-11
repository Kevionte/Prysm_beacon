package kzg

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1testing/require"
)

func TestStart(t *testing.T) {
	require.NoError(t, Start())
	require.NotNil(t, kzgContext)
}
