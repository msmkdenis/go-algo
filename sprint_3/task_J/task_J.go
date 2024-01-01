package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Чтобы выбрать самый лучший алгоритм для решения задачи, Гоша продолжил изучать разные сортировки.
На очереди сортировка пузырьком — https://ru.wikipedia.org/wiki/Сортировка_пузырьком

Её алгоритм следующий (сортируем по неубыванию):

На каждой итерации проходим по массиву, поочередно сравнивая пары соседних элементов.
Если элемент на позиции i больше элемента на позиции i + 1, меняем их местами.
После первой итерации самый большой элемент всплывёт в конце массива.
Проходим по массиву, выполняя указанные действия до тех пор, пока на очередной итерации не окажется,
что обмены больше не нужны, то есть массив уже отсортирован.
После не более чем n – 1 итераций выполнение алгоритма заканчивается,
так как на каждой итерации хотя бы один элемент оказывается на правильной позиции.

Помогите Гоше написать код алгоритма.
*/

func main() {

	const maxCapacity = 5 * 1024 * 1024
	buf := make([]byte, maxCapacity)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)

	var s []string
	for scanner.Scan() {
		line := scanner.Text()
		s = append(s, line)
		if line == "" {
			break
		}
	}
	capacity, _ := strconv.Atoi(s[0])
	unsorted := makeIntSlice(s[1], capacity)
	_ = bubbleSort(unsorted)
}

func makeIntSlice(s string, cap int) []int {
	result := make([]int, 0, cap)
	for _, v := range strings.Split(s, " ") {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}
	return result
}

func bubbleSort(arr []int) []int {
	counter := 0
	isSorted := true
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				counter++
				isSorted = false
			}
		}
		if counter != 0 {
			var answer strings.Builder
			for k, v := range arr {
				if k == len(arr)-1 {
					answer.WriteString(strconv.Itoa(v) + "\n")
				} else {
					answer.WriteString(strconv.Itoa(v) + " ")
				}
			}
			fmt.Print(answer.String())
		}
		counter = 0
	}
	if isSorted {
		var answer strings.Builder
		for k, v := range arr {
			if k == len(arr)-1 {
				answer.WriteString(strconv.Itoa(v) + "\n")
			} else {
				answer.WriteString(strconv.Itoa(v) + " ")
			}
		}
		fmt.Print(answer.String())
	}
	return arr
}
