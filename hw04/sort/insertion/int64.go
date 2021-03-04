package insertion

const nameInt64 string = "insertion" // Название сортировки

var _ Sorter = (*Int64)(nil)

type Int64 struct{}

func NewInt64() Sorter {
	return &Int64{}
}

func (Int64) Sort(items interface{}) interface{} {
	n := len(items.([]int64))

	for i := 1; i < n; i++ {
		j := i

		for j > 0 && items.([]int64)[j] < items.([]int64)[j-1] {
			items.([]int64)[j], items.([]int64)[j-1] = items.([]int64)[j-1], items.([]int64)[j]
			j--
		}
	}

	return items
}

func (Int64) Name() string {
	return nameInt64
}
