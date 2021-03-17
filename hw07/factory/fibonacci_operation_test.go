package factory

import (
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

//nolint:paralleltest
func TestFibonacciOperation(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		o := newFibonacciOperation(10)

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

		require.Equal(t, []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}, items)
	})
}
