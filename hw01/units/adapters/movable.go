package adapters //nolint:dupl

import (
	"github.com/andreyAKor/otus_arch_hw/hw01/actions"
	"github.com/andreyAKor/otus_arch_hw/hw01/geometry"
	"github.com/andreyAKor/otus_arch_hw/hw01/repository"
)

const (
	UnitPosition = "position"
	UnitVelocity = "velocity"
)

var _ actions.Movable = (*movable)(nil)

// Адаптер движения.
type movable struct {
	storage repository.Storager
}

func NewMovable(storage repository.Storager) (actions.Movable, error) {
	if storage == nil {
		return nil, ErrStorageNil
	}

	return &movable{storage}, nil
}

func (m *movable) SetPosition(v *geometry.Vector) {
	m.storage.Set(UnitPosition, v)
}

func (m *movable) GetPosition() *geometry.Vector {
	if v, ok := m.storage.Get(UnitPosition).(*geometry.Vector); ok {
		return v
	}

	return nil
}

func (m *movable) GetVelocity() *geometry.Vector {
	if v, ok := m.storage.Get(UnitVelocity).(*geometry.Vector); ok {
		return v
	}

	return nil
}
