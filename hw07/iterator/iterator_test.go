package iterator

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:paralleltest
func TestSimpleIterator(t *testing.T) {
	items := generateItems(0, 10)
	computeNext := next(items)

	iterator := New(computeNext)
	total := 0
	for range items {
		next, err := iterator.Next()
		assert.Nil(t, err)
		assert.NotNil(t, next)

		i := next
		total += i
	}
	assert.Equal(t, 45, total)
}

//nolint:paralleltest
func TestCloseHandler(t *testing.T) {
	items := generateItems(0, 10)
	computeNext, idx := nextAndIndex(items)
	iterator := New(computeNext)

	for range items {
		next, err := iterator.Next()
		assert.NotNil(t, next)
		assert.Nil(t, err)
	}

	assert.Equal(t, 10, *idx)
}

//nolint:unused
type Items []int

func (i Items) Iterator() Iterator {
	return New(next(i))
}

func next(items []int) ComputeNext {
	index := 0

	return func() (int, error) {
		if index >= len(items) {
			return 0, io.EOF
		}

		n := items[index]
		index++

		return n, nil
	}
}

func nextAndIndex(items []int) (ComputeNext, *int) {
	index := 0

	return func() (int, error) {
		if index >= len(items) {
			return 0, io.EOF
		}

		n := items[index]
		index++

		return n, nil
	}, &index
}

func generateItems(from int, to int) []int {
	var items []int
	for from < to {
		items = append(items, from)
		from++
	}

	return items
}
