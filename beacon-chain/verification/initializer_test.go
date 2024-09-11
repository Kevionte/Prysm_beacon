package verification

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/Kevionte/prysm_beacon/v1beacon-chain/startup"
	"github.com/Kevionte/prysm_beacon/v1encoding/bytesutil"
	"github.com/Kevionte/prysm_beacon/v1testing/require"
)

func TestInitializerWaiter(t *testing.T) {
	ctx := context.Background()
	vr := bytesutil.ToBytes32([]byte{0, 1, 1, 2, 3, 5})
	gen := time.Now()
	c := startup.NewClock(gen, vr)
	cs := startup.NewClockSynchronizer()
	require.NoError(t, cs.SetClock(c))

	w := NewInitializerWaiter(cs, &mockForkchoicer{}, &mockStateByRooter{})
	ini, err := w.WaitForInitializer(ctx)
	require.NoError(t, err)
	csc, ok := ini.shared.sc.(*sigCache)
	require.Equal(t, true, ok)
	require.Equal(t, true, bytes.Equal(vr[:], csc.valRoot))
}
