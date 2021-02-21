package factory

import (
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

//nolint:paralleltest
func TestBinetOperation(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		o := newBinetOperation(10)

		items := []int{}

		i := o.Iterator()
		for {
			v, err := i.Next()
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				t.Fatalf("next iterate fail: %s", err)
			}

			items = append(items, v)
		}

		require.Equal(t, []int{55, 34, 21, 13, 8, 5, 3, 2, 1, 1}, items)
	})
}
