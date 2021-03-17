package merge

const nameInt64 string = "merge" // Название сортировки

var _ Sorter = (*Int64)(nil)

type Int64 struct{}

func NewInt64() Sorter {
	return &Int64{}
}

//nolint:unconvert
func (i Int64) Sort(items interface{}) interface{} {
	n := len(items.([]int64))

	if n == 1 {
		return items
	}

	middle := int(n / 2)
	left := make([]int64, middle)
	right := make([]int64, n-middle)

	for i := 0; i < n; i++ {
		if i < middle {
			left[i] = items.([]int64)[i]
		} else {
			right[i-middle] = items.([]int64)[i]
		}
	}

	return i.merging(i.Sort(left).([]int64), i.Sort(right).([]int64))
}

func (Int64) Name() string {
	return nameInt64
}

func (Int64) merging(left, right []int64) interface{} {
	result := make([]int64, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}
