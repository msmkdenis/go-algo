package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
-- АЛГОРИТМ ВСТАВКИ ПОДАРЕННЫХ СТРОК В СТРОКУ S --
1. Инициализация переменных:
   - `accumulator` для отслеживания суммарного смещения вставок.
   - `result` представляет собой буфер, в который будет записываться результирующая строка.

2. Итерация по новой длине строки:
   Мы итерируемся по новой длине строки, которая будет включать в себя все вставки подаренных строк.

3. Проверка, остались ли подаренные строки:
   Если мин-куча `wordsToAdd` пуста, это означает, что все подаренные строки были добавлены,
   и мы добавляем оставшуюся часть строки `s` в результат и завершаем итерацию.

4. Извлечение подаренной строки для текущей позиции:
   Мы извлекаем подаренное слово из вершины мин-кучи `wordsToAdd`.

5. Добавление символов до текущей позиции вставки:
   Мы добавляем символы строки `s` до текущей позиции вставки подаренной строки.

6. Добавление подаренной строки:
   Мы добавляем подаренную строку в результат.

7. Обновление счетчика смещения итерации:
   Мы обновляем `accumulator` на длину добавленной строки, чтобы учесть смещение в следующей итерации.

8. Корректировка переменной итерации:
   Мы корректируем переменную итерации `i`, чтобы пропустить уже добавленные символы подаренной строки.

9. Вывод результирующей строки:
   В конце цикла мы выводим результирующую строку.
*/

func main() {
	input := getInputData()
	s := input[0]
	n, _ := strconv.Atoi(input[1])

	wordsToAdd := make(MinHeap, 0)
	newStrLength := len(s)
	for i := 0; i < n; i++ {
		listString := strings.Split(input[i+2], " ")
		t := listString[0]
		k, _ := strconv.Atoi(listString[1])
		wordsToAdd.Push(&word{pos: k, str: t})
		newStrLength += k
	}

	heap.Init(&wordsToAdd)

	accumulator := 0
	var result strings.Builder
	for i := 0; i < newStrLength; i++ {
		if wordsToAdd.Len() == 0 {
			result.WriteString(s[i-accumulator:])
			break
		}

		addingWord := heap.Pop(&wordsToAdd).(*word)

		for i != addingWord.pos+accumulator {
			result.WriteByte(s[i-accumulator])
			i++
		}
		result.WriteString(addingWord.str)
		accumulator += len(addingWord.str)
		i += len(addingWord.str) - 1
	}
	fmt.Print(result.String())
}

type MinHeap []*word

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i].pos < h[j].pos
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*word))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type word struct {
	pos int
	str string
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
