package factory

var _ NumbersFactoryIface = (*fibonacci)(nil)

// Конкретная фабрика для Фибоначчи.
type fibonacci struct{}

func (a *fibonacci) MakeOperation(limit int) OperationIface {
	return newFibonacciOperation(limit)
}
