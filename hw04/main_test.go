package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

//nolint:paralleltest
func TestRead(t *testing.T) {
	t.Run("empty files", func(t *testing.T) {
		_, err := read("")
		require.True(t, os.IsNotExist(errors.Unwrap(err)))

		_, err = read("/etc/hosts")
		require.False(t, os.IsNotExist(err))
	})
	t.Run("unlimit file", func(t *testing.T) {
		_, err := read("/dev/urandom")
		require.EqualError(t, err, ErrUnsupportedFile.Error())
	})
	t.Run("permission file", func(t *testing.T) {
		_, err := read("/dev/psaux")
		require.True(t, os.IsPermission(errors.Unwrap(err)))
	})
}

//nolint:paralleltest
func TestWrite(t *testing.T) {
	t.Run("empty files", func(t *testing.T) {
		err := write("", "", nil)
		require.True(t, os.IsNotExist(errors.Unwrap(err)))
	})
	t.Run("permission file", func(t *testing.T) {
		err := write("/dev/psaux", "", nil)
		require.True(t, os.IsPermission(errors.Unwrap(err)))
	})
}
