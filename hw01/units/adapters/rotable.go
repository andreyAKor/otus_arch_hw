package adapters

import (
	"github.com/andreyAKor/otus_arch_hw/hw01/actions"
	"github.com/andreyAKor/otus_arch_hw/hw01/geometry"
	"github.com/andreyAKor/otus_arch_hw/hw01/repository"
)

const (
	UnitDirection       = "direction"
	UnitAngularVelocity = "angularVelocity"
)

var _ actions.Rotable = (*Rotable)(nil)

// Адаптер поворота
type Rotable struct {
	storage repository.Storager
}

func NewRotable(s repository.Storager) *Rotable {
	return &Rotable{s}
}

func (r *Rotable) SetDirection(v *geometry.Vector) {
	r.storage.Set(UnitDirection, v)
}

func (r *Rotable) GetDirection() *geometry.Vector {
	if v, ok := r.storage.Get(UnitDirection).(*geometry.Vector); ok {
		return v
	}
	return nil
}

func (r *Rotable) GetAngularVelocity() *geometry.Vector {
	if v, ok := r.storage.Get(UnitAngularVelocity).(*geometry.Vector); ok {
		return v
	}
	return nil
}
