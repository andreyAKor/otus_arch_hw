package iterator

type ComputeNext func() (next int, err error)

type Iterator interface {
	Next() (int, error)
}

var _ Iterator = (*iterator)(nil)

// Итератор.
type iterator struct {
	computeNext ComputeNext
}

func New(computeNext ComputeNext) Iterator {
	return &iterator{computeNext}
}

// Возвращает следующий элемент итерации.
func (i *iterator) Next() (int, error) {
	next, err := i.computeNext()
	if err != nil {
		return 0, err
	}

	return next, nil
}
