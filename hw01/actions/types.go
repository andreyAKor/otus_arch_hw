package actions

import (
	"github.com/andreyAKor/otus_arch_hw/hw01/geometry"
)

// Деятель
type Doer interface {
	Do()
}

// Движения
type Movable interface {
	// Позиция
	SetPosition(v *geometry.Vector)
	GetPosition() *geometry.Vector

	// Скорость
	GetVelocity() *geometry.Vector
}

// Поворота
type Rotable interface {
	// Направление
	SetDirection(v *geometry.Vector)
	GetDirection() *geometry.Vector

	// Угловая скорость
	GetAngularVelocity() *geometry.Vector
}
