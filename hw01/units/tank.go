package units

import (
	"github.com/andreyAKor/otus_arch_hw/hw01/repository"
)

// Танк
type Tank struct {
	storage repository.Storager
}

func NewTank(s repository.Storager) *Tank {
	return &Tank{s}
}
