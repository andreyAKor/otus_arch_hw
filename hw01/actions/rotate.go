package actions

var _ Doer = (*Rotate)(nil)

// Структура поворота
type Rotate struct {
	r Rotable
}

func NewRotate(r Rotable) *Rotate {
	return &Rotate{r}
}

func (r *Rotate) Do() {
	r.r.SetDirection(r.r.GetDirection().Add(r.r.GetAngularVelocity()))
}
