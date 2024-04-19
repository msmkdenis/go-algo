package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/25597/run-report/112239250/

/*
-- ПРИНЦИП РАБОТЫ --
  Для решения задачи воспользуемся методом динамического программирования так как задача имеет оптимальную подструктуру
  и повторяющиеся подзадачи:

	Оптимальная подструктура:
		В данной задаче оптимальная подструктура означает, что оптимальное решение задачи может быть получено
		из оптимальных решений ее подзадач.	Наша исходная задача состоит в том, чтобы определить,
		можно ли разделить все набранные Гошей очки на две части с одинаковой суммой.

		Это решение может быть выведено из решений более мелких задач,
		таких как "можно ли разделить подмножество первых k очков на две части с одинаковой суммой".
		Если мы можем разделить каждое подмножество из первых k очков на две части с одинаковой суммой,
		то мы можем сделать то же самое и для всего набора очков.

	Повторяющиеся подзадачи:
		В процессе решения этой задачи мы сталкиваемся с повторяющимися подзадачами,
		когда нам нужно проверять возможность разделения подмножеств очков на две части с одинаковой суммой.

		Например, если у нас есть набор очков [3, 1, 4, 1], и мы хотим проверить,
		можно ли разделить 3 очка на две части с одинаковой суммой, мы должны проверить,
		можно ли разделить первые 2 очка на две части с одинаковой суммой, и так далее.

		Благодаря применению метода динамического программирования, мы можем сохранять результаты вычислений для каждой
		подзадачи и использовать их для избежания повторных вычислений.

-- РЕАЛИЗАЦИЯ --
  Сначала вычисляется общее количество очков, которые Гоша набрал за все партии.
  Если общее количество очков нечетное, то невозможно разделить их на две части с одинаковой суммой, и функция возвращает false.
  Если общее количество очков четное, то мы пытаемся найти такую сумму, которая составит половину от общего количества очков.
  Для этого используется алгоритм динамического программирования.

  Алгоритм строит булевый массив dp, где dp[i] равно true, если возможно набрать сумму i из подмножества набранных очков.
  После завершения выполнения алгоритма, если значение dp[target] (где target - половина от общего количества очков) равно true,
  значит, можно разделить очки на две части с одинаковой суммой, и функция возвращает true, в противном случае - false.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
  Временная сложность алгоритма определяется двумя вложенными циклами: внешний цикл по всем баллам, который работает за O(n),
  и внутренний цикл, который работает за O(target), где target - половина от общего количества очков.
  Таким образом, временная сложность составляет O(n * target), где n - количество партий, а target - сумма всех очков, деленная на 2.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
  Пространственная сложность алгоритма определяется размером массива dp, который имеет размер target + 1.
  Поэтому пространственная сложность составляет O(target).

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
  Алгоритм использует метод динамического программирования для определения возможности разбиения очков на две части с одинаковой суммой.
  Доказательство корректности заключается в том, что если сумма всех очков четная, и если существует подмножество баллов,
  сумма которых равна половине от общей суммы, то можно разделить все очки на две части с одинаковой суммой.
  Алгоритм динамического программирования обеспечивает полное исследование всех возможных вариантов разбиения,
  что гарантирует корректность результата.
*/

func canSplitScores(scores []int) bool {
	total := 0
	for _, score := range scores {
		total += score
	}

	// Если сумма очков нечётная, невозможно разбить на две части с одинаковой суммой
	if total%2 != 0 {
		return false
	}

	target := total / 2

	dp := make([]bool, target+1)
	dp[0] = true

	for _, score := range scores {
		for j := target; j >= score; j-- {
			dp[j] = dp[j] || dp[j-score]
		}
	}

	return dp[target]
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

func main() {
	lines := getInputData()

	n, _ := strconv.Atoi(lines[0])
	scoreStr := strings.Fields(lines[1])

	scores := make([]int, n)
	for i := 0; i < n; i++ {
		score, _ := strconv.Atoi(scoreStr[i])
		scores[i] = score
	}

	possible := canSplitScores(scores)
	if possible {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}