package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/andreyAKor/otus_arch_hw/hw02/matrix"
)

// Размерность матрицы N*N.
const n = 1000

func main() {
	// Готовим матрицы A и B
	a, b := makeMatrix(n)

	// Засекаем время
	start := time.Now()

	// Перемножение матриц A и B.
	// Причем первым параметром указали количество ядер в системе как количество потоков
	c, err := matrix.Mul(runtime.NumCPU(), a, b)
	if err != nil {
		panic(err)
	}

	// Сравниваем результаты перемножений последних элементов
	fmt.Printf("a: %v b: %v a*b: %v c: %v\n", (*a)[n-1][n-1], (*b)[n-1][n-1], (*a)[n-1][n-1]*(*b)[n-1][n-1], (*c)[n-1][n-1])

	// Итоговое время выполнения
	fmt.Printf("elapsed: %v\n", time.Since(start))
}

// Подготовка матриц A и B заданной размерности.
func makeMatrix(n int) (*matrix.Matrix, *matrix.Matrix) {
	var a, b matrix.Matrix

	a = make([]matrix.Row, n)
	b = make([]matrix.Row, n)

	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		b[i] = make([]int, n)

		for j := 0; j < n; j++ {
			a[i][j] = i + j
			b[i][j] = i + j + 1
		}
	}

	return &a, &b
}
