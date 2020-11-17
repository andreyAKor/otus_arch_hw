package geometry

// Вектор
type Vector struct {
	Start Point
	End   Point
}

func (v *Vector) Add(v2 *Vector) *Vector {
	return v // TODO: реализовать сложение векторов
}
