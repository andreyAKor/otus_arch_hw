package factory

var _ NumbersFactoryIface = (*binet)(nil)

// Конкретная фабрика для Бине.
type binet struct{}

func (b *binet) MakeOperation(limit int) OperationIface {
	return newBinetOperation(limit)
}
