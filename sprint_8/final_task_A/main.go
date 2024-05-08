package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
-- ПРИНЦИП РАБОТЫ --
  Алгоритм нахождения общего префикса для распакованных строк основан на принципе последовательной распаковки строк и поиске общего префикса.
  Мы распаковываем каждую строку, начиная с первой, и сравниваем ее с общим префиксом.
  Если общий префикс пустой, значит, строки не имеют общего префикса, и работа завершается.
  Иначе мы обновляем общий префикс, находя наибольший общий префикс текущей распакованной строки с текущим общим префиксом.
  Далее процесс повторяется для всех строк ввода.

-- РЕАЛИЗАЦИЯ --
  Последовательно распаковываем каждую строку с помощью функции unpackString и
  находим общий префикс с помощью функции longestCommonPrefix.
  Если общий префикс становится пустым, работа завершается.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
  Временная сложность алгоритма зависит от длины входных строк и количества строк.
  Распаковка каждой строки занимает O(|s|) времени, где |s| - длина строки.
  Поиск общего префикса занимает O(m*n) времени, где m - количество строк, n - длина наибольшей строки.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
  Пространственная сложность алгоритма зависит от длины входных строк и количества строк.
  Она может быть O(|s|), где |s| - длина наибольшей распакованной строки, так как для распаковки используется буфер.
*/

func main() {
	lines := getInputData()
	n, _ := strconv.Atoi(lines[0])

	// Инициализация переменной для хранения общего префикса.
	commonPrefix := ""
	// Проход по всем строкам для обработки.
	for i := 1; i <= n; i++ {
		// Распаковка текущей строки.
		unpackedString := unpackString(lines[i])
		// Если это первая строка, то присвоить ее значение общему префиксу.
		if i == 1 {
			commonPrefix = unpackedString
		} else {
			// Найти общий префикс между текущим общим префиксом и распакованной строкой.
			commonPrefix = longestCommonPrefix(commonPrefix, unpackedString)
		}
		// Если общий префикс становится пустым, завершить выполнение программы.
		if len(commonPrefix) == 0 {
			return
		}
	}
	// Вывести найденный общий префикс.
	fmt.Println(commonPrefix)
}

// Функция для распаковки строки.
func unpackString(s string) string {
	var result strings.Builder
	var multipliersStack []int
	var toMultiplyStack []*strings.Builder

	// Посимвольный проход по строке.
	for i := 0; i < len(s); i++ {
		// Если текущий символ - цифра, то добавляем ее в стек множителей.
		if unicode.IsDigit(rune(s[i])) {
			multiplier, _ := strconv.Atoi(string(s[i]))
			multipliersStack = append(multipliersStack, multiplier)
			continue
		}
		// Если текущий символ - буква, то добавляем ее к текущей умножаемой строке (или к результату, если умножаемых строк в стеке нет).
		if unicode.IsLetter(rune(s[i])) {
			if len(toMultiplyStack) != 0 {
				toMultiplyStack[len(toMultiplyStack)-1].WriteByte(s[i])
			} else {
				result.WriteByte(s[i])
			}
			continue
		}
		// Если текущий символ - открывающая скобка, то создаем новый билдер строк для умножаемой строки.
		if rune(s[i]) == '[' {
			toMultiplyStack = append(toMultiplyStack, &strings.Builder{})
			continue
		}
		// Если текущий символ - закрывающая скобка, то извлекаем последний множитель и умножаемую строку из соответствующих стеков.
		if rune(s[i]) == ']' {
			multiplier := multipliersStack[len(multipliersStack)-1]
			multipliersStack = multipliersStack[:len(multipliersStack)-1]

			toMultiply := toMultiplyStack[len(toMultiplyStack)-1].String()
			toMultiplyStack = toMultiplyStack[:len(toMultiplyStack)-1]

			// Умножаем умножаемую строку на множитель и добавляем к текущему билдеру строк в стеке (или к результату, если стек пуст).
			for j := 0; j < multiplier; j++ {
				if len(toMultiplyStack) == 0 {
					result.WriteString(toMultiply)
				} else {
					toMultiplyStack[len(toMultiplyStack)-1].WriteString(toMultiply)
				}
			}
			continue
		}
	}
	// Возвращаем результат распаковки.
	return result.String()
}

// Функция для нахождения общего префикса между двумя строками.
func longestCommonPrefix(a, b string) string {
	var commonPrefix strings.Builder
	// Посимвольное сравнение строк.
	for i := range a {
		// Если символы не совпадают, возвращаем текущий общий префикс.
		if i == len(b) || a[i] != b[i] {
			return commonPrefix.String()
		}
		// Добавляем текущий символ к общему префиксу.
		commonPrefix.WriteByte(a[i])
	}
	// Возвращаем общий префикс.
	return commonPrefix.String()
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
