package ewsType

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTimeSpan(t *testing.T) {
	tm := NewTimeSpan(26*time.Hour + 13*time.Minute + 8*time.Second + 120*time.Millisecond + 55*time.Microsecond)
	require.Equal(t, int64((26*60+13)*60+8), tm.TotalSeconds())
	require.Equal(t, 120, tm.Milliseconds())
	require.Equal(t, 8, tm.Seconds())
	require.Equal(t, 13, tm.Minutes())
	require.Equal(t, 2, tm.Hours())
	require.Equal(t, 1, tm.Days())
}
