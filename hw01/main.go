package main

import (
	"fmt"

	"github.com/andreyAKor/otus_arch_hw/hw01/repository"
	"github.com/andreyAKor/otus_arch_hw/hw01/units"
	//	"github.com/andreyAKor/otus_arch_hw/hw01/units/adapters"
)

func main() {
	fmt.Println("Start")

	var storage repository.Storager

	_ = units.NewTank(storage)

	fmt.Println("End")
}
