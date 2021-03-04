package merge

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

func (f Factory) Factory() Sorter {
	return NewInt64()
}
