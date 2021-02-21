package main

import (
	"errors"
	"os"
	"testing"

	"github.com/andreyAKor/otus_arch_hw/hw07/factory"
	"github.com/stretchr/testify/require"
)

//nolint:paralleltest
func TestWrite(t *testing.T) {
	t.Run("empty files", func(t *testing.T) {
		err := write("", nil)
		require.True(t, os.IsNotExist(errors.Unwrap(err)))
	})
	t.Run("permission file", func(t *testing.T) {
		fibonacciFactory, _ := factory.Factory("fibonacci")
		fibonacciOperation := fibonacciFactory.MakeOperation(10)

		err := write("/dev/psaux", fibonacciOperation)
		require.True(t, os.IsPermission(errors.Unwrap(err)))
	})
}
