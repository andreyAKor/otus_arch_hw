package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/andreyAKor/otus_arch_hw/hw04/factory"
)

const (
	bufSize  int  = 1024 // Размер буффера на чтение/запись
	maxLines int  = 50   // Максимальное количество элементов для сортировки
	lineSerp byte = '\n' // Разделитель строк
)

var (
	ErrUnsupportedFile = errors.New("unsupported file")

	algo int    // Алгоритм сортировки
	in   string // Входной файл
	out  string // Выходной файл
)

func init() {
	flag.IntVar(&algo, "algo", 1, "Алгоритм сортировки: 1 - выбором, 2 - вставки, 3 - слиянием")
	flag.StringVar(&in, "in", "", "Входной файл")
	flag.StringVar(&out, "out", "", "Выходной файл")
}

func main() {
	flag.Parse()

	lines, err := read(in)
	if err != nil {
		log.Fatal(err)
	}
	if len(lines) > maxLines {
		lines = lines[:maxLines]
	}

	items, err := makeItems(lines)
	if err != nil {
		log.Fatal(err)
	}

	sorter := factory.NewSorter().Factory(algo)
	if err := write(out, sorter.Name(), sorter.Sort(items).([]int64)); err != nil {
		log.Fatal(err)
	}
}

// Чтение файла.
func read(in string) ([]string, error) {
	rf, err := os.Open(in)
	if err != nil {
		return nil, fmt.Errorf("read fail: %w", err)
	}
	defer rf.Close()

	fi, err := rf.Stat()
	if err != nil {
		return nil, fmt.Errorf("read fail: %w", err)
	}

	if !fi.Mode().IsRegular() {
		return nil, ErrUnsupportedFile
	}

	br := bufio.NewReaderSize(rf, bufSize)

	res := []string{}

	for {
		str, err := br.ReadString(lineSerp)

		if len(str) > 0 {
			res = append(res, str[:len(str)-1])
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return res, fmt.Errorf("read fail: %w", err)
		}
	}

	return res, nil
}

// Запись в файл.
func write(out, name string, items []int64) error {
	// Write file
	wf, err := os.Create(out)
	if err != nil {
		return fmt.Errorf("write fail: %w", err)
	}
	defer wf.Close()

	bw := bufio.NewWriterSize(wf, bufSize)
	defer bw.Flush()

	if _, err := bw.WriteString(fmt.Sprintf("%s\n", name)); err != nil {
		return fmt.Errorf("write fail: %w", err)
	}

	for _, item := range items {
		if _, err := bw.WriteString(fmt.Sprintf("%d\n", item)); err != nil {
			return fmt.Errorf("write fail: %w", err)
		}
	}

	return nil
}

// Формируем их строк список занчений.
func makeItems(lines []string) ([]int64, error) {
	res := []int64{}

	for _, str := range lines {
		val, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("write fail: %w", err)
		}

		res = append(res, int64(val))
	}

	return res, nil
}
