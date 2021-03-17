package selection

type Sorter interface {
	Sort(items interface{}) interface{}
	Name() string
}
