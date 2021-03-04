package factory

type SorterFactory interface {
	Sort(items interface{}) interface{}
	Name() string
}
