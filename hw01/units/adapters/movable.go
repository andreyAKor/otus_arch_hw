package adapters

import (
	"github.com/andreyAKor/otus_arch_hw/hw01/actions"
	"github.com/andreyAKor/otus_arch_hw/hw01/geometry"
	"github.com/andreyAKor/otus_arch_hw/hw01/repository"
)

const (
	UnitPosition = "position"
	UnitVelocity = "velocity"
)

var _ actions.Movable = (*Movable)(nil)

// Адаптер движения
type Movable struct {
	storage repository.Storager
}

func NewMovable(s repository.Storager) *Movable {
	return &Movable{s}
}

func (m *Movable) SetPosition(v *geometry.Vector) {
	m.storage.Set(UnitPosition, v)
}

func (m *Movable) GetPosition() *geometry.Vector {
	if v, ok := m.storage.Get(UnitPosition).(*geometry.Vector); ok {
		return v
	}
	return nil
}

func (m *Movable) GetVelocity() *geometry.Vector {
	if v, ok := m.storage.Get(UnitVelocity).(*geometry.Vector); ok {
		return v
	}
	return nil
}
