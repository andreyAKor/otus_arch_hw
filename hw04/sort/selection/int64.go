package selection

const nameInt64 string = "selection" // Название сортировки

var _ Sorter = (*Int64)(nil)

type Int64 struct{}

func NewInt64() Sorter {
	return &Int64{}
}

func (Int64) Sort(items interface{}) interface{} {
	n := len(items.([]int64))

	for i := 0; i < n; i++ {
		minIdx := i

		for j := i; j < n; j++ {
			if items.([]int64)[j] < items.([]int64)[minIdx] {
				minIdx = j
			}
		}

		items.([]int64)[i], items.([]int64)[minIdx] = items.([]int64)[minIdx], items.([]int64)[i]
	}

	return items
}

func (Int64) Name() string {
	return nameInt64
}
