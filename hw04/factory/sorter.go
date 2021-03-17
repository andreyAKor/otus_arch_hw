package factory

import (
	"github.com/andreyAKor/otus_arch_hw/hw04/sort/insertion"
	"github.com/andreyAKor/otus_arch_hw/hw04/sort/merge"
	"github.com/andreyAKor/otus_arch_hw/hw04/sort/selection"
)

type Sorter struct{}

func NewSorter() *Sorter {
	return &Sorter{}
}

func (f Sorter) Factory(_type int) (s SorterFactory) {
	// В зависимости от типа, создаем сортировщика
	switch _type {
	case 1:
		s = selection.NewFactory().Factory()
	case 2:
		s = insertion.NewFactory().Factory()
	case 3:
		s = merge.NewFactory().Factory()
	}

	return
}
