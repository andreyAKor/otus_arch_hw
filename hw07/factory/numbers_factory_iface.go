package factory

import (
	"fmt"
)

// Интерфейс абстрактной фабрики.
type NumbersFactoryIface interface {
	MakeOperation(limit int) OperationIface
}

func Factory(operation string) (NumbersFactoryIface, error) {
	if operation == "fibonacci" {
		return &fibonacci{}, nil
	}

	if operation == "binet" {
		return &binet{}, nil
	}

	return nil, fmt.Errorf(`wrong operation type passed by "%s"`, operation)
}
