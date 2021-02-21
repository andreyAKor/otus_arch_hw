package factory

import (
	"io"
	"math"

	"github.com/andreyAKor/otus_arch_hw/hw07/iterator"
)

var _ OperationIface = (*binetOperation)(nil)

// Конкретный продукт для Бине.
type binetOperation struct {
	operation
}

func newBinetOperation(limit int) OperationIface {
	b := &binetOperation{
		operation: operation{
			limit:   limit,
			counter: 0,
		},
	}
	b.operation.iterator = iterator.New(b.next)

	return b
}

func (b *binetOperation) next() (int, error) {
	if b.operation.counter == b.operation.limit {
		return 0, io.EOF
	}

	b.operation.counter++

	return int(b.binet(float64(b.operation.limit - b.operation.counter + 1))), nil
}

func (b *binetOperation) binet(numIndex float64) float64 {
	Phi := (1 + math.Sqrt(5)) / 2
	phi := (1 - math.Sqrt(5)) / 2

	return math.Ceil((math.Pow(Phi, numIndex) - math.Pow(phi, numIndex)) / math.Sqrt(5))
}
