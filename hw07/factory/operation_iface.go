package factory

import (
	"github.com/andreyAKor/otus_arch_hw/hw07/iterator"
)

// Абстрактный продукт операции.
type OperationIface interface {
	Limit() int
	Iterator() iterator.Iterator
}

type operation struct {
	iterator iterator.Iterator
	limit    int
	counter  int
}

func (o *operation) Limit() int {
	return o.limit
}

func (o *operation) Iterator() iterator.Iterator {
	return o.iterator
}
