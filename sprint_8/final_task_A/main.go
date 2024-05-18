package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// https://contest.yandex.ru/contest/26133/run-report/114283955/

/*
-- ПРИНЦИП РАБОТЫ --
  Алгоритм нахождения общего префикса для распакованных строк основан на последовательной распаковке каждой строки и поиске общего префикса.
  Сначала считываются все строки, и первая строка используется для инициализации общего префикса.
  Затем последовательно распаковывается каждая строка и сравнивается с текущим общим префиксом.
  Если общий префикс становится пустым, дальнейшая обработка прекращается.

-- РЕАЛИЗАЦИЯ --
  Последовательно распаковываем каждую строку с помощью функции unpackString и
  находим общий префикс с помощью функции longestCommonPrefix.
  Если общий префикс становится пустым, работа завершается.

  Используются два основных этапа:
    распаковка строки с помощью функции unpackString
    поиск общего префикса с помощью функции longestCommonPrefix.
  Распаковка строки осуществляется с использованием стека для корректной обработки вложенных структур.
  Поиск общего префикса происходит путем поэтапного сравнения текущего префикса с каждой новой распакованной строкой.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
  Корректность распаковки строк гарантируется использованием стека, который обеспечивает правильную обработку вложенных структур.
  Первая распакованная строка инициализируется как общий префикс.
  Алгоритм нахождения общего префикса посимвольно сравнивает текущий префикс с каждой новой распакованной строкой.
  Если длина общего префикса становится равной 0, дальнейшее сравнение не имеет смысла, так как общий префикс отсутствует.
  Таким образом, будут рассмотрены все строки (если общий префикс не стал равен 0).

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
  Временная сложность алгоритма зависит от длины входных строк и количества строк.
  Распаковка занимает O(n*|s|) времени, где n - кол-во строк, |s| - длина наибольшей строки.
  Поиск общего префикса занимает O(n*|s|) времени, где n - количество строк, |s| - длина наибольшей строки
  (в худшем случае если все строки равны и надо перебрать все символы).
  Итого сложность составит: распаковка + поиск префикса O(2(n*|s|)) или O(n*|s|).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
  Пространственная сложность алгоритма зависит от длины входных строк и количества строк.
  В худшем случае потребуется O(n * |s|) памяти, где n - количество строк, |s| - длина наибольшей строки.
  Это включает:
  - Хранение всех строк: O(n * |s|).
  - Промежуточные структуры данных (стек множителей, билдеры строк и т.д.): O(|s|) для каждой строки в процессе распаковки.
  Общая пространственная сложность составляет O(n * |s|).
*/

func main() {
	lines := getInputData()
	n, _ := strconv.Atoi(lines[0])

	// Инициализируем commonPrefix распаковкой первой строки
	commonPrefix := unpackString(lines[1])
	// Проход по всем строкам, начиная со 2-й для обработки.
	for i := 2; i <= n; i++ {
		// Найти общий префикс между текущим общим префиксом и распакованной строкой.
		commonPrefix = longestCommonPrefix(commonPrefix, unpackString(lines[i]))
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
