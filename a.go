package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// ключевая идея:
// чтобы сортировка была возможной, для каждого контейнера его сумма шаров должна равняться какой-то сумме шаров одного цвета
// то есть, если рассматривать каждый контенйнер с шарами как строку матрицы, то для каждой строки её сумма должна равняться сумме какого-либо столбца

func solve(r io.Reader) (string, error) {
	sc := bufio.NewScanner(r)
	sc.Scan()

	n, err := strconv.ParseInt(sc.Text(), 10, 64)
	if err != nil {
		return "no", err
	}

	rowSums := make([]int64, n)
	colSums := make([]int64, n)

	for i := int64(0); i < n; i++ {
		sc.Scan()

		ln := sc.Text()
		splittedLn := strings.Fields(ln)

		for j, ln := range splittedLn {
			num, err := strconv.ParseInt(ln, 10, 64)
			if err != nil {
				return "no", err
			}

			rowSums[i] += num
			colSums[j] += num
		}
	}

	slices.Sort(rowSums)
	slices.Sort(colSums)

	for i := int64(0); i < n; i++ {
		if rowSums[i] != colSums[i] {
			return "no", nil
		}
	}
	return "yes", nil
}

func main() {
	str, err := solve(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(str)
}
