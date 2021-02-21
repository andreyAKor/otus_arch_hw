package factory

import (
	"io"

	"github.com/andreyAKor/otus_arch_hw/hw07/iterator"
)

var _ OperationIface = (*fibonacciOperation)(nil)

// Конкретный продукт для Фибоначчи.
type fibonacciOperation struct {
	operation
	f func() int
}

func newFibonacciOperation(limit int) OperationIface {
	f := &fibonacciOperation{
		operation: operation{
			limit: limit,
		},
	}
	f.f = f.iterationFunc()
	f.operation.iterator = iterator.New(f.next)

	return f
}

func (f *fibonacciOperation) next() (int, error) {
	if f.operation.counter == f.operation.limit {
		return 0, io.EOF
	}

	f.operation.counter++

	return f.f(), nil
}

func (f fibonacciOperation) iterationFunc() func() int {
	x := 0
	y := 1

	return func() int {
		x, y = y, x+y

		return x
	}
}
