package matrix

import (
	"errors"
	"sync"
)

var (
	ErrThreadsLessOne = errors.New("threads less that one")
	ErrMatrixAIsNil   = errors.New("matrix A is nil")
	ErrMatrixBIsNil   = errors.New("matrix B is nil")
	ErrMatrixAIsEmpty = errors.New("matrix A is empty")
	ErrMatrixBIsEmpty = errors.New("matrix B is empty")
	ErrMatrixABVary   = errors.New("matrix A and B vary in size")
)

// Строка матрицы.
type Row []int

// Матрица.
type Matrix []Row

// Структура задачи для воркера.
type task struct {
	i, j int
}

// Перемножение матриц A и B в threads потоках/воркерах.
func Mul(threads int, a, b *Matrix) (*Matrix, error) {
	if threads < 1 {
		return nil, ErrThreadsLessOne
	}

	if a == nil {
		return nil, ErrMatrixAIsNil
	}
	if b == nil {
		return nil, ErrMatrixBIsNil
	}

	if len(*a) == 0 {
		return nil, ErrMatrixAIsEmpty
	}
	if len(*b) == 0 {
		return nil, ErrMatrixBIsEmpty
	}

	if len(*a) != len(*b) {
		return nil, ErrMatrixABVary
	}

	// Котроль выполнений группы горутин
	wg := &sync.WaitGroup{}

	// Канал заданий для воркеров размерностью в threads
	taskCh := make(chan task, threads)

	// Дейтсвие выполнится при выходе из функции
	defer func() {
		// Закрываем канал заданий и ждем, когда воркеры все прекратят свою работу
		close(taskCh)
		wg.Wait()
	}()

	// Инициализируем матрицу C
	c := &Matrix{}
	*c = make([]Row, len(*a))
	for i := range *a {
		if len((*a)[i]) != len((*b)[i]) {
			return nil, ErrMatrixABVary
		}

		(*c)[i] = make([]int, len((*a)[i]))
	}

	// Формируем набор воркеров-потоков количеством в threads
	wg.Add(threads)
	for i := 0; i < threads; i++ {
		go worker(wg, taskCh, a, b, c)
	}

	// Отправляем воркерам координаты заданий на перемножение
	for i := range *a {
		for j := range (*a)[i] {
			taskCh <- task{i, j}
		}
	}

	return c, nil
}

// Воркер для операции перемножения.
func worker(
	wg *sync.WaitGroup,
	taskCh chan task,
	a, b, c *Matrix,
) {
	defer wg.Done()

	// Перебираем задания из канала и выполяем перемножение
	for task := range taskCh {
		// Перемножаем элементы матриц A и B
		(*c)[task.i][task.j] = (*a)[task.i][task.j] * (*b)[task.i][task.j]
	}
}
