package copy

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDeepCopy(t *testing.T) {
	type Demo struct {
		mp  map[string]int
		arr []int
		i   int
		s   string
		f   float64
		t   time.Time
	}

	src := &Demo{
		mp: map[string]int{
			"test": 1,
		},
		arr: []int{1, 2, 3},
		i:   1,
		s:   "str",
		f:   1.1,
		t:   time.Date(0, 0, 0, 0, 1, 1, 1, time.Local),
	}

	deepCopy, err := DeepCopy(src)
	require.Nil(t, err)
	require.EqualValues(t, deepCopy, src)
}
