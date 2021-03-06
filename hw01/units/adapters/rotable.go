package adapters //nolint:dupl

import (
	"github.com/andreyAKor/otus_arch_hw/hw01/actions"
	"github.com/andreyAKor/otus_arch_hw/hw01/geometry"
	"github.com/andreyAKor/otus_arch_hw/hw01/repository"
)

const (
	UnitDirection       = "direction"
	UnitAngularVelocity = "angularVelocity"
)

var _ actions.Rotable = (*rotable)(nil)

// Адаптер поворота.
type rotable struct {
	storage repository.Storager
}

func NewRotable(storage repository.Storager) (actions.Rotable, error) {
	if storage == nil {
		return nil, ErrStorageNil
	}

	return &rotable{storage}, nil
}

func (r *rotable) SetDirection(v *geometry.Vector) {
	r.storage.Set(UnitDirection, v)
}

func (r *rotable) GetDirection() *geometry.Vector {
	if v, ok := r.storage.Get(UnitDirection).(*geometry.Vector); ok {
		return v
	}

	return nil
}

func (r *rotable) GetAngularVelocity() *geometry.Vector {
	if v, ok := r.storage.Get(UnitAngularVelocity).(*geometry.Vector); ok {
		return v
	}

	return nil
}
