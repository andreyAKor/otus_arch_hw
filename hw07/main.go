package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/andreyAKor/otus_arch_hw/hw07/factory"
)

const bufSize int = 1024 // Размер буффера на запись

var (
	limit int    // Максимальное количество итераций
	algo  string // Алгоритм генерации чисел
	out   string // Выходной файл
)

func init() {
	flag.IntVar(&limit, "limit", 1, "Максимальное количество итераций")
	flag.StringVar(&algo, "algo", "", "Алгоритм генерации чисел: fibonacci - Фибоначчи, binet - Бине")
	flag.StringVar(&out, "out", "", "Выходной файл")
}

// Клиентский код.
func main() {
	flag.Parse()

	fact, err := factory.Factory(algo)
	if err != nil {
		log.Fatal(err)
	}

	if err := write(out, fact.MakeOperation(limit)); err != nil {
		log.Fatal(err)
	}
}

// Запись в файл.
func write(out string, o factory.OperationIface) error {
	// Write file
	wf, err := os.Create(out)
	if err != nil {
		return fmt.Errorf("write fail: %w", err)
	}
	defer wf.Close()

	bw := bufio.NewWriterSize(wf, bufSize)
	defer bw.Flush()

	i := o.Iterator()
	for {
		v, err := i.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return fmt.Errorf("next iterate fail: %w", err)
		}

		if _, err := bw.WriteString(fmt.Sprintf("%d\n", v)); err != nil {
			return fmt.Errorf("write fail: %w", err)
		}
	}

	return nil
}
