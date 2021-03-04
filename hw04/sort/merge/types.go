package merge

type Sorter interface {
	Sort(items interface{}) interface{}
	Name() string
}
