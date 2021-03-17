package factory

import (
	"io"

	"github.com/andreyAKor/otus_arch_hw/hw07/iterator"
)

var _ OperationIface = (*fibonacciOperation)(nil)

// Конкретный продукт для Фибоначчи.
type fibonacciOperation struct {
	operation
	fn func() int
}

//nolint:exhaustivestruct
func newFibonacciOperation(limit int) OperationIface {
	fo := &fibonacciOperation{
		operation: operation{
			limit:   limit,
			counter: 0,
		},
	}
	fo.fn = fo.iterationFunc()
	fo.operation.iterator = iterator.New(fo.next)

	return fo
}

func (f *fibonacciOperation) next() (int, error) {
	if f.operation.counter == f.operation.limit {
		return 0, io.EOF
	}

	f.operation.counter++

	return f.fn(), nil
}

func (f fibonacciOperation) iterationFunc() func() int {
	x := 0
	y := 1

	return func() int {
		x, y = y, x+y

		return x
	}
}
