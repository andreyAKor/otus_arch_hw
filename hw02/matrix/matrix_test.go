package matrix

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

// nolint:funlen,paralleltest
func TestMul(t *testing.T) {
	// Пост-проверка на утечку горутин после прогонов тестов
	defer goleak.VerifyNone(t)

	a := Matrix{
		{1, 2},
		{3, 4},
	}
	b := Matrix{
		{5, 6},
		{7, 8},
	}
	finalC := Matrix{
		{5, 12},
		{21, 32},
	}

	// Нормальные кейсы работы перемножений
	t.Run("normal multiples", func(t *testing.T) {
		// Перемножение матриц в 1 поток
		t.Run("in one threads", func(t *testing.T) {
			c, err := Mul(1, &a, &b)
			require.NotNil(t, c)
			require.NoError(t, err)
			require.Equal(t, &finalC, c)
		})

		// Перемножение матриц в 2 потока
		t.Run("in two threads", func(t *testing.T) {
			c, err := Mul(2, &a, &b)
			require.NotNil(t, c)
			require.NoError(t, err)
			require.Equal(t, &finalC, c)
		})

		// Перемножение матриц в 2 000 потоков
		t.Run("in 2 000 threads", func(t *testing.T) {
			c, err := Mul(2_000, &a, &b)
			require.NotNil(t, c)
			require.NoError(t, err)
			require.Equal(t, &finalC, c)
		})
	})

	// Ошибочные кейсы переменожений матриц
	t.Run("errorneous multiples", func(t *testing.T) {
		// Перемножение матриц в 0 потоков
		t.Run("threads less that one", func(t *testing.T) {
			c, err := Mul(0, &a, &b)
			require.Nil(t, c)
			require.EqualError(t, ErrThreadsLessOne, err.Error())
		})

		// Матрица A не проинициализирована
		t.Run("matrix A is nil", func(t *testing.T) {
			c, err := Mul(1, nil, &b)
			require.Nil(t, c)
			require.EqualError(t, ErrMatrixAIsNil, err.Error())
		})

		// Матрица B не проинициализирована
		t.Run("matrix B is nil", func(t *testing.T) {
			c, err := Mul(1, &a, nil)
			require.Nil(t, c)
			require.EqualError(t, ErrMatrixBIsNil, err.Error())
		})

		// Матрица A пустая
		t.Run("matrix A is empty", func(t *testing.T) {
			c, err := Mul(1, &Matrix{}, &b)
			require.Nil(t, c)
			require.EqualError(t, ErrMatrixAIsEmpty, err.Error())
		})

		// Матрица B пустая
		t.Run("matrix B is empty", func(t *testing.T) {
			c, err := Mul(1, &a, &Matrix{})
			require.Nil(t, c)
			require.EqualError(t, ErrMatrixBIsEmpty, err.Error())
		})

		// Различия матриц по размерам
		t.Run("matrix A and B vary in size", func(t *testing.T) {
			// По количеству строк
			t.Run("by rows", func(t *testing.T) {
				c, err := Mul(
					1,
					&Matrix{
						{1, 2},
					},
					&b,
				)
				require.Nil(t, c)
				require.EqualError(t, ErrMatrixABVary, err.Error())
			})

			// По количеству значений в строке
			t.Run("by items in some row", func(t *testing.T) {
				c, err := Mul(
					1,
					&Matrix{
						{1, 2},
						{3, 4, 5},
					},
					&b,
				)
				require.Nil(t, c)
				require.EqualError(t, ErrMatrixABVary, err.Error())
			})
		})
	})
}

// nolint:paralleltest
func TestWorker(t *testing.T) {
	// Пост-проверка на утечку горутин после прогонов тестов
	defer goleak.VerifyNone(t)

	// Проверка на нормальное завершение воркера
	t.Run("worker normal shutdown", func(t *testing.T) {
		wg := &sync.WaitGroup{}
		taskCh := make(chan task, 1)

		wg.Add(1)
		go worker(wg, taskCh, nil, nil, nil)

		close(taskCh)
		wg.Wait()
	})

	// Проверка на адекватную калькуляцию воркером
	t.Run("worker normal calc", func(t *testing.T) {
		a := Matrix{
			{1, 2},
			{3, 4},
		}
		b := Matrix{
			{5, 6},
			{7, 8},
		}
		c := Matrix{
			{0, 0},
			{0, 0},
		}
		finalC := Matrix{
			{0, 0},
			{0, 32},
		}

		wg := &sync.WaitGroup{}
		taskCh := make(chan task, 1)

		wg.Add(1)
		go worker(wg, taskCh, &a, &b, &c)

		taskCh <- task{len(finalC) - 1, len(finalC[len(finalC)-1]) - 1}

		close(taskCh)
		wg.Wait()

		require.Equal(t, finalC, c)
	})
}
