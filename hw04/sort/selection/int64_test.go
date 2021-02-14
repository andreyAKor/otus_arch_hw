package selection

import (
	"testing"

	"github.com/stretchr/testify/require"
)

//nolint:paralleltest
func TestInt64(t *testing.T) {
	t.Run("normal numbers", func(t *testing.T) {
		items := []int64{4, 1, 5, 2, 3}
		res := NewInt64().Sort(items)
		require.Equal(t, res, []int64{1, 2, 3, 4, 5})
	})
	t.Run("negative numbers", func(t *testing.T) {
		items := []int64{-4, 1, 5, -2, 3}
		res := NewInt64().Sort(items)
		require.Equal(t, res, []int64{-4, -2, 1, 3, 5})
	})
}
