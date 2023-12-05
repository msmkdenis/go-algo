package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
В первой строке задано n — количество строк матрицы. Во второй — количество столбцов m.
Числа m и n не превосходят 1000. В следующих n строках задана матрица.
Элементы матрицы — целые числа, по модулю не превосходящие 1000.
В последних двух строках записаны координаты элемента, соседей которого нужно найти.
Индексация начинается с нуля.
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var s []string
	for scanner.Scan() {
		line := scanner.Text()
		s = append(s, line)
		if line == "" {
			break
		}
	}
	matrixRows := makeIntSlice(s[0])[0]
	x := makeIntSlice(s[2+matrixRows])[0]
	y := makeIntSlice(s[3+matrixRows])[0]
	matrix := s[2:(2 + matrixRows)]
	neighbors := findNeighbors(makeMatrix(matrix), x, y)
	slices.Sort(neighbors)
	for n, v := range neighbors {
		if n == len(neighbors)-1 {
			fmt.Print(v)
		} else {
			fmt.Print(v, " ")
		}
	}
}

func makeMatrix(matrix []string) [][]int {
	var result [][]int
	for _, v := range matrix {
		result = append(result, makeIntSlice(v))
	}
	return result
}

func makeIntSlice(s string) []int {
	var result []int
	for _, v := range strings.Split(s, " ") {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}
	return result
}

func findNeighbors(slice [][]int, row, col int) []int {
	neighbors := make([]int, 0)

	// Check top neighbor
	if row > 0 {
		neighbors = append(neighbors, slice[row-1][col])
	}

	// Check bottom neighbor
	if row < len(slice)-1 {
		neighbors = append(neighbors, slice[row+1][col])
	}

	// Check left neighbor
	if col > 0 {
		neighbors = append(neighbors, slice[row][col-1])
	}

	// Check right neighbor
	if col < len(slice[row])-1 {
		neighbors = append(neighbors, slice[row][col+1])
	}

	return neighbors
}
