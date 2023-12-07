package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Алла получила задание, связанное с мониторингом работы различных серверов.
Требуется понять, сколько времени обрабатываются определённые запросы на конкретных серверах.
Эту информацию нужно хранить в матрице, где номер столбца соответствуют идентификатору запроса, а номер строки — идентификатору сервера.
Алла перепутала строки и столбцы местами. С каждым бывает. Помогите ей исправить баг.

Есть матрица размера m × n. Нужно написать функцию, которая её транспонирует.

Транспонированная матрица получается из исходной заменой строк на столбцы.
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
	matrixCols := makeIntSlice(s[1])[0]
	if matrixRows == 0 && matrixCols == 0 {
		fmt.Println("")
	} else {
		matrix := s[2:(2 + matrixRows)]
		m := makeMatrix(matrix)
		_, stringMatrix := transpose(m)
		for _, v := range stringMatrix {
			fmt.Println(v)
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

func transpose(matrix [][]int) ([][]int, []string) {
	var result [][]int
	var m []string
	for i := 0; i < len(matrix[0]); i++ {
		var row []int
		var s strings.Builder
		for j := 0; j < len(matrix); j++ {
			row = append(row, matrix[j][i])
			if j < len(matrix)-1 {
				s.WriteString(strconv.Itoa(matrix[j][i]))
				s.WriteString(" ")
			} else {
				s.WriteString(strconv.Itoa(matrix[j][i]))
			}
		}
		result = append(result, row)
		m = append(m, s.String())
	}
	return result, m
}
