package bls

import (
	"testing"

	"github.com/Kevionte/prysm_beacon/v1crypto/bls/common"
	"github.com/Kevionte/prysm_beacon/v1testing/require"
)

func TestDisallowZeroSecretKeys(t *testing.T) {
	t.Run("blst", func(t *testing.T) {
		// Blst does a zero check on the key during deserialization.
		_, err := SecretKeyFromBytes(common.ZeroSecretKey[:])
		require.Equal(t, common.ErrSecretUnmarshal, err)
	})
}

func TestDisallowZeroPublicKeys(t *testing.T) {
	t.Run("blst", func(t *testing.T) {
		_, err := PublicKeyFromBytes(common.InfinitePublicKey[:])
		require.Equal(t, common.ErrInfinitePubKey, err)
	})
}

func TestDisallowZeroPublicKeys_AggregatePubkeys(t *testing.T) {
	t.Run("blst", func(t *testing.T) {
		_, err := AggregatePublicKeys([][]byte{common.InfinitePublicKey[:], common.InfinitePublicKey[:]})
		require.Equal(t, common.ErrInfinitePubKey, err)
	})
}
