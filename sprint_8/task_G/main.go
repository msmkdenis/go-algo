/*
-- ПРИНЦИП РАБОТЫ --
  Алгоритм находит все вхождения шаблона A в последовательность X, допускающие сдвиг на константу.
  Для этого используется итеративный подход. Перебираются все возможные сдвиги в последовательности X,
  и для каждого сдвига вызывается функция поиска find, которая ищет вхождение шаблона A в последовательности X
  с учетом текущего сдвига. Если вхождение найдено, его позиция печатается, и сдвиг обновляется на следующую позицию после найденного вхождения.

-- РЕАЛИЗАЦИЯ --
  Функция main читает входные данные из файла с помощью функции getInputData.
  Далее, последовательность измерений X и шаблон A преобразуются в целочисленные массивы с помощью функции parseIntArray.
  Затем итеративно вызывается функция find для каждого сдвига в последовательности X.
  Функция find ищет вхождение шаблона A в последовательности X, начиная с заданной позиции start.
  Если вхождение найдено, возвращается его позиция. Иначе возвращается -1.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
  Временная сложность алгоритма зависит от длин последовательности X и шаблона A.
  Если n - длина последовательности X, а m - длина шаблона A, то временная сложность будет O(n*m),
  так как в худшем случае для каждого элемента последовательности X мы будем проверять весь шаблон A.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
  Пространственная сложность зависит от размера входных данных и составляет O(n + m),
  где n и m - длины последовательности X и шаблона A соответственно.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := getInputData()

	x := parseIntArray(strings.Fields(lines[1]))
	a := parseIntArray(strings.Fields(lines[3]))

	start := 0
	for {
		pos := find(x, a, start)
		if pos == -1 {
			break
		}
		fmt.Print(pos+1, " ")
		start = pos + 1
	}
}

func find(arr []int, pattern []int, start int) int {
	if len(arr) < len(pattern) {
		return -1
	}

	for pos := start; pos < len(arr)-len(pattern)+1; pos++ {
		match := true
		diff := pattern[0] - arr[pos]
		for offset := 0; offset < len(pattern); offset++ {
			if pattern[offset]-arr[pos+offset] != diff {
				match = false
				break
			}
		}
		if match {
			return pos
		}
	}
	return -1
}

func parseIntArray(arr []string) []int {
	res := make([]int, len(arr))
	for i, str := range arr {
		num, _ := strconv.Atoi(str)
		res[i] = num
	}
	return res
}

func getInputData() []string {
	input, _ := os.Open("input.txt")
	defer input.Close()

	const maxCapacity = 10240 * 1024
	buf := make([]byte, maxCapacity)

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(buf, maxCapacity)

	var s []string
	for scanner.Scan() {
		bufStr := scanner.Text()
		s = append(s, bufStr)
	}

	return s
}
