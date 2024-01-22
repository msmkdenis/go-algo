package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/23815/run-report/105531507/

/*
-- ПРИНЦИП РАБОТЫ --
Быстрая in-place сортировка основана на стратегии "разделяй и властвуй".
In-place означает, что все операции совершаются в текущем массиве, дополнительные массивы не создаются (доп. память не требуется).
1. Выбираем опорный элемент (обычно первый, последний или средний).
2. Разделяем массив относительно опорного на две части, при этом элементы "меньше" опорного окажутся слева, элементы "больше" - справа.
3. Рекурсивно вызываем сортировку для двух частей, который получились относительно индекса для разделения массива в операции 2.
4. Процесс будет продолжаться пока не останется один элемент в подмассиве (а значит массив отсортирован полностью).

-- РЕАЛИЗАЦИЯ --
Сортировать необходимо структуры по нескольким полям.
Для этого создадим функцию compare, осуществляющую сравнение элементов по необходимой логике.

В вызове функции quicksort осуществим три операции:
- отсортируем полученный массив относительно опорного элемента, одновременно определим индекс для разделения массива
- рекурсивно вызовем quicksort для двух массивов, полученных относительно первоначального массива с разделением по найденному индексу
- будем продолжать операции пока в подмассивах не останется по 1 элементу

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Алгоритм выбирает опорный элемент, сортирует элементы массива относительно опорного и определяет индекс для деления массива.
В результате массив отсортирован относительно опорного элемента: все элементы меньше опорного - слева, больше - справа.
По найденному индексу для разделения массива рекурсивно вызываем quicksort для двух частей массива.
Операции продолжаются пока после деления не останется только один элемент в подмассиве.
В результате все элементы массива окажутся отсортированными.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Временная сложность алгоритма быстрой сортировки:
- в худшем случае составляет О(n^2) - в случае если массив уже отсортирован или все элементы массива равны
- в среднем составляет O(n log n) по следующим причинам:

Разделение (partSort):
В каждом рекурсивном вызове быстрой сортировки выполняется операция разделения массива на две части.
Эта операция требует просмотра каждого элемента в массиве один раз, что занимает O(n) времени, где n - это количество элементов в массиве.

Рекурсия:
После разделения массива, быстрая сортировка рекурсивно вызывается для каждой из двух частей.
Если разделение происходит равномерно, то есть каждый раз массив делится пополам, то глубина рекурсии будет log n,
где n - это количество элементов в массиве.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Пространственная сложность алгоритма в среднем составит О(log n), в худшем случае составит O(n), например если массив отсортирован.
Дополнительная память потребуется для хранения стека вызовов (в силу использования рекурсии).
*/

type person struct {
	login   string
	solved  int
	penalty int
}

func main() {
	participants := getInputData()

	if len(participants) > 1 {
		quicksort(participants, 0, len(participants)-1)
	}

	for _, p := range participants {
		fmt.Println(p.login)
	}
}

func comparator(a, b person) bool {
	// сначала сравниваем по количеству задач
	if a.solved != b.solved {
		return a.solved < b.solved
	}

	// если равенство числа решённых задач
	// сравниваем по штрафам
	if a.penalty != b.penalty {
		return a.penalty > b.penalty
	}

	// если равенство решенных задач и штрафов
	// сравниваем по лексике
	if strings.Compare(a.login, b.login) > 0 {
		return true
	} else {
		return false
	}
}

func quicksort(arr []person, left, right int) {
	if left < right {
		divideIdx := partSort(arr, left, right, comparator)
		// разделяем массив на две части
		quicksort(arr, left, divideIdx)
		quicksort(arr, divideIdx+1, right)
	}
}

func partSort(arr []person, left, right int, compare func(a, b person) bool) int {
	//выбираем опорный элемент
	pivot := arr[(left+right)/2]
	for {
		// двигаемся слева направо до тех пор, пока не встретим значение ">" опорного
		for compare(pivot, arr[left]) {
			left++
		}
		// двигаемся справа налево до тех пор, пока не встретим значение "<" опорного
		for compare(arr[right], pivot) {
			right--
		}
		// определяем индекс для разделения массива
		if left >= right {
			return right
		}
		// меняем значения местами (в центре опорный элемент)
		arr[left], arr[right] = arr[right], arr[left]
		// сдвигаем указатели после замены, чтобы продолжить цикл
		left++
		right--
	}
}

func getInputData() []person {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	var bufStr string
	bufStr = scanner.Text()
	n, _ := strconv.Atoi(bufStr)

	users := make([]person, n)
	var user person
	for i := 0; i < n; i++ {
		scanner.Scan()
		bufStr = scanner.Text()
		strArr := strings.Split(bufStr, " ")
		user = person{}
		user.login = strArr[0]
		user.solved, _ = strconv.Atoi(strArr[1])
		user.penalty, _ = strconv.Atoi(strArr[2])

		users[i] = user
	}

	return users
}
